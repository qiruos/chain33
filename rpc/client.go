// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rpc chain33 RPC模块包含JSONRpc以及grpc
package rpc

import (
	"bytes"
	"encoding/hex"
	"time"

	"github.com/33cn/chain33/system/crypto/btcscript"
	"github.com/33cn/chain33/system/crypto/btcscript/script"

	"github.com/33cn/chain33/account"
	"github.com/33cn/chain33/client"
	"github.com/33cn/chain33/common"
	"github.com/33cn/chain33/common/address"
	"github.com/33cn/chain33/common/log/log15"
	"github.com/33cn/chain33/queue"
	"github.com/33cn/chain33/types"
)

var log = log15.New("module", "rpc")

type channelClient struct {
	client.QueueProtocolAPI
	accountdb *account.DB
}

// Init channel client
func (c *channelClient) Init(q queue.Client, api client.QueueProtocolAPI) {
	if api == nil {
		var err error
		api, err = client.New(q, nil)
		if err != nil {
			panic(err)
		}
	}
	c.QueueProtocolAPI = api
	c.accountdb = account.NewCoinsAccount(q.GetConfig())
}

// CreateRawTransaction create rawtransaction
func (c *channelClient) CreateRawTransaction(param *types.CreateTx) ([]byte, error) {
	if param == nil {
		log.Error("CreateRawTransaction", "Error", types.ErrInvalidParam)
		return nil, types.ErrInvalidParam
	}
	//构建交易时to地址不为空时需要检测地址的合法性
	if param.GetTo() != "" {
		if err := address.CheckAddress(param.GetTo()); err != nil {
			return nil, types.ErrInvalidAddress
		}
	}
	//因为历史原因，这里还是有部分token 的字段，但是没有依赖token dapp
	//未来这个调用可能会被废弃
	types.AssertConfig(c.QueueProtocolAPI)
	cfg := c.QueueProtocolAPI.GetConfig()
	execer := cfg.ExecName(cfg.GetCoinExec())
	if param.IsToken {
		execer = cfg.ExecName("token")
	}
	if param.Execer != "" {
		execer = param.Execer
	}
	reply, err := types.CallCreateTx(cfg, execer, "", param)
	if err != nil {
		return nil, err
	}

	//add tx fee setting
	tx := &types.Transaction{}
	err = types.Decode(reply, tx)
	if err != nil {
		return nil, err
	}
	tx.Fee = param.Fee
	//set proper fee if zero fee
	if tx.Fee <= 0 {
		proper, err := c.GetProperFee(nil)
		if err != nil {
			return nil, err
		}
		fee, err := tx.GetRealFee(proper.GetProperFee())
		if err != nil {
			return nil, err
		}
		tx.Fee = fee
	}

	return types.Encode(tx), nil
}

func (c *channelClient) ReWriteRawTx(param *types.ReWriteRawTx) ([]byte, error) {
	types.AssertConfig(c.QueueProtocolAPI)
	cfg := c.QueueProtocolAPI.GetConfig()
	if param == nil || param.Tx == "" {
		log.Error("ReWriteRawTx", "Error", types.ErrInvalidParam)
		return nil, types.ErrInvalidParam
	}

	tx, err := decodeTx(param.Tx)
	if err != nil {
		return nil, err
	}
	if param.To != "" {
		tx.To = param.To
	}
	if param.Fee != 0 && param.Fee > tx.Fee {
		tx.Fee = param.Fee
	}
	var expire int64
	if param.Expire != "" {
		expire, err = types.ParseExpire(param.Expire)
		if err != nil {
			return nil, err
		}
		tx.SetExpire(cfg, time.Duration(expire))
	}
	group, err := tx.GetTxGroup()
	if err != nil {
		return nil, err
	}

	//单笔交易直接返回
	if group == nil {
		txHex := types.Encode(tx)
		return txHex, nil
	}

	//交易组的处理
	index := param.Index
	if int(index) > len(group.GetTxs()) {
		return nil, types.ErrIndex
	}

	//修改交易组中所有成员交易
	if index <= 0 {
		if param.Fee != 0 && param.Fee > group.Txs[0].Fee {
			group.Txs[0].Fee = param.Fee
		}
		if param.Expire != "" {
			for i := 0; i < len(group.Txs); i++ {
				group.SetExpire(cfg, i, time.Duration(expire))
			}
		}
		group.RebuiltGroup()
		grouptx := group.Tx()
		txHex := types.Encode(grouptx)
		return txHex, nil
	}
	//修改交易组中指定成员交易
	index--
	if param.Fee != 0 && index == 0 && param.Fee > group.Txs[0].Fee {
		group.Txs[0].Fee = param.Fee
	}
	if param.Expire != "" {
		group.SetExpire(cfg, int(index), time.Duration(expire))
	}
	group.RebuiltGroup()
	grouptx := group.Tx()
	txHex := types.Encode(grouptx)
	return txHex, nil
}

// CreateRawTxGroup create rawtransaction for group
func (c *channelClient) CreateRawTxGroup(param *types.CreateTransactionGroup) ([]byte, error) {
	types.AssertConfig(c.QueueProtocolAPI)
	cfg := c.QueueProtocolAPI.GetConfig()
	if param == nil || len(param.Txs) <= 1 {
		return nil, types.ErrTxGroupCountLessThanTwo
	}
	var transactions []*types.Transaction
	for _, t := range param.Txs {
		txByte, err := common.FromHex(t)
		if err != nil {
			return nil, err
		}
		var transaction types.Transaction
		err = types.Decode(txByte, &transaction)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &transaction)
	}
	feeRate := cfg.GetMinTxFeeRate()
	//get proper fee rate
	proper, err := c.GetProperFee(nil)
	if err != nil {
		log.Error("CreateNoBalance", "GetProperFeeErr", err)
		return nil, err
	}
	if proper.GetProperFee() > feeRate {
		feeRate = proper.ProperFee
	}
	group, err := types.CreateTxGroup(transactions, feeRate)
	if err != nil {
		return nil, err
	}

	txGroup := group.Tx()
	txHex := types.Encode(txGroup)
	return txHex, nil
}

// CreateNoBalanceTxs create the multiple transaction with no balance
// 实际使用的时候要注意，一般情况下，不要传递 private key 到服务器端，除非是本地localhost 的服务。
func (c *channelClient) CreateNoBalanceTxs(in *types.NoBalanceTxs) (*types.Transaction, error) {
	types.AssertConfig(c.QueueProtocolAPI)
	cfg := c.QueueProtocolAPI.GetConfig()
	txNone := &types.Transaction{Execer: []byte(cfg.ExecName(types.NoneX)), Payload: []byte("no-fee-transaction")}
	txNone.To = address.ExecAddress(string(txNone.Execer))
	txNone, err := types.FormatTx(cfg, cfg.ExecName(types.NoneX), txNone)
	if err != nil {
		return nil, err
	}
	//不设置时默认为永不超时
	if in.Expire == "" {
		in.Expire = "0"
	}
	expire, err := types.ParseExpire(in.Expire)
	if err != nil {
		return nil, err
	}
	//交易组只需要设置单笔交易超时
	txNone.SetExpire(cfg, time.Duration(expire))
	isParaTx := false
	transactions := []*types.Transaction{txNone}
	for _, txhex := range in.TxHexs {
		tx, err := decodeTx(txhex)
		if err != nil {
			return nil, err
		}
		if types.IsParaExecName(string(tx.GetExecer())) {
			isParaTx = true
		}
		transactions = append(transactions, tx)
	}

	//平行链下不允许设置高度作为过期判定, issue#706
	if expire > 0 && expire <= types.ExpireBound && isParaTx {
		return nil, types.ErrInvalidExpire
	}

	feeRate := cfg.GetMinTxFeeRate()
	//get proper fee rate
	proper, err := c.GetProperFee(nil)
	if err != nil {
		log.Error("CreateNoBalance", "GetProperFeeErr", err)
		return nil, err
	}
	if proper.GetProperFee() > feeRate {
		feeRate = proper.ProperFee
	}
	group, err := types.CreateTxGroup(transactions, feeRate)
	if err != nil {
		return nil, err
	}
	err = group.Check(cfg, 0, cfg.GetMinTxFeeRate(), cfg.GetMaxTxFee())
	if err != nil {
		return nil, err
	}

	newtx := group.Tx()
	//如果可能要做签名
	if in.PayAddr != "" || in.Privkey != "" {
		rawTx := hex.EncodeToString(types.Encode(newtx))
		req := &types.ReqSignRawTx{Addr: in.PayAddr, Privkey: in.Privkey, Expire: in.Expire, TxHex: rawTx, Index: 1}
		signedTx, err := c.ExecWalletFunc("wallet", "SignRawTx", req)
		if err != nil {
			return nil, err
		}
		return decodeTx(signedTx.(*types.ReplySignRawTx).TxHex)
	}
	return newtx, nil
}

func decodeTx(hexstr string) (*types.Transaction, error) {
	var tx types.Transaction
	data, err := common.FromHex(hexstr)
	if err != nil {
		return nil, err
	}
	err = types.Decode(data, &tx)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

// GetAddrOverview get overview of address
func (c *channelClient) GetAddrOverview(parm *types.ReqAddr) (*types.AddrOverview, error) {
	err := address.CheckAddress(parm.Addr)
	if err != nil {
		return nil, types.ErrInvalidAddress
	}
	reply, err := c.QueueProtocolAPI.GetAddrOverview(parm)
	if err != nil {
		log.Error("GetAddrOverview", "Error", err.Error())
		return nil, err
	}
	//获取地址账户的余额通过account模块
	addrs := make([]string, 1)
	addrs[0] = parm.Addr
	accounts, err := c.accountdb.LoadAccounts(c.QueueProtocolAPI, addrs)
	if err != nil {
		log.Error("GetAddrOverview", "Error", err.Error())
		return nil, err
	}
	if len(accounts) != 0 {
		reply.Balance = accounts[0].Balance
	}
	return reply, nil
}

// GetBalance get balance
func (c *channelClient) GetBalance(in *types.ReqBalance) ([]*types.Account, error) {
	// in.AssetExec & in.AssetSymbol 新增参数，
	// 不填时兼容原来的调用
	if in.AssetExec == "" || in.AssetSymbol == "" {
		in.AssetSymbol = "bty"
		in.AssetExec = c.GetConfig().GetCoinExec()
		return c.accountdb.GetBalance(c.QueueProtocolAPI, in)
	}

	acc, err := account.NewAccountDB(c.QueueProtocolAPI.GetConfig(), in.AssetExec, in.AssetSymbol, nil)
	if err != nil {
		log.Error("GetBalance", "Error", err.Error())
		return nil, err
	}
	return acc.GetBalance(c.QueueProtocolAPI, in)
}

// GetAllExecBalance get balance of exec
func (c *channelClient) GetAllExecBalance(in *types.ReqAllExecBalance) (*types.AllExecBalance, error) {
	types.AssertConfig(c.QueueProtocolAPI)
	cfg := c.QueueProtocolAPI.GetConfig()
	addr := in.Addr
	err := address.CheckAddress(addr)
	if err != nil {
		if err = address.CheckMultiSignAddress(addr); err != nil {
			return nil, types.ErrInvalidAddress
		}
	}
	var addrs []string
	addrs = append(addrs, addr)
	allBalance := &types.AllExecBalance{Addr: addr}
	for _, exec := range types.AllowUserExec {
		execer := cfg.ExecName(string(exec))
		params := &types.ReqBalance{
			Addresses:   addrs,
			Execer:      execer,
			StateHash:   in.StateHash,
			AssetExec:   in.AssetExec,
			AssetSymbol: in.AssetSymbol,
		}
		res, err := c.GetBalance(params)
		if err != nil {
			continue
		}
		if len(res) < 1 {
			continue
		}
		acc := res[0]
		if acc.Balance == 0 && acc.Frozen == 0 {
			continue
		}
		execAcc := &types.ExecAccount{Execer: execer, Account: acc}
		allBalance.ExecAccount = append(allBalance.ExecAccount, execAcc)
	}
	return allBalance, nil
}

// GetTotalCoins get total of coins
func (c *channelClient) GetTotalCoins(in *types.ReqGetTotalCoins) (*types.ReplyGetTotalCoins, error) {
	//获取地址账户的余额通过account模块
	resp, err := c.accountdb.GetTotalCoins(c.QueueProtocolAPI, in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DecodeRawTransaction decode rawtransaction
func (c *channelClient) DecodeRawTransaction(param *types.ReqDecodeRawTransaction) (*types.Transaction, error) {
	var tx types.Transaction
	bytes, err := common.FromHex(param.TxHex)
	if err != nil {
		return nil, err
	}
	err = types.Decode(bytes, &tx)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

// GetTimeStatus get status of time
func (c *channelClient) GetTimeStatus() (*types.TimeStatus, error) {
	ntpTime := common.GetRealTimeRetry(c.GetConfig().GetModuleConfig().NtpHosts, 2)
	local := time.Now()
	if ntpTime.IsZero() {
		return &types.TimeStatus{NtpTime: "", LocalTime: local.Format("2006-01-02 15:04:05"), Diff: 0}, nil
	}
	diff := local.Sub(ntpTime) / time.Second
	return &types.TimeStatus{NtpTime: ntpTime.Format("2006-01-02 15:04:05"), LocalTime: local.Format("2006-01-02 15:04:05"), Diff: int64(diff)}, nil
}

// GetExecBalance get balance with exec by channelclient
func (c *channelClient) GetExecBalance(in *types.ReqGetExecBalance) (*types.ReplyGetExecBalance, error) {
	//通过account模块获取地址账户在合约中的余额
	resp, err := c.accountdb.GetExecBalance(c.QueueProtocolAPI, in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *channelClient) getWalletRecoverAddr(param *types.ReqGetWalletRecoverAddr) ([]byte, []byte, error) {

	ctrPub, err := common.FromHex(param.GetCtrPubKey())
	if err != nil {
		log.Error("getWalletRecoverAddr", "control pubKey", param.GetCtrPubKey(), "from hex err", err)
		return nil, nil, types.ErrFromHex
	}

	recoverPub, err := common.FromHex(param.GetRecoverPubKey())
	if err != nil {
		log.Error("getWalletRecoverAddr", "recover pubKey", param.GetCtrPubKey(), "from hex err", err)
		return nil, nil, types.ErrFromHex
	}

	pkScript, err := script.NewWalletRecoveryScript(ctrPub, recoverPub, param.GetRelativeDelayHeight())
	if err != nil {
		log.Error("getWalletRecoverAddr", "new wallet recover script err", err)
		return nil, nil, err
	}
	return recoverPub, pkScript, nil
}

// GetWalletRecoverAddr get wallet recover chain33 address
func (c *channelClient) GetWalletRecoverAddr(req *types.ReqGetWalletRecoverAddr) (*types.ReplyString, error) {

	if len(req.GetCtrPubKey()) <= 0 || len(req.GetRecoverPubKey()) <= 0 ||
		req.GetRelativeDelayHeight() <= 0 {
		log.Error("GetWalletRecoverAddr", "invalid req", req.String())
		return nil, types.ErrInvalidParam
	}
	_, pkScript, err := c.getWalletRecoverAddr(req)
	if err != nil {
		log.Error("GetWalletRecoverAddr", "getWalletRecoverAddr err", err)
		return nil, err
	}
	reply := &types.ReplyString{}
	reply.Data = address.PubKeyToAddr(script.Script2PubKey(pkScript))
	return reply, nil
}

// SignWalletRecoverTx sign wallet recover transaction
func (c *channelClient) SignWalletRecoverTx(req *types.ReqSignWalletRecoverTx) (*types.ReplySignRawTx, error) {

	if req.GetWalletRecoverParam() == nil || len(req.GetRawTx()) <= 0 {
		log.Error("SignWalletRecoverTx", "invalid req", req.String())
		return nil, types.ErrInvalidParam
	}
	privKeyHex := req.PrivKey
	if privKeyHex == "" {
		reply, err := c.ExecWalletFunc("wallet", "DumpPrivkey", &types.ReqString{Data: req.SignAddr})
		if err != nil {
			log.Error("SignWalletRecoverTx", "execWalletFunc err", err)
			return nil, err
		}
		privKeyHex = reply.(*types.ReplyString).GetData()
	}

	recoverPub, wrScript, err := c.getWalletRecoverAddr(req.GetWalletRecoverParam())
	if err != nil {
		log.Error("SignWalletRecoverTx", "getWalletRecoverAddr err", err)
		return nil, err
	}
	tx := &types.Transaction{}
	txBytes, err := common.FromHex(req.GetRawTx())
	if err != nil || types.Decode(txBytes, tx) != nil {
		log.Error("SignWalletRecoverTx", "rawTx", req.GetRawTx(), "decode raw tx err", err)
		return nil, types.ErrDecode
	}
	tx.Signature = nil
	signMsg := types.Encode(tx)
	privKey, err := common.FromHex(privKeyHex)
	if err != nil {
		log.Error("SignWalletRecoverTx", "privKey from hex err", err)
		return nil, types.ErrFromHex
	}
	_, pk := script.NewBtcKeyFromBytes(privKey)
	// if wallet recover with recover addr
	isRetrive := bytes.Equal(pk.SerializeCompressed(), recoverPub)
	sig, pubKey, err := script.GetWalletRecoverySignature(isRetrive, signMsg, privKey,
		wrScript, req.GetWalletRecoverParam().GetRelativeDelayHeight())
	if err != nil {
		log.Error("SignWalletRecoverTx", "get wallet recover sig err", err)
		return nil, err
	}

	tx.Signature = &types.Signature{
		Ty:        btcscript.ID,
		Signature: sig,
		Pubkey:    pubKey,
	}
	reply := &types.ReplySignRawTx{}
	reply.TxHex = hex.EncodeToString(types.Encode(tx))
	return reply, nil
}
