# Changelog

## Unreleased
## 1.4.1

*November 28, 2022*

### Application

* [\#2780](https://github.com/gridiron-zone/gridiron/pull/2780) Bump tibc-go version to v0.4.2
* [\#2779](https://github.com/gridiron-zone/gridiron/pull/2779) Bump up gridmod version to v1.7.2
* [\#2777](https://github.com/gridiron-zone/gridiron/pull/2777) Add SetIAVLCacheSize and SetIAVLDisableFastNode
* [\#2775](https://github.com/gridiron-zone/gridiron/pull/2775) Remove group module

## 1.4.0

*November 15, 2022*
### Application

* [\#2759](https://github.com/gridiron-zone/gridiron/pull/2759) Fix export error when with flag `--for-zero-height`
* [\#2766](https://github.com/gridiron-zone/gridiron/pull/2766) Bump up cosmos sdk to v0.46.4
* [\#2768](https://github.com/gridiron-zone/gridiron/pull/2768) Bump up ibc-go to v5.0.1
* [\#2770](https://github.com/gridiron-zone/gridiron/pull/2770) Bump up gridmod to v1.7.0 & Bump up tibc-go to v0.4.0 
* [gridmod \#309](https://github.com/irisnet/irismod/pull/309) Refactor nft with cosmos-sdk nft module.
* [gridmod \#308](https://github.com/irisnet/irismod/pull/308) Coinswap module adds unilateral injection liquidity function.

### API Breaking Changes

* [gridmod \#309](https://github.com/irisnet/irismod/pull/309) GRPC method `Owner` rename to `NFTsOfOwner`, Remove deprecated `Queries` api
### Bug Fixes

* [gridmod \#304](https://github.com/irisnet/irismod/pull/304) Fix nft module import error.
* [gridmod \#314](https://github.com/irisnet/irismod/pull/314) Fix `addLiquidity` panic error.

## 1.3.0

*March 19, 2022*

### Application

* [\#2735](https://github.com/gridiron-zone/gridiron/pull/2735) Bump up gridmod
* [\#2734](https://github.com/gridiron-zone/gridiron/pull/2734) Bump up TIBC-Go
* [TIBC-Go \#247](https://github.com/bianjieai/tibc-go/pull/103) Support cross chain MT transfer via TIBC
* [gridmod \#247](https://github.com/irisnet/irismod/pull/247) Added the Farm Proposal function, allowing users to submit on-chain proposals to apply for a specified amount of GRID in GRIDnet’s community pool as farming rewards.
* [gridmod \#249](https://github.com/irisnet/irismod/pull/249) An added gas fee of 5,000 GRID, with a tax rate of 40%, for the creation of new liquidity pools.
* [gridmod \#245](https://github.com/irisnet/irismod/pull/245) Compatibilized & expanded the functions of the NFT module to match ERC-721 spec, and to support application requirements in a more flexible manner.
* [gridmod \#269](https://github.com/irisnet/irismod/pull/269) Introduced ERC-1155 compatible MT (Multi Token) module

## 1.2.0

*November 03th, 2021*

### Application

* [\#2681](https://github.com/gridiron-zone/gridiron/pull/2681) Bump cosmos-sdk version to [v0.44.2](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.44.2)
* [\#2668](https://github.com/gridiron-zone/gridiron/pull/2668) Integrated tibc protocol
* [\#2623](https://github.com/gridiron-zone/gridiron/pull/2623) Import farm module
* [gridmod \#219](https://github.com/irisnet/irismod/pull/219) Refactor coinswap module
* [gridmod \#189](https://github.com/irisnet/irismod/pull/189) Enhance nft module

### Breaking Changes

* [gridmod \#219](https://github.com/irisnet/irismod/pull/219) Liquidity tokens are named as *lpt-{number}*, and the existing liquidity will be automatically modified during the upgrade
* [cosmos-sdk \#10041](https://github.com/cosmos/cosmos-sdk/pull/10041) Remove broadcast & encode legacy REST endpoints. Please see the [REST Endpoints Migration guide](https://docs.cosmos.network/master/migrations/rest.html) to migrate to the new REST endpoints.

## 1.1.1

*April 21th, 2021*

### Application

* [\#2611](https://github.com/gridiron-zone/gridiron/pull/2611) Bump cosmos-sdk version to [v0.42.4](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.42.4)
* [\#2611](https://github.com/gridiron-zone/gridiron/pull/2611) Bump tendermint version to [v0.34.9](https://github.com/tendermint/tendermint/releases/tag/v0.34.9)

## 1.1.0

*March 26th, 2021*

### CLI

* [\#2592](https://github.com/gridiron-zone/gridiron/pull/2592) Bump gridmod version to [v1.4.0](https://github.com/irisnet/irismod/releases/tag/v1.4.0)

### Application

* [\#2602](https://github.com/gridiron-zone/gridiron/pull/2602) Bump cosmos-sdk version to [v0.42.2](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.42.3)
* [\#2588](https://github.com/gridiron-zone/gridiron/pull/2588) Bump tendermint verion to [v0.34.8](https://github.com/tendermint/tendermint/releases/tag/v0.34.8)
* [\#2592](https://github.com/gridiron-zone/gridiron/pull/2592) Bump gridmod version to [v1.4.0](https://github.com/irisnet/irismod/releases/tag/v1.4.0)
* [\#2589](https://github.com/gridiron-zone/gridiron/issues/2589) Migrate gridiron from v1.0 to v1.1

## 1.0.1

*February 18th, 2021*

### Application

This release fixes a downstream security issue which impacts Cosmos SDK users.
See the [Tendermint v0.34.7 SDK changelog](https://github.com/tendermint/tendermint/blob/v0.34.x/CHANGELOG.md#v0347) for details.

* [\#2573](https://github.com/gridiron-zone/gridiron/pull/2573) Bump cosmos-sdk version to [v0.41.3](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.41.3)
* [\#2571](https://github.com/gridiron-zone/gridiron/pull/2571) Bump tendermint verion to [v0.34.7](https://github.com/tendermint/tendermint/releases/tag/v0.34.7)

## 1.0.0

*February 9th, 2021*

### CLI

* [\#2541](https://github.com/gridiron-zone/gridiron/pull/2541) Bump cosmos-sdk version to [v0.41.0](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.41.0)
* [\#2515](https://github.com/gridiron-zone/gridiron/pull/2515) Bump tendermint verion to [v0.34.3](https://github.com/tendermint/tendermint/releases/tag/v0.34.3)
* [\#2567](https://github.com/gridiron-zone/gridiron/pull/2567) Bump gridmod version to [v1.3.1](https://github.com/irisnet/irismod/releases/tag/v1.3.1)
* [\#2505](https://github.com/gridiron-zone/gridiron/pull/2505) Remove duplicate cmd
* [\#2154](https://github.com/gridiron-zone/gridiron/issues/2154) Support native token unit conversion in command

### Application

* [\#2541](https://github.com/gridiron-zone/gridiron/pull/2541) Bump cosmos-sdk version to [v0.41.0](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.41.0)
* [\#2515](https://github.com/gridiron-zone/gridiron/pull/2515) Bump tendermint verion to [v0.34.3](https://github.com/tendermint/tendermint/releases/tag/v0.34.3)
* [\#2567](https://github.com/gridiron-zone/gridiron/pull/2567) Bump gridmod version to [v1.3.1](https://github.com/irisnet/irismod/releases/tag/v1.3.1)
* [\#2551](https://github.com/gridiron-zone/gridiron/pull/2551) Disable repeated service invocation
* [\#2542](https://github.com/gridiron-zone/gridiron/pull/2542) Migrate withdraw infos
* [\#2524](https://github.com/gridiron-zone/gridiron/pull/2524) Fix proto package and path
* [\#2518](https://github.com/gridiron-zone/gridiron/pull/2518) Move abandoned tokens to communityTax
* [\#2512](https://github.com/gridiron-zone/gridiron/pull/2512) Normalize msg and genesis validation
* [\#2484](https://github.com/gridiron-zone/gridiron/pull/2484) Bump cosmos-sdk version to [v0.40.0](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.40.0)
* [\#2484](https://github.com/gridiron-zone/gridiron/pull/2484) Bump tendermint verion to [v0.34.1](https://github.com/tendermint/tendermint/releases/tag/v0.34.1)
* [\#2502](https://github.com/gridiron-zone/gridiron/pull/2502) Bump gridmod version to [v1.2.0](https://github.com/irisnet/irismod/releases/tag/v1.2.0)
* [\#2501](https://github.com/gridiron-zone/gridiron/issues/2501) Merge the swagger docs of cosmos-sdk and gridmod
* [\#2485](https://github.com/gridiron-zone/gridiron/pull/2485) Support key import 0.16.3 and earlier version keystore file
* [\#2488](https://github.com/gridiron-zone/gridiron/pull/2488) Restrict the use of certain token for specified Msg
* [\#2098](https://github.com/gridiron-zone/gridiron/issues/2098) Migrate gridcli test
* [\#2097](https://github.com/gridiron-zone/gridiron/issues/2097) Add scripts to migrate genesis data from v0.16.3
* [\#2090](https://github.com/gridiron-zone/gridiron/issues/2090) Refactor guardian module
* [\#2089](https://github.com/gridiron-zone/gridiron/issues/2089) Overwrite mint module
* [\#2500](https://github.com/gridiron-zone/gridiron/issues/2500) Migrate GRIDhub modules to gridmod
* [\#2381](https://github.com/gridiron-zone/gridiron/pull/2381) Rebuild gridiron v1.0 using cosmos-sdk v0.40

## 0.16.3

*Aug 25th, 2020*

### Application

* [\#4d4c06](https://github.com/gridiron-zone/gridiron/commit/4d4c06a6dfdccc4271734ef6eef95960f000384f) Bump ledger-cosmos-go to support cosmos ledger 2.0

## 0.16.2

*Apr 9th, 2020*

### Tendermint

* [\#110](https://github.com/irisnet/tendermint/pull/110) Defend against DoS attacks

## 0.16.1

*Jan 9th, 2020*

### CLI

* [\#2109](https://github.com/gridiron-zone/gridiron/issues/2109) Fix the bug that the hash lock can not be specified in the CLI

### Application

* [\#2118](https://github.com/gridiron-zone/gridiron/pull/2118) Improve the UX of snapshot
* [\#2119](https://github.com/gridiron-zone/gridiron/issues/2119) Prune iavlstore version using rootstore version

## 0.16.0

*Nov 22th, 2019*

### Breaking Changes

* [API Breaking Changes](./docs/light-client/CHANGELOG.md#v0160)
* [\#1912](https://github.com/gridiron-zone/gridiron/issues/1912) Update prometheus metrics

### LCD (REST API)

* [\#1858](https://github.com/gridiron-zone/gridiron/issues/1858) Add new function - AddLiquidity
* [\#1859](https://github.com/gridiron-zone/gridiron/issues/1859) Add new function - RemoveLiquidity
* [\#1860](https://github.com/gridiron-zone/gridiron/issues/1860) Add new function - Swap Coin
* [\#1861](https://github.com/gridiron-zone/gridiron/issues/1861) Add new function - Query Exchange
* [\#1948](https://github.com/gridiron-zone/gridiron/issues/1948) Add client for HTLC

### CLI

* [\#1948](https://github.com/gridiron-zone/gridiron/issues/1948) Add client for HTLC
* [\#2030](https://github.com/gridiron-zone/gridiron/issues/2030) Fix CLI test for HTLC

### Application

* [\#1858](https://github.com/gridiron-zone/gridiron/issues/1858) Add new function - AddLiquidity
* [\#1859](https://github.com/gridiron-zone/gridiron/issues/1859) Add new function - RemoveLiquidity
* [\#1860](https://github.com/gridiron-zone/gridiron/issues/1860) Add new function - Swap Coin
* [\#1861](https://github.com/gridiron-zone/gridiron/issues/1861) Add new function - Query Exchange
* [\#1872](https://github.com/gridiron-zone/gridiron/issues/1872) Replace time.Time with uint64 in coinswap msgs
* [\#1811](https://github.com/gridiron-zone/gridiron/issues/1811) Add docs for coin-swap
* [\#1879](https://github.com/gridiron-zone/gridiron/issues/1879) Add tags for coinswap module
* [\#1910](https://github.com/gridiron-zone/gridiron/issues/1910) Add coin-flow for coin-swap module
* [\#1912](https://github.com/gridiron-zone/gridiron/issues/1912) Update prometheus metrics
* [\#1936](https://github.com/gridiron-zone/gridiron/issues/1936) Update denom's specification
* [\#1941](https://github.com/gridiron-zone/gridiron/issues/1941) Refactor document website documentation
* [\#1946](https://github.com/gridiron-zone/gridiron/issues/1946) Add new function - Create HTLC
* [\#1949](https://github.com/gridiron-zone/gridiron/issues/1949) Add new function - Claim and Refund in HTLC
* [\#1965](https://github.com/gridiron-zone/gridiron/issues/1965) Add user docs and specification for HTLC
* [\#1984](https://github.com/gridiron-zone/gridiron/issues/1984) Update HTLC docs
* [\#1985](https://github.com/gridiron-zone/gridiron/issues/1985) Code formatting and cleaning
* [\#1988](https://github.com/gridiron-zone/gridiron/issues/1988) Add HTLC genesis
* [\#1991](https://github.com/gridiron-zone/gridiron/issues/1991) Add Coinswap feature doc
* [\#1995](https://github.com/gridiron-zone/gridiron/issues/1995) Improve HTLC
* [\#2008](https://github.com/gridiron-zone/gridiron/issues/2008) Enhance HTLC genesis test
* [\#2013](https://github.com/gridiron-zone/gridiron/issues/2013) Modify timestamp type in LCD
* [\#2015](https://github.com/gridiron-zone/gridiron/issues/2015) Add HTLC feature docs

### Tendermint

* [\#1880](https://github.com/gridiron-zone/gridiron/issues/1880) Ensure generated private keys are valid
* [\#1885](https://github.com/gridiron-zone/gridiron/issues/1885) tendermint use Go modules instead of dep
* [\#1908](https://github.com/gridiron-zone/gridiron/issues/1908) Update p2p to v0.32
* [\#1909](https://github.com/gridiron-zone/gridiron/issues/1909) Support boltdb
* [\#1913](https://github.com/gridiron-zone/gridiron/issues/1913) Remove db from in favor of tendermint/tm-db
* [\#1917](https://github.com/gridiron-zone/gridiron/issues/1917) Create a peer state in consensus reactor before the peer is started
* [\#1918](https://github.com/gridiron-zone/gridiron/issues/1918) Improve error message returned from AddSignatureFromPubKey
* [\#1919](https://github.com/gridiron-zone/gridiron/issues/1919) Self node in addrbook.json and node tries to dial itself
* [\#1920](https://github.com/gridiron-zone/gridiron/issues/1920) Fix profiler blocking the entire node
* [\#1921](https://github.com/gridiron-zone/gridiron/issues/1921) Exit if SwitchToConsensus fails
* [\#1922](https://github.com/gridiron-zone/gridiron/issues/1922) Improve transaction search

## 0.15.5

*Oct 30th, 2019*

### Tendermint

Bump Tendermint version to irisnet/tendermint [v0.31.3](https://github.com/irisnet/tendermint/releases/tag/v0.31.3) to fix the consensus security bug.

## 0.15.4

*Oct 14th, 2019*

### Tendermint

Bump Tendermint version to irisnet/tendermint [v0.31.2](https://github.com/irisnet/tendermint/releases/tag/v0.31.2) to fix the p2p panic error.

## 0.15.3

*Oct 2th, 2019*

### Tendermint

Bump Tendermint version to irisnet/tendermint [v0.31.1](https://github.com/irisnet/tendermint/releases/tag/v0.31.1) to fix the p2p panic error.

## 0.15.2

*Sep 11th, 2019*

### Application

* [\#1940](https://github.com/gridiron-zone/gridiron/pull/1940) Do not update gov params when network not equal mainnet
* [\#1945](https://github.com/gridiron-zone/gridiron/pull/1945) Fix protocol loading on replay-last-block

## 0.15.1

*Aug 22th, 2019*

### Application

* [\#1895](https://github.com/gridiron-zone/gridiron/issues/1895) Run tendermint cleanupFunc before Exit
* [\#1897](https://github.com/gridiron-zone/gridiron/issues/1897) Fix possible panic on exporting genesis file

## 0.15.0

*Aug 20th, 2019*

### LCD (REST API)

* [\#1473](https://github.com/gridiron-zone/gridiron/issues/1473) Add APIs for Tokens
* [\#1521](https://github.com/gridiron-zone/gridiron/issues/1521) Add APIs for Gateways
* [\#1745](https://github.com/gridiron-zone/gridiron/issues/1745) POST APIs only generate transactions, not broadcast transactions
* [\#1750](https://github.com/gridiron-zone/gridiron/issues/1750) Move params query from `/gov/params` to `/params`
* [\#1802](https://github.com/gridiron-zone/gridiron/issues/1802) Remove key-related APIs and enhance API `tx/broadcast`

### CLI

* [\#1750](https://github.com/gridiron-zone/gridiron/issues/1750) Replace `gridcli gov query-params` by `gridcli params`

### Application

* [\#1336](https://github.com/gridiron-zone/gridiron/issues/1336) Ensure field length checking in the service module
* [\#1466](https://github.com/gridiron-zone/gridiron/issues/1466) Support multi-signature account
* [\#1468](https://github.com/gridiron-zone/gridiron/issues/1468) Add asset module
* [\#1493](https://github.com/gridiron-zone/gridiron/issues/1493) Back up keys using keystore file
* [\#1511](https://github.com/gridiron-zone/gridiron/issues/1511) Split the export command into reset and export
* [\#1544](https://github.com/gridiron-zone/gridiron/issues/1544) Refactor CommunityPool and BurnedCoins into AccAddresses
* [\#1568](https://github.com/gridiron-zone/gridiron/issues/1568) Optimize Dockerfile
* [\#1603](https://github.com/gridiron-zone/gridiron/issues/1603) Refactor total supply to support multiple tokens
* [\#1604](https://github.com/gridiron-zone/gridiron/issues/1604) Enhance gov TallyResult
* [\#1677](https://github.com/gridiron-zone/gridiron/issues/1677) Refactor coin_type to support multiple tokens
* [\#1728](https://github.com/gridiron-zone/gridiron/issues/1728) Refactor gov module and add proposal types `PlainText` and `TokenAddition`
* [\#1757](https://github.com/gridiron-zone/gridiron/issues/1757) Add a random number generator
* [\#1783](https://github.com/gridiron-zone/gridiron/issues/1783) Optimize error messages
* [\#1643](https://github.com/gridiron-zone/gridiron/issues/1643) Fix tx search by tx.height
* [\#1854](https://github.com/gridiron-zone/gridiron/issues/1854) Enforce unbonding_time for redelegation
* [\#1867](https://github.com/gridiron-zone/gridiron/issues/1867) Enable memo validation for "flagged" accounts
* [\#1851](https://github.com/gridiron-zone/gridiron/issues/1851) Reduce gas consumption to support larger data storage

### Tendermint

* Update to irisnet/tendermint [v0.31.0](https://github.com/irisnet/tendermint/releases/tag/v0.31.0)
  * [\#1379](https://github.com/gridiron-zone/gridiron/issues/1379) [p2p] Simplify GetSelectionWithBias for addressbook
  * [\#1432](https://github.com/gridiron-zone/gridiron/issues/1432) [crypto] Allow PubKeyMultisigThreshold to unmarshal into crypto.PubKey
  * [\#1451](https://github.com/gridiron-zone/gridiron/issues/1451) [db] Close WriteBatch to prevent memory leak
  * [\#1460](https://github.com/gridiron-zone/gridiron/issues/1460) [kms] Shut down tmkms when a validator enters proposal round crashes the app
  * [\#1554](https://github.com/gridiron-zone/gridiron/issues/1554) [consensus] Log block status on replay block
  * [\#1646](https://github.com/gridiron-zone/gridiron/issues/1646) [consensus] Fix possible halt by resetting TriggeredTimeoutPrecommit before starting the next height
  * [\#1647](https://github.com/gridiron-zone/gridiron/issues/1647) [consensus] Flush WAL on stop to prevent data corruption during a graceful shutdown
  * [\#1648](https://github.com/gridiron-zone/gridiron/issues/1648) [p2p] Fix reconnecting report duplicate ID error due to race condition between adding peer to peerSet and starting it
  * [\#1649](https://github.com/gridiron-zone/gridiron/issues/1649) [p2p] Do not panic when filter times out
  * [\#1650](https://github.com/gridiron-zone/gridiron/issues/1650) [p2p] Reject all-zero shared secrets in the Diffie-Hellman step of secret-connection
  * [\#1660](https://github.com/gridiron-zone/gridiron/issues/1660) [instrumentation] Add chain_id label for all metrics
  * [\#1665](https://github.com/gridiron-zone/gridiron/issues/1665) [rpc] Return maxPerPage if per_page is greater than max
  * [\#1679](https://github.com/gridiron-zone/gridiron/issues/1679) [p2p] Fix nil pointer deference on DNS resolution failure
  * [\#1692](https://github.com/gridiron-zone/gridiron/issues/1692) [blockchain] Fix pool timer leak bug
  * [\#1693](https://github.com/gridiron-zone/gridiron/issues/1693) [db] Close Iterator in RemoteDB
  * [\#1697](https://github.com/gridiron-zone/gridiron/issues/1697) [lite] Fix error proxy endpoints `validators` in lite client
  * [\#1702](https://github.com/gridiron-zone/gridiron/issues/1702) [blockchain] Update the maxHeight when a peer is removed
  * [\#1703](https://github.com/gridiron-zone/gridiron/issues/1703) [mempool] Stop gossiping tx's back to where they come from
  * [\#1704](https://github.com/gridiron-zone/gridiron/issues/1704) [rpc] Disable compression for HTTP client to prevent GZIP-bomb DoS attacks
  * [\#1714](https://github.com/gridiron-zone/gridiron/issues/1714) [mempool] Bound mempool memory usage

## 0.14.1

*May 31th, 2019*

### LCD (REST API)

* [\#1486](https://github.com/gridiron-zone/gridiron/issues/1486) Ensure `/bank/account/{address}` has consistent json output
* [\#1495](https://github.com/gridiron-zone/gridiron/issues/1495) Improve error handling for query parameters

### Application

* [\#1506](https://github.com/gridiron-zone/gridiron/issues/1506) Enforce `unbonding_time` regardless of validator status

## 0.14.0

*May 27th, 2019*

### LCD (REST API)

* [\#1245](https://github.com/gridiron-zone/gridiron/issues/1245) Follow best-practice URI naming guide -- API BREAKING!
* [\#1416](https://github.com/gridiron-zone/gridiron/issues/1416) Drop three useless distribution queries -- API BREAKING!
* [\#1444](https://github.com/gridiron-zone/gridiron/pull/1444) Fix `/bank/token-stats` output format -- API BREAKING!
* [\#1374](https://github.com/gridiron-zone/gridiron/issues/1374) Use `Querier` pattern to improve query implementation
* [\#1426](https://github.com/gridiron-zone/gridiron/issues/1426) Add query for community tax
* [\#1386](https://github.com/gridiron-zone/gridiron/issues/1386) Fix `memo` support bug in LCD

### CLI

* [\#1245](https://github.com/gridiron-zone/gridiron/issues/1245) Move `sign` and `broadcast` subcmd under `tx` cmd
* [\#1375](https://github.com/gridiron-zone/gridiron/issues/1375) Unify the output formats of return data
* [\#1411](https://github.com/gridiron-zone/gridiron/issues/1411) Fix incorrect json indent output for `keys` commands
* [\#1419](https://github.com/gridiron-zone/gridiron/pull/1419) Fix incorrect decimal output in plain-text format
* [\#1443](https://github.com/gridiron-zone/gridiron/issues/1443) Allow users to generate send tx offline

### Application

* [\#1383](https://github.com/gridiron-zone/gridiron/issues/1383) Improve result tags for unbond and redelegate
* [\#1409](https://github.com/gridiron-zone/gridiron/issues/1409) Fix validation bug for `community_tax` parameter
* [\#1422](https://github.com/gridiron-zone/gridiron/issues/1422) Replace dep with Go Module

### Tendermint

* Update to irisnet/tendermint [v0.28.0](https://github.com/irisnet/tendermint/releases/tag/v0.28.0)
  * [\#1408](https://github.com/gridiron-zone/gridiron/issues/1408) [mempool] Fix `broadcastTxRoutine` leak
  * [\#1428](https://github.com/gridiron-zone/gridiron/issues/1428) [rpc] Fix `/tx_search` bug when results are empty
  * [\#1429](https://github.com/gridiron-zone/gridiron/issues/1429) [privval] Retry `RemoteSigner` connections on error
  * [\#1430](https://github.com/gridiron-zone/gridiron/issues/1430) [privval] Memorize pubkey on startup
  * [\#1431](https://github.com/gridiron-zone/gridiron/issues/1431) [p2p] Make `SecretConnection` thread safe
  * [\#1434](https://github.com/gridiron-zone/gridiron/issues/1434) [consensus] Log `peerID` on ignored votes
  * [\#1435](https://github.com/gridiron-zone/gridiron/issues/1435) [rpc] Include peer's remote IP in `/net_info`
  * [\#1436](https://github.com/gridiron-zone/gridiron/issues/1436) [crypto] Update btcd fork for rare signRFC6979 bug
  * [\#1438](https://github.com/gridiron-zone/gridiron/issues/1438) [privval] Fix race between sign and ping requests
  * [\#1439](https://github.com/gridiron-zone/gridiron/issues/1439) [p2p] Fix MITM bug on `SecretConnection`
  * [\#1440](https://github.com/gridiron-zone/gridiron/issues/1440) [node] Start `EventBus` and `IndexerService` before first block

### Documentation

* [\#1376](https://github.com/gridiron-zone/gridiron/issues/1376) Improve v0.13.1 docs

## 0.13.1

*Mar 22th, 2019*

### LCD (REST API)

* [\#1339](https://github.com/gridiron-zone/gridiron/pull/1339) Add pagination params for lcd validators query
* [\#1355](https://github.com/gridiron-zone/gridiron/pull/1355) Fix pagination error message
* [\#1360](https://github.com/gridiron-zone/gridiron/pull/1360) Add query API for delegator rewards

### CLI

* [\#1360](https://github.com/gridiron-zone/gridiron/pull/1360) Add query command for delegator rewards

### Application

* [\#1329](https://github.com/gridiron-zone/gridiron/pull/1329) Improve error message for insufficient balance
* [\#1340](https://github.com/gridiron-zone/gridiron/pull/1340) Remove coin flow tags if tx is out of gas
* [\#1341](https://github.com/gridiron-zone/gridiron/pull/1341) Check validator existence and status before getting its pubkey
* [\#1344](https://github.com/gridiron-zone/gridiron/pull/1344) Reset the init value for metrics
* [\#1354](https://github.com/gridiron-zone/gridiron/pull/1354) Fix the bug of metric data accumulation
* [\#1362](https://github.com/gridiron-zone/gridiron/pull/1362) Fix testnet build for Docker
* [\#1370](https://github.com/gridiron-zone/gridiron/pull/1370) Add more tags for Redelgate TxResult

### Tendermint

* [\#51](https://github.com/irisnet/tendermint/pull/51) Update to irisnet/Tendermint **v0.27.4**
  * [\#44](https://github.com/irisnet/tendermint/pull/44) [p2p] Cleanup rejected inbound connections
  * [\#45](https://github.com/irisnet/tendermint/pull/45) [consensus] Fix consensus round issue
  * [\#46](https://github.com/irisnet/tendermint/pull/46) [mempool] Check max msg size in `CheckTx()`
  * [\#47](https://github.com/irisnet/tendermint/pull/47) [mempool] Fix the bug of LRU cache update
  * [\#48](https://github.com/irisnet/tendermint/pull/48) [p2p] Fix infinite loop in `AddrBook`
  * [\#50](https://github.com/irisnet/tendermint/pull/50) [p2p] Fix FlushStop() in `MConnection`

## 0.12.3

*February 27th, 2019*

* Fix issue of build config for Ledger Nano
* Fix Dockerfile to be compatible with Ledger build
* Improve documents about address prefix/fee/chain-id for the mainnet

## 0.12.2

*February 26th, 2019*

* Fix the document issue

## 0.12.2-rc0

*February 26th, 2019*

* Set max commission rate and max commission change rate to 100%
* Implement coin flow record feature for "internal transactions"
* Support Ledger Nano S and KMS
* Update the default gas_price_threshold to be 6000grid-nano
* No slashing for Censorship or Downtime
* No slashing for non-voting for proposals
* Configure the default build environment as mainnet
* Set the default gas limit to be 50000
* Fix wrong withdraw address
* Fix gas simulate issue
* Display correct validator consensus pubkey
* Improve the documents

## 0.12.1

*February 14th, 2019*

* Fix the bug of repeatability check about evidence in the tendermint
* Change the invariant level for mainnet

## 0.12.0

*February 11th, 2019*

* [grid] Refactor and add more promethus metrics for monitor
* [grid] Enrich log message for all modules
* [grid] Close all unclosed iterators
* [grid] Add invariant check level configuration in grid.toml
* [gridcli] Add share percent in lcd unbond and redelegate
* [gridcli] Fix tx search bug by page and size query parameter name
* [gridcli] Improve error message for gov module
* [gridlcd] Upgrade swagger-ui to 3.0 which can support dynamic query parameters
* [gridtool] Remove monitor command
* [tendermint] Fix DynamicVerifier for large validator set changes
* [tendermint] Optimize txs search to handle huge search result
* [doc] Improve slashing document

## 0.11.0

*January 25th, 2019*

* [grid] Handle the expected abort during replay last block
* [grid] Go through and make sure all panic are reasonable
* [grid] Add the set-withdraw-address feature
* [grid] Update default param value of Upgrade Threshold and Critical Proposal
* [grid] Fix bug that evidence age doesn't take effect

* [gridcli] Improve the upgrade query-signals cmd

* [tendermint] Avoid one evidence be committed multiple times

## 0.11.0-rc0

*January 21th, 2019*

* [grid] Move the upgrade tally threshold into the software upgrade proposal
* [grid] Limit the size of transaction
* [grid] Tendermint's blockstore also needs to consume gas when storing transactions
* [grid] Proposer censorship slashing to prevent the proposer from submitting block containing garbage tx data
* [grid] The proposer must deposit 30% of the mindeposit when submitting the proposal
* [grid] Make more check about every msg's ValidateBasic()
* [grid] Add flag --output-file to save export result and ensure result is consistent
* [grid] Add new param service/TxSizeLimit to limit the service tx size
* [grid] Block mint doesn't depend on BFT time
* [grid] Fix infinite gas meter utilization during aborted ante handler executions
* [grid] Auto-config bech32 prefixes based on network type
* [grid] Improve the system logs

* [gridcli] Make the result of `gridcli tendermint tx` readable
* [gridcli] Improve the output format of the query proposals
* [gridcli] Enhance the query-signals cmd to print the accumulated signal voting power percent

* [gridtool] Add support for consensus address and pubkey

## 0.10.2

*January 17th, 2019*

* [grid] The proposer must deposit 30% of the mindeposit when submitting the proposal

## 0.10.1

*January 17th, 2019*

* [gridcli] Fix issue about query validator information
* [gridcli] Fix cli query proposals error

## 0.10.0

*January 16th, 2019*

* [grid] Add flag --output-file to save export result and ensure result is consistent
* [grid] Improve invariant checking coverage and fix distribution bugs
* [gridcli] Make the result of `gridcli tendermint tx` readable
* [gridcli] Query cmd return details about software upgrade and tax usage proposal
* [tendermint] Fix the inconformity of too many evidences check
* [tendermint] Fix replay bug of `grid export`

## 0.10.0-rc0

*January 8th, 2019*

FEATURES:

* [grid] Make more validation about the `MsgCreateValidator` in CollectStdTxs
* [grid] Remove loosen token in stake pool, use bank to calculate the total loosen token
* [grid] Implement the block mint token-economics
* [grid] Add the service slash feature
* [grid] Redesign and implement the governance module to setup the new voting, tally, and penalty rules for each level of proposals
* [grid] Refactor and redefined all the gov/slashing/service/stake/distribution and gasPrice params
* [grid] Make gov data types codec wires usable across different protocol versions
* [grid] Don't export the unfinished proposals and refund the deposits of these proposals before export snapshot
* [grid] Refund service fee and deposit before export service state
* [grid] Add invariant checking level into makefile
* [grid] Only the genesis type profiler/trustee can initiate the addition or deletion (rather than prohibiting) transactions of the minor type profiler/trustee record. Everyone can view the profiler/trustee list
* [grid] Make sure the destination address is a trustee when the TaxUsage proposal execute
* [grid] Remove the record module
* [grid] Add `grid start --replay-last-block` to reset the app state by replay the last block
* [grid] Add `grid export --height` to export the snapshot of any block height even beyond the maximum cached historical version

* [gridcli] Add cli cmd to query the software upgrade signal status
* [gridcli] Make flag deposit not be required in the gov submit-proposal cmd
* [gridcli] Add token stats query cmd and lcd interface
* [gridcli] Replace decimal with int coins in distribution withdraw tags
* [gridcli] Add the sync tx broadcast type as the default mode in gridcli
* [gridcli] Add burn token cmd and lcd api
* [gridcli] Remove set-withdraw-addr sub-command

* [tendermint] Update tendermint to v0.27.3
* [test] Run cli test suite in parallel

BUG FIXES:

* Withdraw commission on self bond removal
* Use address instead of bond height / intratxcounter for deduplication
* Removal of mandatory self-delegation reward
* Fix bug of the tx result tags
* Fix absence proof verification
* Avoid to export account with no coin
* Correctly reset jailed-validator bond height / unbonding height on export-for-zero-height
* If a validator is jailed, distribute no reward to it
* Fix issue that miss checking the first one in Coins

## 0.9.1-patch01

*January 7th, 2019*

* Hotfix bug of software upgrade

## 0.9.1

*January 4th, 2019*

* Add cli cmd to query the software upgrade signal status
* Remove the text proposal

## 0.9.0

*December 27th, 2018*

* Refactor the gov types
* Make the deposit flag not be required in the gov submit-proposal cmd
* Add withdraw address into the withdraw tags list
* Fix the monitor bug

## 0.9.0-rc0

*December 19th, 2018*

BREAKING CHANGES:

* Use `gridtool` to replace the original `griddebug` and `gridmon`
* `grid init` must specify moniker

FEATURES:

* [gridcli] Optimize the way tags are displayed
* [gridcli] Add `gridcli stake delegations-to [validator-addr]` and `/stake/validators/{validatorAddr}/delegations` interfaces
* [grid] Application framework code refactoring
* [grid] Add a new mechanism to distribute service fee tax
* [grid] Slashing module supports querying slashing history
* [grid] Gov module adds TxTaxUsageProposal/SoftwareHaltProposal proposals
* [grid] Export and import blockchain snapshot at any block height
* [grid] Redesigned to implement class 2 software upgrade
* [grid] Restrict the block gas limit
* [grid] Improve tx search to support multiple tags
* [grid] Improve the default behavior of grid --home
* [grid] `grid tendermint show-address` output begins with `fca`
* [grid] Restrict the number of signatures on the transaction
* [grid] Add a check for the validator private key type and reject the unsupported private key type
* [tendermint] Update tendermint to v0.27.0

BUG FIXES:

* Add chain-id value checking for sign command
* Specify the required flags for cmds `query-proposal`, `query-deposit` and `query-vote`

## 0.8.0

*December 13th, 2018*

* Upgrade tendermint to v0.27.0-dev1

## 0.8.0-rc0

*December 3rd, 2018*

BREAKING CHANGES:

* Genesis.json supports any unit format of GRID CoinType
* The configuration information of the bech32 prefix is dynamically specified by the environment variable
* Improvement of File/directory path specification and the exception handler

FEATURES:

* Upgrade cosmos-sdk to v0.26.1-rc1 and remove the cosmos-sdk dependency
* Upgrade tendermint denpendency to v0.26.1-rc3
* View the current available withdraw balance by simulation mode
* Command line and LCD interface for service invocation request and query
* Implement guardian module for some governance proposal
* Added command add-genesis-account to configure account for genesis.json
* New proposal TerminatorProposal to terminate network consensus

## 0.7.0

*November 27th, 2018*

* Add broadcast command in bank
* Impose upgrade proposal with restrictions
* Fix bech32 prefix error in gridmon
* Improve user documents

## 0.7.0-rc0

*November 19th, 2018*

BREAKING CHANGES:

* [grid] New genesis workflow
* [grid] Validator.Owner renamed to Validator. Validator operator type has now changed to sdk.ValAddress
* [grid] unsafe_reset_all, show_validator, and show_node_id have been renamed to unsafe-reset-all, show-validator, and show-node-id
* [grid]Rename "revoked" to "jailed"
* [grid]Removed CompleteUnbonding and CompleteRedelegation Msg types, and instead added unbonding/redelegation queues to endblocker
* [grid]Removed slashing for governance non-voting validators
* [grid]Validators are no longer deleted until they can no longer possibly be slashed
* [grid]Remove ibc module
* [grid]Validator set updates delayed by one block
* [grid]Drop GenesisTx in favor of a signed StdTx with only one MsgCreateValidator message

FEATURES:

* Upgrade cosmos-sdk denpendency to v0.26.0
* Upgrade tendermint denpendency to v0.26.1-rc0
* [docs]Improve docs
* [grid]Add token inflation
* [grid]Add distribution module to distribute inflation token and collected transaction fee
* [gridcli] --from can now be either an address or a key name
* [gridcli] Passing --gas=simulate triggers a simulation of the tx before the actual execution. The gas estimate obtained via the simulation will be used as gas limit in the actual execution.
* [gridcli]Add --bech to gridcli keys show and respective REST endpoint to
* [gridcli]Introduced new commission flags for validator commands create-validator and edit-validator
* [gridcli]Add commands to query validator unbondings and redelegations
* [gridcli]Add rest apis and commands for distribution

BUG FIXES:

* [gridcli]Mark --to and --amount as required flags for gridcli bank send
* [grid]Add general merkle absence proof (also for empty substores)
* [grid]Fix issue about consumed gas increasing rapidly
* [grid]Return correct Tendermint validator update set on EndBlocker by not including non previously bonded validators that have zero power
* [grid]Add commission data to MsgCreateValidator signature bytes

## 0.6.0

*November 1st, 2018*

* Use --def-chain-id flag to reference the blockchain defined of the iService
* Fix some bugs about iservice definition and record
* Add cli and lcd test for record module
* Update the user doc of iservice definition and record

## 0.6.0-rc0

*October 24th, 2018*

BREAKING CHANGES:

* [monitor] Use new executable binary in monitor

FEATURES:

* [record] Add the record module of the data certification on blockchain
* [iservice] Add the feature of iService definition
* [cli] Add the example description in the cli help
* [test] Add Cli/LCD/Sim auto-test

BUG FIXES:

* Fix software upgrade issue caused by tx fee
* Report Panic when building the lcd proof
* Fix bugs in converting validator power to byte array
* Fix panic bug in wrong account number

## 0.5.0-rc1

*October 11th, 2018*

FEATURES:

* Make all the gov and upgrade parameters can be configured in the genesis.json

BUG FIXES

* Add check for iavl proof and value before building multistore proof

## 0.5.0-rc0

*September 30th, 2018*

BREAKING CHANGES:

* [cointype] Introduce the cointype of grid:
  * 1 grid = 10^18 grid-atto
  * 1 grid-milli = 10^15 grid-atto
  * 1 grid-micro = 10^12 grid-atto
  * 1 grid-nano = 10^9 grid-atto
  * 1 grid-pico = 10^6 grid-atto
  * 1 grid-femto = 10^3 grid-atto

FEATURES:

* [tendermint] Upgrade to Tendermint v0.23.1-rc0
* [cosmos-sdk] Upgrade to cosmos-sdk v0.24.2
  * Move the previous irisnet changeset about cosmos-sdk into gridiron
* [griddebug] Add griddebug tool
* [LCD/cli] Add the proof verification to the LCD and CLI
* [iparam] Support the modification of governance parameters of complex data type through governance, and the submission of modified proposals through json config files
* [software-upgrade] Software upgrade solutions of the irisnet

## 0.4.2

*September 22th, 2018*

BUG FIXES

* Fix consensus failure due to the double sign evidence be broadcasted before the genesis block

## 0.4.1

*September 12th, 2018*

BUG FIXES

* Missing to set validator intraTxCount in stake genesis init

## 0.4.0

*September 6th, 2018*

BREAKING CHANGES:

* [cosmos-sdk] Upgrade to cosmos-sdk v0.23.0
  * Change the address prefix format:
    * cosmosaccaddr --> faa
    * cosmosaccpub --> fap
    * cosmosvaladdr --> fva
    * cosmosvalpub --> fvp
  * Adjust the Route & rootMultiStore Commit for software upgrade
  * Must specify gas and fee in both command line and rest api
  * The fee should be grid token and the token amount should be no less than 2*(10^10)*gas

FEATURES:

* [tendermint] Upgrade to Tendermint v0.22.6
  * Store the pre-state to support the replay function
* [cosmos-sdk] Upgrade to cosmos-sdk v0.23.0
  * Add the paramProposal and softwareUpgradeProposal in gov module
  * Improve fee token mechanism to more reasonably deduct transaction fee and achieve more ability to defent DDOS attack.
  * Introduce the global parameter module

BUG FIXES

* Default account balance in genesis
* Fix grid version issue
* Fix the unit conflict issue in slashing
* Check the voting power when create validator
* Fix evidence amimo register issue

## 0.4.0-rc2

*Sep 5th, 2018*

BUG FIXES

* Fix evidence amimo register issue

## 0.4.0-rc1

*Aug 27th, 2018*

BUG FIXES

* Default account balance in genesis
* grid version issue
* Fix the unit conflict issue in slashing
* Check the voting power when create validator

## 0.3.0

*July 30th, 2018*

BREAKING CHANGES:

* [tendermint] Upgrade to Tendermint v0.22.2
  * Default ports changed from 466xx to 266xx
  * ED25519 addresses are the first 20-bytes of the SHA256 of the raw 32-byte pubkey (Instead of RIPEMD160)
* [cosmos-sdk] Upgrade to cosmos-sdk v0.22.0
* [monitor] Move `gridcli monitor` subcommand to `grid monitor`

FEATURES:

* [lcd] /tx/send is now the only endpoint for posing transaction to gridiron; aminofied all transaction messages
* [monitor] Improve the metrics for grid-monitor

BUG FIXES

* [cli] solve the issue of gridcli stake sign-info

##

## 0.2.0

*July 19th, 2018*

BREAKING CHANGES:

* [tendermint] Upgrade to Tendermint v0.21.0
* [cosmos-sdk] Upgrade to cosmos-sdk v0.19.1-rc1

FEATURES:

* [lcd] code refactor

* [cli] improve sendingand querying the  transactions

* [monitor]export data which is collected by Prometheus Server

  ​

##  
