changelog

## [1.65.5](https://github.com/33cn/chain33/compare/v1.65.4...v1.65.5) (2021-10-19)


### Bug Fixes

* Adjust github action ([9246b83](https://github.com/33cn/chain33/commit/9246b830a84d9a52ae140d09754ed71291c81548))

## [1.65.4](https://github.com/33cn/chain33/compare/v1.65.3...v1.65.4) (2021-10-15)


### Bug Fixes

* **doc:** release 1.65.4 ([4f53148](https://github.com/33cn/chain33/commit/4f531488049a79640121ba5d950166939dedaebd))

# [1.66.0](https://github.com/33cn/chain33/compare/v1.65.2...v1.66.0) (2021-10-15)


### Bug Fixes

* 🐛version control: Add github action for auto publish release and tag version ([22642e1](https://github.com/33cn/chain33/commit/22642e187aecaa21d5904c5d82e459fc6a0f72c4))
* **chain:** add ticker stop method ([aac09d4](https://github.com/33cn/chain33/commit/aac09d45e0ee64f77e81cc36da569444da511ccd))
* chunk key when delete ([f559759](https://github.com/33cn/chain33/commit/f5597596f5f20e02c29eb699d3cefa53cd42b95b))
* chunkStatusCacheMutex unlock bug ([cd7bdc8](https://github.com/33cn/chain33/commit/cd7bdc8111538c4ac1926f54ce85f78071624fe5))
* close pubsub ([aca60b8](https://github.com/33cn/chain33/commit/aca60b86d6e5d8b613d10e15f20181e3b445c2eb))
* close stream without delay after reading ([678ff0a](https://github.com/33cn/chain33/commit/678ff0a78e32d5be98cd123717f1f396467dac4b))
* compact block body in localdb ([6c6a0aa](https://github.com/33cn/chain33/commit/6c6a0aab49ca885ffb60a25ac79c2667ef3f573e))
* compact db ([c704b6a](https://github.com/33cn/chain33/commit/c704b6a3fcb501f546a19a204247cbbc6db69d3f))
* compact db after deleting chunk ([6371fcf](https://github.com/33cn/chain33/commit/6371fcf336241e5430c30a94c303c1dad84c2ac2))
* dht unit test ([1566a58](https://github.com/33cn/chain33/commit/1566a58369fe4c8f4307dc2fcd156840151ddb6d))
* do not push searched peers into peer channel ([1523267](https://github.com/33cn/chain33/commit/152326733cbaba07e0af7d894b2036481db1b58c))
* fetch chunk routine bug ([1a5b5be](https://github.com/33cn/chain33/commit/1a5b5bef855d1966e7944b146a4d2ad456651619))
* fix ci and add manually auto publish release ([28febf7](https://github.com/33cn/chain33/commit/28febf7face3b8842641c72751150a5bd550017a))
* ignore peers without addr when saving peers ([7a15e22](https://github.com/33cn/chain33/commit/7a15e22a5dcf62d426c304dabbdedd6d8ded0264))
* index bug ([fd6a114](https://github.com/33cn/chain33/commit/fd6a114b69aa169f542f1288d532aeb40cd2f4b1))
* libp2p stream leak ([53988d7](https://github.com/33cn/chain33/commit/53988d78a8335ff4ae95a588ea3a970f38acb80c))
* mustFetchChunk context ([c00ce7d](https://github.com/33cn/chain33/commit/c00ce7d738e38b0d7563975f46b9bca404ef558e))
* push to channel unblockedly ([4d884f1](https://github.com/33cn/chain33/commit/4d884f1bf525ffa9296f097763f775cf2a7378e5))
* query chunk records from peers in routing tableif there is no given peer ([5362c0a](https://github.com/33cn/chain33/commit/5362c0a273b4ffc001343f1ab80ba0779b188d07))
* query public ip ([a73dd61](https://github.com/33cn/chain33/commit/a73dd61ebd6ef96d7c67beb7862e10e8ad96b006))
* refresh peer info ([7e95977](https://github.com/33cn/chain33/commit/7e95977ac73d028d776b4556c42b00c559cf5ecb))
* retry to exec block when error occurs in chunk download mode ([330f8b2](https://github.com/33cn/chain33/commit/330f8b2a03a92782314bf18a412d883f039338ed))
* set dht to server mode ([8de2fb9](https://github.com/33cn/chain33/commit/8de2fb98df00900e3d9bc8bbd9b84d35c53cd070))
* unit test ([2269a15](https://github.com/33cn/chain33/commit/2269a1596766ff0bc95e862467a1b657a2bebf93))
* use peerlist instead of best peers ([198f580](https://github.com/33cn/chain33/commit/198f5800ff818b03e19d5d37360655e07f7eacbc))


### Features

* **deps:** bump github.com/decred/base58 from 1.0.2 to 1.0.3 ([cfbde5e](https://github.com/33cn/chain33/commit/cfbde5ef9e4acca23a7e82a47c777f87779969c2))
* **deps:** bump github.com/influxdata/influxdb from 1.7.9 to 1.9.5 ([162a75c](https://github.com/33cn/chain33/commit/162a75c2457d09a7fb1e99a0ae0bd8c29d7a83a0))
* **deps:** bump github.com/multiformats/go-multiaddr ([44b7c10](https://github.com/33cn/chain33/commit/44b7c10f9c6632b188e85ee0b646e06443547c4f))

# [1.66.0](https://github.com/33cn/chain33/compare/v1.65.2...v1.66.0) (2021-10-15)


### Bug Fixes

* 🐛version control: Add github action for auto publish release and tag version ([22642e1](https://github.com/33cn/chain33/commit/22642e187aecaa21d5904c5d82e459fc6a0f72c4))
* **chain:** add ticker stop method ([aac09d4](https://github.com/33cn/chain33/commit/aac09d45e0ee64f77e81cc36da569444da511ccd))
* chunk key when delete ([f559759](https://github.com/33cn/chain33/commit/f5597596f5f20e02c29eb699d3cefa53cd42b95b))
* chunkStatusCacheMutex unlock bug ([cd7bdc8](https://github.com/33cn/chain33/commit/cd7bdc8111538c4ac1926f54ce85f78071624fe5))
* close pubsub ([aca60b8](https://github.com/33cn/chain33/commit/aca60b86d6e5d8b613d10e15f20181e3b445c2eb))
* close stream without delay after reading ([678ff0a](https://github.com/33cn/chain33/commit/678ff0a78e32d5be98cd123717f1f396467dac4b))
* compact block body in localdb ([6c6a0aa](https://github.com/33cn/chain33/commit/6c6a0aab49ca885ffb60a25ac79c2667ef3f573e))
* compact db ([c704b6a](https://github.com/33cn/chain33/commit/c704b6a3fcb501f546a19a204247cbbc6db69d3f))
* compact db after deleting chunk ([6371fcf](https://github.com/33cn/chain33/commit/6371fcf336241e5430c30a94c303c1dad84c2ac2))
* dht unit test ([1566a58](https://github.com/33cn/chain33/commit/1566a58369fe4c8f4307dc2fcd156840151ddb6d))
* do not push searched peers into peer channel ([1523267](https://github.com/33cn/chain33/commit/152326733cbaba07e0af7d894b2036481db1b58c))
* fetch chunk routine bug ([1a5b5be](https://github.com/33cn/chain33/commit/1a5b5bef855d1966e7944b146a4d2ad456651619))
* fix ci and add manually auto publish release ([28febf7](https://github.com/33cn/chain33/commit/28febf7face3b8842641c72751150a5bd550017a))
* ignore peers without addr when saving peers ([7a15e22](https://github.com/33cn/chain33/commit/7a15e22a5dcf62d426c304dabbdedd6d8ded0264))
* index bug ([fd6a114](https://github.com/33cn/chain33/commit/fd6a114b69aa169f542f1288d532aeb40cd2f4b1))
* libp2p stream leak ([53988d7](https://github.com/33cn/chain33/commit/53988d78a8335ff4ae95a588ea3a970f38acb80c))
* mustFetchChunk context ([c00ce7d](https://github.com/33cn/chain33/commit/c00ce7d738e38b0d7563975f46b9bca404ef558e))
* push to channel unblockedly ([4d884f1](https://github.com/33cn/chain33/commit/4d884f1bf525ffa9296f097763f775cf2a7378e5))
* query chunk records from peers in routing tableif there is no given peer ([5362c0a](https://github.com/33cn/chain33/commit/5362c0a273b4ffc001343f1ab80ba0779b188d07))
* query public ip ([a73dd61](https://github.com/33cn/chain33/commit/a73dd61ebd6ef96d7c67beb7862e10e8ad96b006))
* refresh peer info ([7e95977](https://github.com/33cn/chain33/commit/7e95977ac73d028d776b4556c42b00c559cf5ecb))
* retry to exec block when error occurs in chunk download mode ([330f8b2](https://github.com/33cn/chain33/commit/330f8b2a03a92782314bf18a412d883f039338ed))
* set dht to server mode ([8de2fb9](https://github.com/33cn/chain33/commit/8de2fb98df00900e3d9bc8bbd9b84d35c53cd070))
* unit test ([2269a15](https://github.com/33cn/chain33/commit/2269a1596766ff0bc95e862467a1b657a2bebf93))
* use peerlist instead of best peers ([198f580](https://github.com/33cn/chain33/commit/198f5800ff818b03e19d5d37360655e07f7eacbc))


### Features

* **deps:** bump github.com/decred/base58 from 1.0.2 to 1.0.3 ([cfbde5e](https://github.com/33cn/chain33/commit/cfbde5ef9e4acca23a7e82a47c777f87779969c2))
* **deps:** bump github.com/influxdata/influxdb from 1.7.9 to 1.9.5 ([162a75c](https://github.com/33cn/chain33/commit/162a75c2457d09a7fb1e99a0ae0bd8c29d7a83a0))
* **deps:** bump github.com/multiformats/go-multiaddr ([44b7c10](https://github.com/33cn/chain33/commit/44b7c10f9c6632b188e85ee0b646e06443547c4f))

# [6.4.0](https://github.com/33cn/chain33/compare/v6.3.0...v6.4.0) (2021-10-15)


### Bug Fixes

* 🐛version control: Add github action for auto publish release and tag version ([22642e1](https://github.com/33cn/chain33/commit/22642e187aecaa21d5904c5d82e459fc6a0f72c4))
* **chain:** add ticker stop method ([aac09d4](https://github.com/33cn/chain33/commit/aac09d45e0ee64f77e81cc36da569444da511ccd))
* chunk key when delete ([f559759](https://github.com/33cn/chain33/commit/f5597596f5f20e02c29eb699d3cefa53cd42b95b))
* chunkStatusCacheMutex unlock bug ([cd7bdc8](https://github.com/33cn/chain33/commit/cd7bdc8111538c4ac1926f54ce85f78071624fe5))
* close pubsub ([aca60b8](https://github.com/33cn/chain33/commit/aca60b86d6e5d8b613d10e15f20181e3b445c2eb))
* close stream without delay after reading ([678ff0a](https://github.com/33cn/chain33/commit/678ff0a78e32d5be98cd123717f1f396467dac4b))
* compact block body in localdb ([6c6a0aa](https://github.com/33cn/chain33/commit/6c6a0aab49ca885ffb60a25ac79c2667ef3f573e))
* compact db ([c704b6a](https://github.com/33cn/chain33/commit/c704b6a3fcb501f546a19a204247cbbc6db69d3f))
* compact db after deleting chunk ([6371fcf](https://github.com/33cn/chain33/commit/6371fcf336241e5430c30a94c303c1dad84c2ac2))
* dht unit test ([1566a58](https://github.com/33cn/chain33/commit/1566a58369fe4c8f4307dc2fcd156840151ddb6d))
* do not push searched peers into peer channel ([1523267](https://github.com/33cn/chain33/commit/152326733cbaba07e0af7d894b2036481db1b58c))
* fetch chunk routine bug ([1a5b5be](https://github.com/33cn/chain33/commit/1a5b5bef855d1966e7944b146a4d2ad456651619))
* fix ci and add manually auto publish release ([28febf7](https://github.com/33cn/chain33/commit/28febf7face3b8842641c72751150a5bd550017a))
* ignore peers without addr when saving peers ([7a15e22](https://github.com/33cn/chain33/commit/7a15e22a5dcf62d426c304dabbdedd6d8ded0264))
* index bug ([fd6a114](https://github.com/33cn/chain33/commit/fd6a114b69aa169f542f1288d532aeb40cd2f4b1))
* libp2p stream leak ([53988d7](https://github.com/33cn/chain33/commit/53988d78a8335ff4ae95a588ea3a970f38acb80c))
* mustFetchChunk context ([c00ce7d](https://github.com/33cn/chain33/commit/c00ce7d738e38b0d7563975f46b9bca404ef558e))
* push to channel unblockedly ([4d884f1](https://github.com/33cn/chain33/commit/4d884f1bf525ffa9296f097763f775cf2a7378e5))
* query chunk records from peers in routing tableif there is no given peer ([5362c0a](https://github.com/33cn/chain33/commit/5362c0a273b4ffc001343f1ab80ba0779b188d07))
* query public ip ([a73dd61](https://github.com/33cn/chain33/commit/a73dd61ebd6ef96d7c67beb7862e10e8ad96b006))
* refresh peer info ([7e95977](https://github.com/33cn/chain33/commit/7e95977ac73d028d776b4556c42b00c559cf5ecb))
* retry to exec block when error occurs in chunk download mode ([330f8b2](https://github.com/33cn/chain33/commit/330f8b2a03a92782314bf18a412d883f039338ed))
* return nil when no result for list result ([75bf83a](https://github.com/33cn/chain33/commit/75bf83a484677951f9ef6bf5aee56938340a0f2a))
* set dht to server mode ([8de2fb9](https://github.com/33cn/chain33/commit/8de2fb98df00900e3d9bc8bbd9b84d35c53cd070))
* unit test ([2269a15](https://github.com/33cn/chain33/commit/2269a1596766ff0bc95e862467a1b657a2bebf93))
* use peerlist instead of best peers ([198f580](https://github.com/33cn/chain33/commit/198f5800ff818b03e19d5d37360655e07f7eacbc))
* 在没有新区块时, 完成历史数据的推送 ([cf8ba79](https://github.com/33cn/chain33/commit/cf8ba79582617e796185056cc858de0b84835ccc))


### Features

* **deps:** bump github.com/decred/base58 from 1.0.2 to 1.0.3 ([cfbde5e](https://github.com/33cn/chain33/commit/cfbde5ef9e4acca23a7e82a47c777f87779969c2))
* **deps:** bump github.com/influxdata/influxdb from 1.7.9 to 1.9.5 ([162a75c](https://github.com/33cn/chain33/commit/162a75c2457d09a7fb1e99a0ae0bd8c29d7a83a0))
* **deps:** bump github.com/multiformats/go-multiaddr ([44b7c10](https://github.com/33cn/chain33/commit/44b7c10f9c6632b188e85ee0b646e06443547c4f))

# Changelog
All notable changes to this project will be documented in this file.

## [6.0.2]
### Changed
- changed cli version cmd return json format and added title app localdb version info
- change cli seed generate -l 0 from json format to only string

## [6.0.1]
### Changed
- Update configuration files from p2p/verMix to p2p/verMin from[@libangzhu](https://github.com/libangzhu).
- Update p2p version Logic(if you do not fill in the range of p2p version,then verMin=version,verMax=verMin+1) from[@libangzhu](https://github.com/libangzhu).

# [6.2]
### Changed
- Update dapp coins command line name, 'bty' not supported any more, 'coins' recommended
