# ChangeLog

## Current

### Features

### Fixes

### Chore

## 5.2.3-beta

### Features

- [#8382](https://github.com/blockscout/blockscout/pull/8382) - Add sitemap.xml
- [#8313](https://github.com/blockscout/blockscout/pull/8313) - Add batches to TokenInstance fetchers
- [#8285](https://github.com/blockscout/blockscout/pull/8285), [#8399](https://github.com/blockscout/blockscout/pull/8399) - Add CG/CMC coin price sources
- [#8181](https://github.com/blockscout/blockscout/pull/8181) - Insert current token balances placeholders along with historical
- [#8210](https://github.com/blockscout/blockscout/pull/8210) - Drop address foreign keys
- [#8292](https://github.com/blockscout/blockscout/pull/8292) - Add ETHEREUM_JSONRPC_WAIT_PER_TIMEOUT env var
- [#8269](https://github.com/blockscout/blockscout/pull/8269) - Don't push back to sequence on catchup exception
- [#8362](https://github.com/blockscout/blockscout/pull/8362), [#8398](https://github.com/blockscout/blockscout/pull/8398) - Drop token balances tokens foreign key

### Fixes

- [#8446](https://github.com/blockscout/blockscout/pull/8446) - Fix market cap calculation in case of CMC
- [#8431](https://github.com/blockscout/blockscout/pull/8431) - Fix contracts' output decoding
- [#8354](https://github.com/blockscout/blockscout/pull/8354) - Hotfix for proper addresses' tokens displaying
- [#8350](https://github.com/blockscout/blockscout/pull/8350) - Add Base Mainnet support for tx actions
- [#8282](https://github.com/blockscout/blockscout/pull/8282) - NFT fetcher improvements
- [#8287](https://github.com/blockscout/blockscout/pull/8287) - Add separate hackney pool for TokenInstance fetchers
- [#8293](https://github.com/blockscout/blockscout/pull/8293) - Add ETHEREUM_JSONRPC_TRACE_URL for Geth in docker-compose.yml
- [#8240](https://github.com/blockscout/blockscout/pull/8240) - Refactor and fix paging params in API v2
- [#8242](https://github.com/blockscout/blockscout/pull/8242) - Fixing visualizer service CORS issue when running docker-compose
- [#8355](https://github.com/blockscout/blockscout/pull/8355) - Fix current token balances redefining
- [#8338](https://github.com/blockscout/blockscout/pull/8338) - Fix reorgs query
- [#8413](https://github.com/blockscout/blockscout/pull/8413) - Put error in last call for STOP opcode
- [#8447](https://github.com/blockscout/blockscout/pull/8447) - Fix reorg transactions

### Chore

- [#8494](https://github.com/blockscout/blockscout/pull/8494) - Add release announcement in Slack
- [#8493](https://github.com/blockscout/blockscout/pull/8493) - Fix arm docker image build
- [#8478](https://github.com/blockscout/blockscout/pull/8478) - Set integration with Blockscout's eth bytecode DB endpoint by default and other enhancements
- [#8442](https://github.com/blockscout/blockscout/pull/8442) - Unify burn address definition
- [#8321](https://github.com/blockscout/blockscout/pull/8321) - Add curl into resulting Docker image
- [#8319](https://github.com/blockscout/blockscout/pull/8319) - Add MIX_ENV: 'prod' to docker-compose
- [#8281](https://github.com/blockscout/blockscout/pull/8281) - Planned removal of duplicate API endpoints: for CSV export and GraphQL

<details>
  <summary>Dependencies version bumps</summary>
- [#8244](https://github.com/blockscout/blockscout/pull/8244) - Bump core-js from 3.32.0 to 3.32.1 in /apps/block_scout_web/assets
- [#8243](https://github.com/blockscout/blockscout/pull/8243) - Bump sass from 1.65.1 to 1.66.0 in /apps/block_scout_web/assets
- [#8259](https://github.com/blockscout/blockscout/pull/8259) - Bump sweetalert2 from 11.7.23 to 11.7.27 in /apps/block_scout_web/assets
- [#8258](https://github.com/blockscout/blockscout/pull/8258) - Bump sass from 1.66.0 to 1.66.1 in /apps/block_scout_web/assets
- [#8260](https://github.com/blockscout/blockscout/pull/8260) - Bump jest from 29.6.2 to 29.6.3 in /apps/block_scout_web/assets
- [#8261](https://github.com/blockscout/blockscout/pull/8261) - Bump eslint-plugin-import from 2.28.0 to 2.28.1 in /apps/block_scout_web/assets
- [#8262](https://github.com/blockscout/blockscout/pull/8262) - Bump jest-environment-jsdom from 29.6.2 to 29.6.3 in /apps/block_scout_web/assets
- [#8275](https://github.com/blockscout/blockscout/pull/8275) - Bump ecto_sql from 3.10.1 to 3.10.2
- [#8284](https://github.com/blockscout/blockscout/pull/8284) - Bump luxon from 3.4.0 to 3.4.1 in /apps/block_scout_web/assets
- [#8294](https://github.com/blockscout/blockscout/pull/8294) - Bump chart.js from 4.3.3 to 4.4.0 in /apps/block_scout_web/assets
- [#8295](https://github.com/blockscout/blockscout/pull/8295) - Bump jest from 29.6.3 to 29.6.4 in /apps/block_scout_web/assets
- [#8296](https://github.com/blockscout/blockscout/pull/8296) - Bump jest-environment-jsdom from 29.6.3 to 29.6.4 in /apps/block_scout_web/assets
- [#8297](https://github.com/blockscout/blockscout/pull/8297) - Bump @babel/core from 7.22.10 to 7.22.11 in /apps/block_scout_web/assets
- [#8305](https://github.com/blockscout/blockscout/pull/8305) - Bump @amplitude/analytics-browser from 2.2.0 to 2.2.1 in /apps/block_scout_web/assets
- [#8342](https://github.com/blockscout/blockscout/pull/8342) - Bump postgrex from 0.17.2 to 0.17.3
- [#8341](https://github.com/blockscout/blockscout/pull/8341) - Bump hackney from 1.18.1 to 1.18.2
- [#8343](https://github.com/blockscout/blockscout/pull/8343) - Bump @amplitude/analytics-browser from 2.2.1 to 2.2.2 in /apps/block_scout_web/assets
- [#8344](https://github.com/blockscout/blockscout/pull/8344) - Bump postcss from 8.4.28 to 8.4.29 in /apps/block_scout_web/assets
- [#8330](https://github.com/blockscout/blockscout/pull/8330) - Bump bignumber.js from 9.1.1 to 9.1.2 in /apps/block_scout_web/assets
- [#8332](https://github.com/blockscout/blockscout/pull/8332) - Bump jquery from 3.7.0 to 3.7.1 in /apps/block_scout_web/assets
- [#8329](https://github.com/blockscout/blockscout/pull/8329) - Bump viewerjs from 1.11.4 to 1.11.5 in /apps/block_scout_web/assets
- [#8328](https://github.com/blockscout/blockscout/pull/8328) - Bump eslint from 8.47.0 to 8.48.0 in /apps/block_scout_web/assets
- [#8325](https://github.com/blockscout/blockscout/pull/8325) - Bump exvcr from 0.14.3 to 0.14.4
- [#8323](https://github.com/blockscout/blockscout/pull/8323) - Bump ex_doc from 0.30.5 to 0.30.6
- [#8322](https://github.com/blockscout/blockscout/pull/8322) - Bump dialyxir from 1.3.0 to 1.4.0
- [#8326](https://github.com/blockscout/blockscout/pull/8326) - Bump comeonin from 5.3.3 to 5.4.0
- [#8331](https://github.com/blockscout/blockscout/pull/8331) - Bump luxon from 3.4.1 to 3.4.2 in /apps/block_scout_web/assets
- [#8324](https://github.com/blockscout/blockscout/pull/8324) - Bump spandex_datadog from 1.3.0 to 1.4.0
- [#8327](https://github.com/blockscout/blockscout/pull/8327) - Bump bcrypt_elixir from 3.0.1 to 3.1.0
- [#8358](https://github.com/blockscout/blockscout/pull/8358) - Bump @babel/preset-env from 7.22.10 to 7.22.14 in /apps/block_scout_web/assets
- [#8365](https://github.com/blockscout/blockscout/pull/8365) - Bump dialyxir from 1.4.0 to 1.4.1
- [#8374](https://github.com/blockscout/blockscout/pull/8374) - Bump @amplitude/analytics-browser from 2.2.2 to 2.2.3 in /apps/block_scout_web/assets
- [#8373](https://github.com/blockscout/blockscout/pull/8373) - Bump ex_secp256k1 from 0.7.0 to 0.7.1
- [#8391](https://github.com/blockscout/blockscout/pull/8391) - Bump @babel/preset-env from 7.22.14 to 7.22.15 in /apps/block_scout_web/assets
- [#8390](https://github.com/blockscout/blockscout/pull/8390) - Bump photoswipe from 5.3.8 to 5.3.9 in /apps/block_scout_web/assets
- [#8389](https://github.com/blockscout/blockscout/pull/8389) - Bump @babel/core from 7.22.11 to 7.22.15 in /apps/block_scout_web/assets
- [#8392](https://github.com/blockscout/blockscout/pull/8392) - Bump ex_cldr_numbers from 2.31.3 to 2.32.0
- [#8400](https://github.com/blockscout/blockscout/pull/8400) - Bump ex_secp256k1 from 0.7.1 to 0.7.2
- [#8405](https://github.com/blockscout/blockscout/pull/8405) - Bump luxon from 3.4.2 to 3.4.3 in /apps/block_scout_web/assets
- [#8404](https://github.com/blockscout/blockscout/pull/8404) - Bump ex_abi from 0.6.0 to 0.6.1
- [#8410](https://github.com/blockscout/blockscout/pull/8410) - Bump core-js from 3.32.1 to 3.32.2 in /apps/block_scout_web/assets
- [#8418](https://github.com/blockscout/blockscout/pull/8418) - Bump url from 0.11.1 to 0.11.2 in /apps/block_scout_web/assets
- [#8416](https://github.com/blockscout/blockscout/pull/8416) - Bump @babel/core from 7.22.15 to 7.22.17 in /apps/block_scout_web/assets
- [#8419](https://github.com/blockscout/blockscout/pull/8419) - Bump assert from 2.0.0 to 2.1.0 in /apps/block_scout_web/assets
- [#8417](https://github.com/blockscout/blockscout/pull/8417) - Bump photoswipe from 5.3.9 to 5.4.0 in /apps/block_scout_web/assets
- [#8441](https://github.com/blockscout/blockscout/pull/8441) - Bump eslint from 8.48.0 to 8.49.0 in /apps/block_scout_web/assets
- [#8439](https://github.com/blockscout/blockscout/pull/8439) - Bump ex_cldr_numbers from 2.32.0 to 2.32.1
- [#8444](https://github.com/blockscout/blockscout/pull/8444) - Bump ex_cldr_numbers from 2.32.1 to 2.32.2
- [#8445](https://github.com/blockscout/blockscout/pull/8445) - Bump ex_abi from 0.6.1 to 0.6.2
- [#8450](https://github.com/blockscout/blockscout/pull/8450) - Bump jest-environment-jsdom from 29.6.4 to 29.7.0 in /apps/block_scout_web/assets
- [#8451](https://github.com/blockscout/blockscout/pull/8451) - Bump jest from 29.6.4 to 29.7.0 in /apps/block_scout_web/assets
- [#8463](https://github.com/blockscout/blockscout/pull/8463) - Bump sass from 1.66.1 to 1.67.0 in /apps/block_scout_web/assets
- [#8464](https://github.com/blockscout/blockscout/pull/8464) - Bump @babel/core from 7.22.17 to 7.22.19 in /apps/block_scout_web/assets
- [#8462](https://github.com/blockscout/blockscout/pull/8462) - Bump sweetalert2 from 11.7.27 to 11.7.28 in /apps/block_scout_web/assets
- [#8479](https://github.com/blockscout/blockscout/pull/8479) - Bump photoswipe from 5.4.0 to 5.4.1 in /apps/block_scout_web/assets
- [#8483](https://github.com/blockscout/blockscout/pull/8483) - Bump @amplitude/analytics-browser from 2.2.3 to 2.3.1 in /apps/block_scout_web/assets
- [#8481](https://github.com/blockscout/blockscout/pull/8481) - Bump @babel/preset-env from 7.22.15 to 7.22.20 in /apps/block_scout_web/assets
- [#8480](https://github.com/blockscout/blockscout/pull/8480) - Bump @babel/core from 7.22.19 to 7.22.20 in /apps/block_scout_web/assets
- [#8482](https://github.com/blockscout/blockscout/pull/8482) - Bump viewerjs from 1.11.5 to 1.11.6 in /apps/block_scout_web/assets
- [#8489](https://github.com/blockscout/blockscout/pull/8489) - Bump postcss from 8.4.29 to 8.4.30 in /apps/block_scout_web/assets
</details>

## 5.2.2-beta

### Features

- [#8218](https://github.com/blockscout/blockscout/pull/8218) - Add `/api/v2/search/quick` method
- [#8202](https://github.com/blockscout/blockscout/pull/8202) - Add `/api/v2/addresses/:address_hash/tabs-counters` endpoint
- [#8156](https://github.com/blockscout/blockscout/pull/8156) - Add `is_verified_via_admin_panel` property to tokens table
- [#8165](https://github.com/blockscout/blockscout/pull/8165), [#8201](https://github.com/blockscout/blockscout/pull/8201) - Add broadcast of updated address_current_token_balances
- [#7952](https://github.com/blockscout/blockscout/pull/7952) - Add parsing constructor arguments for sourcify contracts
- [#6190](https://github.com/blockscout/blockscout/pull/6190) - Add EIP-1559 support to gas price oracle
- [#7977](https://github.com/blockscout/blockscout/pull/7977) - GraphQL: extend schema with new field for existing objects
- [#8158](https://github.com/blockscout/blockscout/pull/8158), [#8164](https://github.com/blockscout/blockscout/pull/8164) - Include unfetched balances in TokenBalanceOnDemand fetcher

### Fixes

- [#8233](https://github.com/blockscout/blockscout/pull/8233) - Fix API v2 broken tx response
- [#8147](https://github.com/blockscout/blockscout/pull/8147) - Switch sourcify tests from POA Sokol to Gnosis Chiado
- [#8145](https://github.com/blockscout/blockscout/pull/8145) - Handle negative holders count in API v2
- [#8040](https://github.com/blockscout/blockscout/pull/8040) - Resolve issue with Docker image for Mac M1/M2
- [#8060](https://github.com/blockscout/blockscout/pull/8060) - Fix eth_getLogs API endpoint
- [#8082](https://github.com/blockscout/blockscout/pull/8082), [#8088](https://github.com/blockscout/blockscout/pull/8088) - Fix Rootstock charts API
- [#7992](https://github.com/blockscout/blockscout/pull/7992) - Fix missing range insert
- [#8022](https://github.com/blockscout/blockscout/pull/8022) - Don't add reorg block number to missing blocks

### Chore

- [#8222](https://github.com/blockscout/blockscout/pull/8222) - docker-compose for new UI with external backend
- [#8177](https://github.com/blockscout/blockscout/pull/8177) - Refactor address counter functions
- [#8183](https://github.com/blockscout/blockscout/pull/8183) - Update frontend envs in order to pass their validation
- [#8167](https://github.com/blockscout/blockscout/pull/8167) - Manage concurrency for Token and TokenBalance fetcher
- [#8179](https://github.com/blockscout/blockscout/pull/8179) - Enhance nginx config
- [#8146](https://github.com/blockscout/blockscout/pull/8146) - Add method_id to write methods in API v2 response
- [#8105](https://github.com/blockscout/blockscout/pull/8105) - Extend API v1 with endpoints used by new UI
- [#8104](https://github.com/blockscout/blockscout/pull/8104) - remove "TODO" from API v2 response
- [#8100](https://github.com/blockscout/blockscout/pull/8100), [#8103](https://github.com/blockscout/blockscout/pull/8103) - Extend docker-compose configs with new config when front is running externally
- [#8012](https://github.com/blockscout/blockscout/pull/8012) - API v2 smart-contract verification extended logging

<details>
  <summary>Dependencies version bumps</summary>

- [#7980](https://github.com/blockscout/blockscout/pull/7980) - Bump solc from 0.8.20 to 0.8.21 in /apps/explorer
- [#7986](https://github.com/blockscout/blockscout/pull/7986) - Bump sass from 1.63.6 to 1.64.0 in /apps/block_scout_web/assets
- [#8030](https://github.com/blockscout/blockscout/pull/8030) - Bump sweetalert2 from 11.7.18 to 11.7.20 in /apps/block_scout_web/assets
- [#8029](https://github.com/blockscout/blockscout/pull/8029) - Bump viewerjs from 1.11.3 to 1.11.4 in /apps/block_scout_web/assets
- [#8028](https://github.com/blockscout/blockscout/pull/8028) - Bump sass from 1.64.0 to 1.64.1 in /apps/block_scout_web/assets
- [#8026](https://github.com/blockscout/blockscout/pull/8026) - Bump dataloader from 1.0.10 to 1.0.11
- [#8036](https://github.com/blockscout/blockscout/pull/8036) - Bump ex_cldr_numbers from 2.31.1 to 2.31.3
- [#8027](https://github.com/blockscout/blockscout/pull/8027) - Bump absinthe from 1.7.4 to 1.7.5
- [#8035](https://github.com/blockscout/blockscout/pull/8035) - Bump wallaby from 0.30.4 to 0.30.5
- [#8038](https://github.com/blockscout/blockscout/pull/8038) - Bump chart.js from 4.3.0 to 4.3.1 in /apps/block_scout_web/assets
- [#8047](https://github.com/blockscout/blockscout/pull/8047) - Bump chart.js from 4.3.1 to 4.3.2 in /apps/block_scout_web/assets
- [#8000](https://github.com/blockscout/blockscout/pull/8000) - Bump postcss from 8.4.26 to 8.4.27 in /apps/block_scout_web/assets
- [#8052](https://github.com/blockscout/blockscout/pull/8052) - Bump @amplitude/analytics-browser from 2.1.2 to 2.1.3 in /apps/block_scout_web/assets
- [#8054](https://github.com/blockscout/blockscout/pull/8054) - Bump jest-environment-jsdom from 29.6.1 to 29.6.2 in /apps/block_scout_web/assets
- [#8063](https://github.com/blockscout/blockscout/pull/8063) - Bump eslint from 8.45.0 to 8.46.0 in /apps/block_scout_web/assets
- [#8066](https://github.com/blockscout/blockscout/pull/8066) - Bump ex_json_schema from 0.9.3 to 0.10.1
- [#8064](https://github.com/blockscout/blockscout/pull/8064) - Bump core-js from 3.31.1 to 3.32.0 in /apps/block_scout_web/assets
- [#8053](https://github.com/blockscout/blockscout/pull/8053) - Bump jest from 29.6.1 to 29.6.2 in /apps/block_scout_web/assets
- [#8065](https://github.com/blockscout/blockscout/pull/8065) - Bump eslint-plugin-import from 2.27.5 to 2.28.0 in /apps/block_scout_web/assets
- [#8092](https://github.com/blockscout/blockscout/pull/8092) - Bump exvcr from 0.14.1 to 0.14.2
- [#8091](https://github.com/blockscout/blockscout/pull/8091) - Bump sass from 1.64.1 to 1.64.2 in /apps/block_scout_web/assets
- [#8114](https://github.com/blockscout/blockscout/pull/8114) - Bump ex_doc from 0.30.3 to 0.30.4
- [#8115](https://github.com/blockscout/blockscout/pull/8115) - Bump chart.js from 4.3.2 to 4.3.3 in /apps/block_scout_web/assets
- [#8116](https://github.com/blockscout/blockscout/pull/8116) - Bump @fortawesome/fontawesome-free from 6.4.0 to 6.4.2 in /apps/block_scout_web/assets
- [#8142](https://github.com/blockscout/blockscout/pull/8142) - Bump sobelow from 0.12.2 to 0.13.0
- [#8141](https://github.com/blockscout/blockscout/pull/8141) - Bump @babel/core from 7.22.9 to 7.22.10 in /apps/block_scout_web/assets
- [#8140](https://github.com/blockscout/blockscout/pull/8140) - Bump @babel/preset-env from 7.22.9 to 7.22.10 in /apps/block_scout_web/assets
- [#8160](https://github.com/blockscout/blockscout/pull/8160) - Bump exvcr from 0.14.2 to 0.14.3
- [#8159](https://github.com/blockscout/blockscout/pull/8159) - Bump luxon from 3.3.0 to 3.4.0 in /apps/block_scout_web/assets
- [#8169](https://github.com/blockscout/blockscout/pull/8169) - Bump sass from 1.64.2 to 1.65.1 in /apps/block_scout_web/assets
- [#8170](https://github.com/blockscout/blockscout/pull/8170) - Bump sweetalert2 from 11.7.20 to 11.7.22 in /apps/block_scout_web/assets
- [#8188](https://github.com/blockscout/blockscout/pull/8188) - Bump eslint from 8.46.0 to 8.47.0 in /apps/block_scout_web/assets
- [#8204](https://github.com/blockscout/blockscout/pull/8204) - Bump ex_doc from 0.30.4 to 0.30.5
- [#8207](https://github.com/blockscout/blockscout/pull/8207) - Bump wallaby from 0.30.5 to 0.30.6
- [#8212](https://github.com/blockscout/blockscout/pull/8212) - Bump sweetalert2 from 11.7.22 to 11.7.23 in /apps/block_scout_web/assets
- [#8203](https://github.com/blockscout/blockscout/pull/8203) - Bump autoprefixer from 10.4.14 to 10.4.15 in /apps/block_scout_web/assets
- [#8214](https://github.com/blockscout/blockscout/pull/8214) - Bump @amplitude/analytics-browser from 2.1.3 to 2.2.0 in /apps/block_scout_web/assets
- [#8225](https://github.com/blockscout/blockscout/pull/8225) - Bump postcss from 8.4.27 to 8.4.28 in /apps/block_scout_web/assets
- [#8224](https://github.com/blockscout/blockscout/pull/8224) - Bump gettext from 0.22.3 to 0.23.1

</details>

## 5.2.1-beta

### Features

- [#7970](https://github.com/blockscout/blockscout/pull/7970) - Search improvements: add sorting
- [#7771](https://github.com/blockscout/blockscout/pull/7771) - CSV export: speed up
- [#7962](https://github.com/blockscout/blockscout/pull/7962) - Allow indicate CMC id of the coin through env var
- [#7946](https://github.com/blockscout/blockscout/pull/7946) - API v2 rate limit: Put token to cookies & change /api/v2/key method
- [#7888](https://github.com/blockscout/blockscout/pull/7888) - Add token balances info to watchlist address response
- [#7898](https://github.com/blockscout/blockscout/pull/7898) - Add possibility to add extra headers with JSON RPC URL
- [#7836](https://github.com/blockscout/blockscout/pull/7836) - Improve unverified email flow
- [#7784](https://github.com/blockscout/blockscout/pull/7784) - Search improvements: Add new fields, light refactoring
- [#7811](https://github.com/blockscout/blockscout/pull/7811) - Filter addresses before insertion
- [#7895](https://github.com/blockscout/blockscout/pull/7895) - API v2: Add sorting to tokens page
- [#7859](https://github.com/blockscout/blockscout/pull/7859) - Add TokenTotalSupplyUpdater
- [#7873](https://github.com/blockscout/blockscout/pull/7873) - Chunk realtime balances requests
- [#7927](https://github.com/blockscout/blockscout/pull/7927) - Delete token balances only for blocks that lost consensus
- [#7947](https://github.com/blockscout/blockscout/pull/7947) - Improve locks acquiring

### Fixes

- [#8187](https://github.com/blockscout/blockscout/pull/8187) - API v1 500 error convert to 404, if requested path is incorrect
- [#7852](https://github.com/blockscout/blockscout/pull/7852) - Token balances refactoring & fixes
- [#7872](https://github.com/blockscout/blockscout/pull/7872) - Fix pending gas price in pending tx
- [#7875](https://github.com/blockscout/blockscout/pull/7875) - Fix twin compiler version
- [#7825](https://github.com/blockscout/blockscout/pull/7825) - Fix nginx config for the new frontend websockets
- [#7772](https://github.com/blockscout/blockscout/pull/7772) - Fix parsing of database password period(s)
- [#7803](https://github.com/blockscout/blockscout/pull/7803) - Fix additional sources and interfaces, save names for vyper contracts
- [#7758](https://github.com/blockscout/blockscout/pull/7758) - Remove limit for configurable fetchers
- [#7764](https://github.com/blockscout/blockscout/pull/7764) - Fix missing ranges insertion and deletion logic
- [#7843](https://github.com/blockscout/blockscout/pull/7843) - Fix created_contract_code_indexed_at updating
- [#7855](https://github.com/blockscout/blockscout/pull/7855) - Handle internal transactions unique_violation
- [#7899](https://github.com/blockscout/blockscout/pull/7899) - Fix catchup numbers_to_ranges function
- [#7951](https://github.com/blockscout/blockscout/pull/7951) - Fix TX url in email notifications on mainnet

### Chore

- [#7963](https://github.com/blockscout/blockscout/pull/7963) - Op Stack: ignore depositNonce
- [#7954](https://github.com/blockscout/blockscout/pull/7954) - Enhance Account Explorer.Account.Notifier.Email module tests
- [#7950](https://github.com/blockscout/blockscout/pull/7950) - Add GA CI for Eth Goerli chain
- [#7934](https://github.com/blockscout/blockscout/pull/7934), [#7936](https://github.com/blockscout/blockscout/pull/7936) - Explicitly set consensus == true in queries (convenient for search), remove logger requirements, where it is not used anymore
- [#7901](https://github.com/blockscout/blockscout/pull/7901) - Fix Docker image build
- [#7890](https://github.com/blockscout/blockscout/pull/7890), [#7918](https://github.com/blockscout/blockscout/pull/7918) - Resolve warning: Application.get_env/2 is discouraged in the module body, use Application.compile_env/3 instead
- [#7863](https://github.com/blockscout/blockscout/pull/7863) - Add max_age for account sessions
- [#7841](https://github.com/blockscout/blockscout/pull/7841) - CORS setup for docker-compose config with new frontend
- [#7832](https://github.com/blockscout/blockscout/pull/7832), [#7891](https://github.com/blockscout/blockscout/pull/7891) - API v2: Add block_number, block_hash to logs
- [#7789](https://github.com/blockscout/blockscout/pull/7789) - Fix test warnings; Fix name of `MICROSERVICE_ETH_BYTECODE_DB_INTERVAL_BETWEEN_LOOKUPS` env variable
- [#7819](https://github.com/blockscout/blockscout/pull/7819) - Add logging for unknown error verification result
- [#7781](https://github.com/blockscout/blockscout/pull/7781) - Add `/api/v1/health/liveness` and `/api/v1/health/readiness`

<details>
  <summary>Dependencies version bumps</summary>

- [#7759](https://github.com/blockscout/blockscout/pull/7759) - Bump sass from 1.63.4 to 1.63.5 in /apps/block_scout_web/assets
- [#7760](https://github.com/blockscout/blockscout/pull/7760) - Bump @amplitude/analytics-browser from 2.0.0 to 2.0.1 in /apps/block_scout_web/assets
- [#7762](https://github.com/blockscout/blockscout/pull/7762) - Bump webpack from 5.87.0 to 5.88.0 in /apps/block_scout_web/assets
- [#7769](https://github.com/blockscout/blockscout/pull/7769) - Bump sass from 1.63.5 to 1.63.6 in /apps/block_scout_web/assets
- [#7805](https://github.com/blockscout/blockscout/pull/7805) - Bump ssl_verify_fun from 1.1.6 to 1.1.7
- [#7812](https://github.com/blockscout/blockscout/pull/7812) - Bump webpack from 5.88.0 to 5.88.1 in /apps/block_scout_web/assets
- [#7770](https://github.com/blockscout/blockscout/pull/7770) - Bump @amplitude/analytics-browser from 2.0.1 to 2.1.0 in /apps/block_scout_web/assets
- [#7821](https://github.com/blockscout/blockscout/pull/7821) - Bump absinthe from 1.7.1 to 1.7.3
- [#7823](https://github.com/blockscout/blockscout/pull/7823) - Bump @amplitude/analytics-browser from 2.1.0 to 2.1.1 in /apps/block_scout_web/assets
- [#7838](https://github.com/blockscout/blockscout/pull/7838) - Bump gettext from 0.22.2 to 0.22.3
- [#7840](https://github.com/blockscout/blockscout/pull/7840) - Bump eslint from 8.43.0 to 8.44.0 in /apps/block_scout_web/assets
- [#7839](https://github.com/blockscout/blockscout/pull/7839) - Bump photoswipe from 5.3.7 to 5.3.8 in /apps/block_scout_web/assets
- [#7850](https://github.com/blockscout/blockscout/pull/7850) - Bump jest-environment-jsdom from 29.5.0 to 29.6.0 in /apps/block_scout_web/assets
- [#7848](https://github.com/blockscout/blockscout/pull/7848) - Bump @amplitude/analytics-browser from 2.1.1 to 2.1.2 in /apps/block_scout_web/assets
- [#7847](https://github.com/blockscout/blockscout/pull/7847) - Bump @babel/core from 7.22.5 to 7.22.6 in /apps/block_scout_web/assets
- [#7846](https://github.com/blockscout/blockscout/pull/7846) - Bump @babel/preset-env from 7.22.5 to 7.22.6 in /apps/block_scout_web/assets
- [#7856](https://github.com/blockscout/blockscout/pull/7856) - Bump ex_cldr from 2.37.1 to 2.37.2
- [#7870](https://github.com/blockscout/blockscout/pull/7870) - Bump jest from 29.5.0 to 29.6.1 in /apps/block_scout_web/assets
- [#7867](https://github.com/blockscout/blockscout/pull/7867) - Bump postcss from 8.4.24 to 8.4.25 in /apps/block_scout_web/assets
- [#7871](https://github.com/blockscout/blockscout/pull/7871) - Bump @babel/core from 7.22.6 to 7.22.8 in /apps/block_scout_web/assets
- [#7868](https://github.com/blockscout/blockscout/pull/7868) - Bump jest-environment-jsdom from 29.6.0 to 29.6.1 in /apps/block_scout_web/assets
- [#7866](https://github.com/blockscout/blockscout/pull/7866) - Bump @babel/preset-env from 7.22.6 to 7.22.7 in /apps/block_scout_web/assets
- [#7869](https://github.com/blockscout/blockscout/pull/7869) - Bump core-js from 3.31.0 to 3.31.1 in /apps/block_scout_web/assets
- [#7884](https://github.com/blockscout/blockscout/pull/7884) - Bump ecto from 3.10.2 to 3.10.3
- [#7882](https://github.com/blockscout/blockscout/pull/7882) - Bump jason from 1.4.0 to 1.4.1
- [#7880](https://github.com/blockscout/blockscout/pull/7880) - Bump absinthe from 1.7.3 to 1.7.4
- [#7879](https://github.com/blockscout/blockscout/pull/7879) - Bump babel-loader from 9.1.2 to 9.1.3 in /apps/block_scout_web/assets
- [#7881](https://github.com/blockscout/blockscout/pull/7881) - Bump ex_cldr_numbers from 2.31.1 to 2.31.2
- [#7883](https://github.com/blockscout/blockscout/pull/7883) - Bump ex_doc from 0.29.4 to 0.30.1
- [#7916](https://github.com/blockscout/blockscout/pull/7916) - Bump semver from 5.7.1 to 5.7.2 in /apps/explorer
- [#7912](https://github.com/blockscout/blockscout/pull/7912) - Bump sweetalert2 from 11.7.12 to 11.7.16 in /apps/block_scout_web/assets
- [#7913](https://github.com/blockscout/blockscout/pull/7913) - Bump ex_doc from 0.30.1 to 0.30.2
- [#7923](https://github.com/blockscout/blockscout/pull/7923) - Bump postgrex from 0.17.1 to 0.17.2
- [#7921](https://github.com/blockscout/blockscout/pull/7921) - Bump @babel/preset-env from 7.22.7 to 7.22.9 in /apps/block_scout_web/assets
- [#7922](https://github.com/blockscout/blockscout/pull/7922) - Bump @babel/core from 7.22.8 to 7.22.9 in /apps/block_scout_web/assets
- [#7931](https://github.com/blockscout/blockscout/pull/7931) - Bump wallaby from 0.30.3 to 0.30.4
- [#7940](https://github.com/blockscout/blockscout/pull/7940) - Bump postcss from 8.4.25 to 8.4.26 in /apps/block_scout_web/assets
- [#7939](https://github.com/blockscout/blockscout/pull/7939) - Bump eslint from 8.44.0 to 8.45.0 in /apps/block_scout_web/assets
- [#7955](https://github.com/blockscout/blockscout/pull/7955) - Bump sweetalert2 from 11.7.16 to 11.7.18 in /apps/block_scout_web/assets
- [#7958](https://github.com/blockscout/blockscout/pull/7958) - Bump ex_doc from 0.30.2 to 0.30.3
- [#7965](https://github.com/blockscout/blockscout/pull/7965) - Bump webpack from 5.88.1 to 5.88.2 in /apps/block_scout_web/assets
- [#7972](https://github.com/blockscout/blockscout/pull/7972) - Bump word-wrap from 1.2.3 to 1.2.4 in /apps/block_scout_web/assets
</details>

## 5.2.0-beta

### Features

- [#7502](https://github.com/blockscout/blockscout/pull/7502) - Improve performance of some methods, endpoints and SQL queries
- [#7665](https://github.com/blockscout/blockscout/pull/7665) - Add standard-json vyper verification
- [#7685](https://github.com/blockscout/blockscout/pull/7685) - Add yul filter and "language" field for smart contracts
- [#7653](https://github.com/blockscout/blockscout/pull/7653) - Add support for DEPOSIT and WITHDRAW token transfer event in older contracts
- [#7628](https://github.com/blockscout/blockscout/pull/7628) - Support partially verified property from verifier MS; Add property to track contracts automatically verified via eth-bytecode-db
- [#7603](https://github.com/blockscout/blockscout/pull/7603) - Add Polygon Edge and optimism genesis files support
- [#7585](https://github.com/blockscout/blockscout/pull/7585) - Store and display native coin market cap from the DB
- [#7513](https://github.com/blockscout/blockscout/pull/7513) - Add Polygon Edge support
- [#7532](https://github.com/blockscout/blockscout/pull/7532) - Handle empty id in json rpc responses
- [#7544](https://github.com/blockscout/blockscout/pull/7544) - Add ERC-1155 signatures to uncataloged_token_transfer_block_numbers
- [#7363](https://github.com/blockscout/blockscout/pull/7363) - CSV export filters
- [#7697](https://github.com/blockscout/blockscout/pull/7697) - Limit fetchers init tasks

### Fixes

- [#7712](https://github.com/blockscout/blockscout/pull/7712) - Transaction actions import fix
- [#7709](https://github.com/blockscout/blockscout/pull/7709) - Contract args displaying bug
- [#7654](https://github.com/blockscout/blockscout/pull/7654) - Optimize exchange rates requests rate
- [#7636](https://github.com/blockscout/blockscout/pull/7636) - Remove receive from read methods
- [#7635](https://github.com/blockscout/blockscout/pull/7635) - Fix single 1155 transfer displaying
- [#7629](https://github.com/blockscout/blockscout/pull/7629) - Fix NFT fetcher
- [#7614](https://github.com/blockscout/blockscout/pull/7614) - API and smart-contracts fixes and improvements
- [#7611](https://github.com/blockscout/blockscout/pull/7611) - Fix tokens pagination
- [#7566](https://github.com/blockscout/blockscout/pull/7566) - Account: check composed email before sending
- [#7564](https://github.com/blockscout/blockscout/pull/7564) - Return contract type in address view
- [#7562](https://github.com/blockscout/blockscout/pull/7562) - Remove fallback from Read methods
- [#7537](https://github.com/blockscout/blockscout/pull/7537), [#7553](https://github.com/blockscout/blockscout/pull/7553) - Withdrawals fixes and improvements
- [#7546](https://github.com/blockscout/blockscout/pull/7546) - API v2: fix today coin price (use in-memory or cached in DB value)
- [#7545](https://github.com/blockscout/blockscout/pull/7545) - API v2: Check if cached exchange rate is empty before replacing DB value in stats API
- [#7516](https://github.com/blockscout/blockscout/pull/7516) - Fix shrinking logo in Safari
- [#7590](https://github.com/blockscout/blockscout/pull/7590) - Drop genesis block in internal transactions fetcher
- [#7639](https://github.com/blockscout/blockscout/pull/7639) - Fix contract creation transactions
- [#7724](https://github.com/blockscout/blockscout/pull/7724), [#7753](https://github.com/blockscout/blockscout/pull/7753) - Move MissingRangesCollector init logic to handle_continue
- [#7751](https://github.com/blockscout/blockscout/pull/7751) - Add missing method_to_url params for trace transactions

### Chore

- [#7699](https://github.com/blockscout/blockscout/pull/7699) - Add block_number index for address_coin_balances table
- [#7666](https://github.com/blockscout/blockscout/pull/7666), [#7740](https://github.com/blockscout/blockscout/pull/7740), [#7741](https://github.com/blockscout/blockscout/pull/7741) - Search label query
- [#7644](https://github.com/blockscout/blockscout/pull/7644) - Publish docker images CI for prod/staging branches
- [#7594](https://github.com/blockscout/blockscout/pull/7594) - Stats service support in docker-compose config with new frontend
- [#7576](https://github.com/blockscout/blockscout/pull/7576) - Check left blocks in pending block operations in order to decide, if we need to display indexing int tx banner at the top
- [#7543](https://github.com/blockscout/blockscout/pull/7543) - Allow hyphen in DB username

<details>
  <summary>Dependencies version bumps</summary>

- [#7518](https://github.com/blockscout/blockscout/pull/7518) - Bump mini-css-extract-plugin from 2.7.5 to 2.7.6 in /apps/block_scout_web/assets
- [#7519](https://github.com/blockscout/blockscout/pull/7519) - Bump style-loader from 3.3.2 to 3.3.3 in /apps/block_scout_web/assets
- [#7505](https://github.com/blockscout/blockscout/pull/7505) - Bump webpack from 5.83.0 to 5.83.1 in /apps/block_scout_web/assets
- [#7533](https://github.com/blockscout/blockscout/pull/7533) - Bump sass-loader from 13.2.2 to 13.3.0 in /apps/block_scout_web/assets
- [#7534](https://github.com/blockscout/blockscout/pull/7534) - Bump eslint from 8.40.0 to 8.41.0 in /apps/block_scout_web/assets
- [#7541](https://github.com/blockscout/blockscout/pull/7541) - Bump cldr_utils from 2.23.1 to 2.24.0
- [#7542](https://github.com/blockscout/blockscout/pull/7542) - Bump ex_cldr_units from 3.16.0 to 3.16.1
- [#7548](https://github.com/blockscout/blockscout/pull/7548) - Bump briefly from 20d1318 to 678a376
- [#7547](https://github.com/blockscout/blockscout/pull/7547) - Bump webpack from 5.83.1 to 5.84.0 in /apps/block_scout_web/assets
- [#7554](https://github.com/blockscout/blockscout/pull/7554) - Bump webpack from 5.84.0 to 5.84.1 in /apps/block_scout_web/assets
- [#7568](https://github.com/blockscout/blockscout/pull/7568) - Bump @babel/core from 7.21.8 to 7.22.1 in /apps/block_scout_web/assets
- [#7569](https://github.com/blockscout/blockscout/pull/7569) - Bump postcss-loader from 7.3.0 to 7.3.1 in /apps/block_scout_web/assets
- [#7570](https://github.com/blockscout/blockscout/pull/7570) - Bump number from 1.0.3 to 1.0.4
- [#7567](https://github.com/blockscout/blockscout/pull/7567) - Bump @babel/preset-env from 7.21.5 to 7.22.2 in /apps/block_scout_web/assets
- [#7582](https://github.com/blockscout/blockscout/pull/7582) - Bump eslint-config-standard from 17.0.0 to 17.1.0 in /apps/block_scout_web/assets
- [#7581](https://github.com/blockscout/blockscout/pull/7581) - Bump sass-loader from 13.3.0 to 13.3.1 in /apps/block_scout_web/assets
- [#7578](https://github.com/blockscout/blockscout/pull/7578) - Bump @babel/preset-env from 7.22.2 to 7.22.4 in /apps/block_scout_web/assets
- [#7577](https://github.com/blockscout/blockscout/pull/7577) - Bump postcss-loader from 7.3.1 to 7.3.2 in /apps/block_scout_web/assets
- [#7579](https://github.com/blockscout/blockscout/pull/7579) - Bump sweetalert2 from 11.7.5 to 11.7.8 in /apps/block_scout_web/assets
- [#7591](https://github.com/blockscout/blockscout/pull/7591) - Bump sweetalert2 from 11.7.8 to 11.7.9 in /apps/block_scout_web/assets
- [#7593](https://github.com/blockscout/blockscout/pull/7593) - Bump ex_json_schema from 0.9.2 to 0.9.3
- [#7580](https://github.com/blockscout/blockscout/pull/7580) - Bump postcss from 8.4.23 to 8.4.24 in /apps/block_scout_web/assets
- [#7601](https://github.com/blockscout/blockscout/pull/7601) - Bump sweetalert2 from 11.7.9 to 11.7.10 in /apps/block_scout_web/assets
- [#7602](https://github.com/blockscout/blockscout/pull/7602) - Bump mime from 2.0.3 to 2.0.4
- [#7618](https://github.com/blockscout/blockscout/pull/7618) - Bump gettext from 0.22.1 to 0.22.2
- [#7617](https://github.com/blockscout/blockscout/pull/7617) - Bump @amplitude/analytics-browser from 1.10.3 to 1.10.4 in /apps/block_scout_web/assets
- [#7609](https://github.com/blockscout/blockscout/pull/7609) - Bump webpack from 5.84.1 to 5.85.0 in /apps/block_scout_web/assets
- [#7610](https://github.com/blockscout/blockscout/pull/7610) - Bump mime from 2.0.4 to 2.0.5
- [#7634](https://github.com/blockscout/blockscout/pull/7634) - Bump eslint from 8.41.0 to 8.42.0 in /apps/block_scout_web/assets
- [#7633](https://github.com/blockscout/blockscout/pull/7633) - Bump floki from 0.34.2 to 0.34.3
- [#7631](https://github.com/blockscout/blockscout/pull/7631) - Bump phoenix_ecto from 4.4.1 to 4.4.2
- [#7630](https://github.com/blockscout/blockscout/pull/7630) - Bump webpack-cli from 5.1.1 to 5.1.3 in /apps/block_scout_web/assets
- [#7632](https://github.com/blockscout/blockscout/pull/7632) - Bump webpack from 5.85.0 to 5.85.1 in /apps/block_scout_web/assets
- [#7646](https://github.com/blockscout/blockscout/pull/7646) - Bump sweetalert2 from 11.7.10 to 11.7.11 in /apps/block_scout_web/assets
- [#7647](https://github.com/blockscout/blockscout/pull/7647) - Bump @amplitude/analytics-browser from 1.10.4 to 1.10.6 in /apps/block_scout_web/assets
- [#7659](https://github.com/blockscout/blockscout/pull/7659) - Bump webpack-cli from 5.1.3 to 5.1.4 in /apps/block_scout_web/assets
- [#7658](https://github.com/blockscout/blockscout/pull/7658) - Bump @amplitude/analytics-browser from 1.10.6 to 1.10.7 in /apps/block_scout_web/assets
- [#7657](https://github.com/blockscout/blockscout/pull/7657) - Bump webpack from 5.85.1 to 5.86.0 in /apps/block_scout_web/assets
- [#7672](https://github.com/blockscout/blockscout/pull/7672) - Bump @babel/preset-env from 7.22.4 to 7.22.5 in /apps/block_scout_web/assets
- [#7674](https://github.com/blockscout/blockscout/pull/7674) - Bump ecto from 3.10.1 to 3.10.2
- [#7673](https://github.com/blockscout/blockscout/pull/7673) - Bump @babel/core from 7.22.1 to 7.22.5 in /apps/block_scout_web/assets
- [#7671](https://github.com/blockscout/blockscout/pull/7671) - Bump sass from 1.62.1 to 1.63.2 in /apps/block_scout_web/assets
- [#7681](https://github.com/blockscout/blockscout/pull/7681) - Bump sweetalert2 from 11.7.11 to 11.7.12 in /apps/block_scout_web/assets
- [#7679](https://github.com/blockscout/blockscout/pull/7679) - Bump @amplitude/analytics-browser from 1.10.7 to 1.10.8 in /apps/block_scout_web/assets
- [#7680](https://github.com/blockscout/blockscout/pull/7680) - Bump sass from 1.63.2 to 1.63.3 in /apps/block_scout_web/assets
- [#7693](https://github.com/blockscout/blockscout/pull/7693) - Bump sass-loader from 13.3.1 to 13.3.2 in /apps/block_scout_web/assets
- [#7692](https://github.com/blockscout/blockscout/pull/7692) - Bump postcss-loader from 7.3.2 to 7.3.3 in /apps/block_scout_web/assets
- [#7691](https://github.com/blockscout/blockscout/pull/7691) - Bump url from 0.11.0 to 0.11.1 in /apps/block_scout_web/assets
- [#7690](https://github.com/blockscout/blockscout/pull/7690) - Bump core-js from 3.30.2 to 3.31.0 in /apps/block_scout_web/assets
- [#7701](https://github.com/blockscout/blockscout/pull/7701) - Bump css-minimizer-webpack-plugin from 5.0.0 to 5.0.1 in /apps/block_scout_web/assets
- [#7702](https://github.com/blockscout/blockscout/pull/7702) - Bump @amplitude/analytics-browser from 1.10.8 to 1.11.0 in /apps/block_scout_web/assets
- [#7708](https://github.com/blockscout/blockscout/pull/7708) - Bump phoenix_pubsub from 2.1.2 to 2.1.3
- [#7707](https://github.com/blockscout/blockscout/pull/7707) - Bump @amplitude/analytics-browser from 1.11.0 to 2.0.0 in /apps/block_scout_web/assets
- [#7706](https://github.com/blockscout/blockscout/pull/7706) - Bump webpack from 5.86.0 to 5.87.0 in /apps/block_scout_web/assets
- [#7705](https://github.com/blockscout/blockscout/pull/7705) - Bump sass from 1.63.3 to 1.63.4 in /apps/block_scout_web/assets
- [#7714](https://github.com/blockscout/blockscout/pull/7714) - Bump ex_cldr_units from 3.16.1 to 3.16.2
- [#7748](https://github.com/blockscout/blockscout/pull/7748) - Bump mock from 0.3.7 to 0.3.8
- [#7746](https://github.com/blockscout/blockscout/pull/7746) - Bump eslint from 8.42.0 to 8.43.0 in /apps/block_scout_web/assets
- [#7747](https://github.com/blockscout/blockscout/pull/7747) - Bump cldr_utils from 2.24.0 to 2.24.1

</details>

## 5.1.5-beta

### Features

- [#7439](https://github.com/blockscout/blockscout/pull/7439) - Define batch size for token balance fetcher via runtime env var
- [#7298](https://github.com/blockscout/blockscout/pull/7298) - Add changes to support force email verification
- [#7422](https://github.com/blockscout/blockscout/pull/7422) - Refactor state changes
- [#7416](https://github.com/blockscout/blockscout/pull/7416) - Add option to disable reCAPTCHA
- [#6694](https://github.com/blockscout/blockscout/pull/6694) - Add withdrawals support (EIP-4895)
- [#7355](https://github.com/blockscout/blockscout/pull/7355) - Add endpoint for token info import
- [#7393](https://github.com/blockscout/blockscout/pull/7393) - Realtime fetcher max gap
- [#7436](https://github.com/blockscout/blockscout/pull/7436) - TokenBalanceOnDemand ERC-1155 support
- [#7469](https://github.com/blockscout/blockscout/pull/7469), [#7485](https://github.com/blockscout/blockscout/pull/7485), [#7493](https://github.com/blockscout/blockscout/pull/7493) - Clear missing block ranges after every success import
- [#7489](https://github.com/blockscout/blockscout/pull/7489) - INDEXER_CATCHUP_BLOCK_INTERVAL env var

### Fixes

- [#7490](https://github.com/blockscout/blockscout/pull/7490) - Fix pending txs is not a map
- [#7474](https://github.com/blockscout/blockscout/pull/7474) - Websocket v2 improvements
- [#7472](https://github.com/blockscout/blockscout/pull/7472) - Fix RE_CAPTCHA_DISABLED variable parsing
- [#7391](https://github.com/blockscout/blockscout/pull/7391) - Fix: cannot read properties of null (reading 'value')
- [#7377](https://github.com/blockscout/blockscout/pull/7377), [#7454](https://github.com/blockscout/blockscout/pull/7454) - API v2 improvements

### Chore

- [#7496](https://github.com/blockscout/blockscout/pull/7496) - API v2: Pass backend version to the frontend
- [#7468](https://github.com/blockscout/blockscout/pull/7468) - Refactoring queries with blocks
- [#7435](https://github.com/blockscout/blockscout/pull/7435) - Add `.exs` and `.eex` checking in cspell
- [#7450](https://github.com/blockscout/blockscout/pull/7450) - Resolve unresponsive navbar in verification form page
- [#7449](https://github.com/blockscout/blockscout/pull/7449) - Actualize docker-compose readme and use latest tags instead main
- [#7417](https://github.com/blockscout/blockscout/pull/7417) - Docker compose for frontend
- [#7349](https://github.com/blockscout/blockscout/pull/7349) - Proxy pattern with getImplementation()
- [#7360](https://github.com/blockscout/blockscout/pull/7360) - Manage visibility of indexing progress alert

<details>
  <summary>Dependencies version bumps</summary>

- [#7351](https://github.com/blockscout/blockscout/pull/7351) - Bump decimal from 2.0.0 to 2.1.1
- [#7356](https://github.com/blockscout/blockscout/pull/7356) - Bump @amplitude/analytics-browser from 1.10.0 to 1.10.1 in /apps/block_scout_web/assets
- [#7366](https://github.com/blockscout/blockscout/pull/7366) - Bump mixpanel-browser from 2.46.0 to 2.47.0 in /apps/block_scout_web/assets
- [#7365](https://github.com/blockscout/blockscout/pull/7365) - Bump @amplitude/analytics-browser from 1.10.1 to 1.10.2 in /apps/block_scout_web/assets
- [#7368](https://github.com/blockscout/blockscout/pull/7368) - Bump cowboy from 2.9.0 to 2.10.0
- [#7370](https://github.com/blockscout/blockscout/pull/7370) - Bump ex_cldr_units from 3.15.0 to 3.16.0
- [#7364](https://github.com/blockscout/blockscout/pull/7364) - Bump chart.js from 4.2.1 to 4.3.0 in /apps/block_scout_web/assets
- [#7382](https://github.com/blockscout/blockscout/pull/7382) - Bump @babel/preset-env from 7.21.4 to 7.21.5 in /apps/block_scout_web/assets
- [#7381](https://github.com/blockscout/blockscout/pull/7381) - Bump highlight.js from 11.7.0 to 11.8.0 in /apps/block_scout_web/assets
- [#7379](https://github.com/blockscout/blockscout/pull/7379) - Bump @babel/core from 7.21.4 to 7.21.5 in /apps/block_scout_web/assets
- [#7380](https://github.com/blockscout/blockscout/pull/7380) - Bump postcss-loader from 7.2.4 to 7.3.0 in /apps/block_scout_web/assets
- [#7395](https://github.com/blockscout/blockscout/pull/7395) - Bump @babel/core from 7.21.5 to 7.21.8 in /apps/block_scout_web/assets
- [#7402](https://github.com/blockscout/blockscout/pull/7402) - Bump webpack from 5.81.0 to 5.82.0 in /apps/block_scout_web/assets
- [#7411](https://github.com/blockscout/blockscout/pull/7411) - Bump cldr_utils from 2.22.0 to 2.23.1
- [#7409](https://github.com/blockscout/blockscout/pull/7409) - Bump @amplitude/analytics-browser from 1.10.2 to 1.10.3 in /apps/block_scout_web/assets
- [#7410](https://github.com/blockscout/blockscout/pull/7410) - Bump sweetalert2 from 11.7.3 to 11.7.5 in /apps/block_scout_web/assets
- [#7434](https://github.com/blockscout/blockscout/pull/7434) - Bump ex_cldr from 2.37.0 to 2.37.1
- [#7433](https://github.com/blockscout/blockscout/pull/7433) - Bump eslint from 8.39.0 to 8.40.0 in /apps/block_scout_web/assets
- [#7432](https://github.com/blockscout/blockscout/pull/7432) - Bump tesla from 1.6.0 to 1.6.1
- [#7431](https://github.com/blockscout/blockscout/pull/7431) - Bump webpack-cli from 5.0.2 to 5.1.0 in /apps/block_scout_web/assets
- [#7430](https://github.com/blockscout/blockscout/pull/7430) - Bump core-js from 3.30.1 to 3.30.2 in /apps/block_scout_web/assets
- [#7443](https://github.com/blockscout/blockscout/pull/7443) - Bump webpack-cli from 5.1.0 to 5.1.1 in /apps/block_scout_web/assets
- [#7457](https://github.com/blockscout/blockscout/pull/7457) - Bump web3 from 1.9.0 to 1.10.0 in /apps/block_scout_web/assets
- [#7456](https://github.com/blockscout/blockscout/pull/7456) - Bump webpack from 5.82.0 to 5.82.1 in /apps/block_scout_web/assets
- [#7458](https://github.com/blockscout/blockscout/pull/7458) - Bump phoenix_ecto from 4.4.0 to 4.4.1
- [#7455](https://github.com/blockscout/blockscout/pull/7455) - Bump solc from 0.8.19 to 0.8.20 in /apps/explorer
- [#7460](https://github.com/blockscout/blockscout/pull/7460) - Bump jquery from 3.6.4 to 3.7.0 in /apps/block_scout_web/assets
- [#7488](https://github.com/blockscout/blockscout/pull/7488) - Bump exvcr from 0.13.5 to 0.14.1
- [#7486](https://github.com/blockscout/blockscout/pull/7486) - Bump redix from 1.2.2 to 1.2.3
- [#7487](https://github.com/blockscout/blockscout/pull/7487) - Bump tesla from 1.6.1 to 1.7.0
- [#7494](https://github.com/blockscout/blockscout/pull/7494) - Bump webpack from 5.82.1 to 5.83.0 in /apps/block_scout_web/assets
- [#7495](https://github.com/blockscout/blockscout/pull/7495) - Bump ex_cldr_numbers from 2.31.0 to 2.31.1

</details>

## 5.1.4-beta

### Features

- [#7273](https://github.com/blockscout/blockscout/pull/7273) - Support reCAPTCHA v3 in CSV export page
- [#7345](https://github.com/blockscout/blockscout/pull/7345) - Manage telegram link and its visibility in the footer
- [#7313](https://github.com/blockscout/blockscout/pull/7313) - API v2 new endpoints: watchlist transactions
- [#7286](https://github.com/blockscout/blockscout/pull/7286) - Split token instance fetcher
- [#7246](https://github.com/blockscout/blockscout/pull/7246) - Fallback JSON RPC option
- [#7329](https://github.com/blockscout/blockscout/pull/7329) - Delete pending block operations for empty blocks

### Fixes

- [#7317](https://github.com/blockscout/blockscout/pull/7317) - Fix tokensupply API v1 endpoint: handle nil total_supply
- [#7290](https://github.com/blockscout/blockscout/pull/7290) - Allow nil gas price for pending tx (Erigon node case)
- [#7288](https://github.com/blockscout/blockscout/pull/7288) - API v2 improvements: Fix tx type for pending contract creation; Remove owner for not unique ERC-1155 token instances
- [#7283](https://github.com/blockscout/blockscout/pull/7283) - Fix status for dropped/replaced tx
- [#7270](https://github.com/blockscout/blockscout/pull/7270) - Fix default `TOKEN_EXCHANGE_RATE_REFETCH_INTERVAL`
…- [#1559](https://github.com/blockscout/blockscout/pull/1559) - Change v column type for Transactions table

### Chore

- [#1579](https://github.com/blockscout/blockscout/pull/1579) - Add SpringChain to the list of Additional Chains Utilizing BlockScout
- [#1578](https://github.com/blockscout/blockscout/pull/1578) - Refine contributing procedure
- [#1572](https://github.com/blockscout/blockscout/pull/1572) - Add option to disable block rewards in indexer config

## 1.3.5-beta

### Features

- [#1560](https://github.com/blockscout/blockscout/pull/1560) - Allow executing smart contract functions in arbitrarily sized batches
- [#1543](https://github.com/blockscout/blockscout/pull/1543) - Use trace_replayBlockTransactions API for faster tracing
- [#1558](https://github.com/blockscout/blockscout/pull/1558) - Allow searching by token symbol
- [#1551](https://github.com/blockscout/blockscout/pull/1551) Exact date and time for Transaction details page
- [#1547](https://github.com/blockscout/blockscout/pull/1547) - Verify smart contracts with evm versions
- [#1540](https://github.com/blockscout/blockscout/pull/1540) - Fetch ERC721 token balances if sender is '0x0..0'
- [#1539](https://github.com/blockscout/blockscout/pull/1539) - Add the link to release in the footer
- [#1519](https://github.com/blockscout/blockscout/pull/1519) - Create contract methods
- [#1496](https://github.com/blockscout/blockscout/pull/1496) - Remove dropped/replaced transactions in pending transactions list
- [#1492](https://github.com/blockscout/blockscout/pull/1492) - Disable usd value for an empty exchange rate
- [#1466](https://github.com/blockscout/blockscout/pull/1466) - Decoding candidates for unverified contracts

### Fixes

- [#1545](https://github.com/blockscout/blockscout/pull/1545) - Fix scheduling of latest block polling in Realtime Fetcher
- [#1554](https://github.com/blockscout/blockscout/pull/1554) - Encode integer parameters when calling smart contract functions
- [#1537](https://github.com/blockscout/blockscout/pull/1537) - Fix test that depended on date
- [#1534](https://github.com/blockscout/blockscout/pull/1534) - Render a nicer error when creator cannot be determined
- [#1527](https://github.com/blockscout/blockscout/pull/1527) - Add index to value_fetched_at
- [#1518](https://github.com/blockscout/blockscout/pull/1518) - Select only distinct failed transactions
- [#1516](https://github.com/blockscout/blockscout/pull/1516) - Fix coin balance params reducer for pending transaction
- [#1511](https://github.com/blockscout/blockscout/pull/1511) - Set correct log level for production
- [#1510](https://github.com/blockscout/blockscout/pull/1510) - Fix test that fails every 1st day of the month
- [#1509](https://github.com/blockscout/blockscout/pull/1509) - Add index to blocks' consensus
- [#1508](https://github.com/blockscout/blockscout/pull/1508) - Remove duplicated indexes
- [#1505](https://github.com/blockscout/blockscout/pull/1505) - Use https instead of ssh for absinthe libs
- [#1501](https://github.com/blockscout/blockscout/pull/1501) - Constructor_arguments must be type `text`
- [#1498](https://github.com/blockscout/blockscout/pull/1498) - Add index for created_contract_address_hash in transactions
- [#1493](https://github.com/blockscout/blockscout/pull/1493) - Do not do work in process initialization
- [#1487](https://github.com/blockscout/blockscout/pull/1487) - Limit geth sync to 128 blocks
- [#1484](https://github.com/blockscout/blockscout/pull/1484) - Allow decoding input as utf-8
- [#1479](https://github.com/blockscout/blockscout/pull/1479) - Remove smoothing from coin balance chart

### Chore

- [https://github.com/blockscout/blockscout/pull/1532](https://github.com/blockscout/blockscout/pull/1532) - Upgrade elixir to 1.8.1
- [https://github.com/blockscout/blockscout/pull/1553](https://github.com/blockscout/blockscout/pull/1553) - Dockerfile: remove 1.7.1 version pin FROM bitwalker/alpine-elixir-phoenix
- [https://github.com/blockscout/blockscout/pull/1465](https://github.com/blockscout/blockscout/pull/1465) - Resolve lodash security alert

