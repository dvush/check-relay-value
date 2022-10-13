# Prepare

```
go build
```

# Usage

See `./boost-cli -h`

## Relays

List relays

```
./boost-cli relays

Available relays:
flashbots            https://boost-relay.flashbots.net
bloxroute.max-profit https://bloxroute.max-profit.blxrbdn.com
bloxroute.ethical    https://bloxroute.ethical.blxrbdn.com
bloxroute.regulated  https://bloxroute.regulated.blxrbdn.com
blocknative          https://builder-relay-mainnet.blocknative.com
eden                 https://relay.edennetwork.io
manifold             https://mainnet-relay.securerpc.com
```

## Status

```
./boost-cli -R flashbots status

OK
```

## Validators
List validators for the current epoch

```
./boost-cli -R flashbots validators

slot    proposer pubkey                                                                                    fee recepient                              timestamp  gaslimit 
4903680 0x8d2b505bbf55185201c2df4863b553fe2305a924668e1383a4010a677205c4be6ba4ef221a23b6f425e1fcd5b81b6e4b 0x388c818ca8b9251b393131c08a736a67ccb19297 1665496511 30000000
4903681 0xb613811f915a207f30a076fd4f3c4d33e5ee6e44eefacd2867d73090e31596c7fba75e27381eaff3276d4a56f5620ce7 0xebec795c9c8bbd61ffc14a6662944748f299cacf 1663240912 30000000
4903682 0xb57e034bc138c79738e1e2df18c021de5f923d99991c4b0c801ad22896fd6196310f219f0155a59b438e94bea9461107 0xe688b84b23f322a994a53dbf8e15fa82cdb71127 1663645375 30000000
4903685 0xa3715f86e94ef159e7471395869a3b40b5502c58ffe412dcc44c12777c4434586ac6473ea59cc4532381c5e080b8be33 0x388c818ca8b9251b393131c08a736a67ccb19297 1663924800 30000000
4903687 0xb7f25118ae4762e90ac78e6ff0facdbe89a9844d143913ff9f937698ac48f2dc6931cbea69e92c577efe3ab1d7f54c42 0xa1a547358a9ca8e7b320d7742729e3334ad96546 1663635416 30000000
4903691 0xab3694d4b6de619b3700a147b0ca8b2200164a4cd9aa7c626e50dd3248c959196aadd3b418cc0dccb52c37af3ff2c2f8 0x388c818ca8b9251b393131c08a736a67ccb19297 1664187021 30000000
4903692 0xa18d05ee05c6f5536d61098f9525a293e278140f1d92668250c742861ed57bc64853e6630eb8a4297032ff6d72141b77 0x63fc27188ba91c75bd1ee4f807fdbce6d3251db3 1662986243 30000000
4903693 0xb925f9179cb87dbbe3d4a7ee75931dcfd4142c38a2089d6358f277a14913e0ca24e6f591ed140ad2f673fed32e9aa6c9 0x0077732907bfc6208933cfd2a51afb8f33ca5958 1663262896 30000000
4903694 0x999bf42c85c224c2167659a1945bc60b3fd106ee14f187c606f38152a024083d3b3fb03e77f7aca20146ed167be8603b 0x1cedc0f3af8f9841b0a1f5c1a4ddc6e1a1629074 1663292675 30000000
4903696 0x941ddeacdf3998238491fbc5caed51acd6416eef74fa32d708ec8a1543c071f493a08f7e68f21f00619b8943e316a371 0x388c818ca8b9251b393131c08a736a67ccb19297 1663670087 30000000
4903700 0xa881ca1f2c7e5eff22d9d4ee476fa61c0f27c9f07828c57ed82f9573ea04d696afc959049f79596d83f4b34281fcbed1 0xebec795c9c8bbd61ffc14a6662944748f299cacf 1663240900 30000000
4903702 0xb5edd9bc3651e17ad3a797fe047fc9d687d9c64bf088cc969743b1ca1b8435f944af7480f27f02ad8e55660d971059bd 0x221507c5cae31196a535f223c022eb0e38c3377d 1663514327 30000000
4903703 0xa274a1c114f6676981ebcf96100d28facf7ec9937deb388953fcc6d30c4c27c761c57e2ad724bbac239b6216b38e9e5c 0xebec795c9c8bbd61ffc14a6662944748f299cacf 1663239367 30000000
4903705 0xa20ccb305672d0a3fecb7c66f01fe27419b8f7f4b225d42a6592f50212f9bbaa23d4e0b953445c4e1ad6cceb36d0eda1 0x388c818ca8b9251b393131c08a736a67ccb19297 1663735638 30000000
4903707 0xac3b7cd8038e0004232976e2c83d5213224ad24a540442b05acf6f1b25486df7eb5b30f0bd16c4e2eceef4811302fbda 0xe688b84b23f322a994a53dbf8e15fa82cdb71127 1665458584 30000000
4903708 0x987ef2beadbcbe49b1f0fa679b97780710cbbd959d2209f715def67fac73a645a0f30962d8da5a96e9e24c6c14b806d2 0xea7459c7e0ce2ce114ee4da948c5d994d6cd7ed1 1663234767 30000000
4903710 0xab422594bb2f1683915f20f6a32d717a9d45bae21cb1b69a7b310deeb6ec1a4c14bf15d9ad67e82c5598d32e4a4e8703 0xe688b84b23f322a994a53dbf8e15fa82cdb71127 1665460955 30000000
4903712 0xa4a82987fd29b962367ec6da8605c69e5ab38016e31833f369145e2a0f00a228a21d6b707b05e8dcd8fa97d25ac3abbf 0x388c818ca8b9251b393131c08a736a67ccb19297 1663670123 30000000
4903713 0xb522942ef75ae2fdd6dcabdafc9723451371c47129e060c40a522605b64cd428f5033087803a39891fe74c4673d5710e 0xe688b84b23f322a994a53dbf8e15fa82cdb71127 1665464747 30000000
4903715 0x94ec8668cdb7b81babf556e45d19f42bf11472072e328add38b18024721718f83922b10ecf7fb37e69f56742c9a32bf7 0x388c818ca8b9251b393131c08a736a67ccb19297 1663245539 30000000
4903716 0x98151055c5b6cb5306273887317fd4e80f5b4eb538370796b43d2c2d02822667bc205e1e100225611bf75a60aef02ca3 0xe94f1fa4f27d9d288ffea234bb62e1fbc086ca0c 1663773843 30000000
4903719 0x816a33dde4999a18339289f3c4e07dee9860c5e3b624664c923d94548967192e8e4fb811fb145967c4e3e5e06dd47d26 0x906219a7b1843e432a80714b584b8ce93c2005ed 1664242956 30000000
4903720 0xb174823009bac8c3a7097bc7449af9766e828340b2bca3e2f69bde2cea6f36e5b386c14ecfd41b04c989b5ae0403b55d 0xeee27662c2b8eba3cd936a23f039f3189633e4c8 1663239475 30000000
4903723 0x9932b3b2c6bafd8cbeac05581bb04d7a7ad9acefe30dd824659c9dc0ca3be7855a8c637831dc6d13b24d8374ddbc65d0 0x388c818ca8b9251b393131c08a736a67ccb19297 1664468001 30000000
4903726 0xa5988db5d4c6758c8c9aff6faf1392faf368e02447db93a8546ef9bdbbe68237efc80b4490ba128857306b12386e8869 0xffee087852cb4898e6c3532e776e68bc68b1143b 1664305475 30000000
4903728 0xb0f114ebfd09b3c24a121f3512c61778c5eec7e805f985509294255144feb2a236a1ec59fb63a75a8330c22020d81162 0x326ce3994d16e8ac6193d4a282ecd6be58818866 1665339705 30000000
4903732 0xb78dea7af9410ee9d18327079affef88263bc37748107cf59616a380d33e4d3292e1fd6d6ab73b967f0c5c20294be3f1 0xadae1798f761fa7fce29b6673d453d1a48a2931a 1664134118 30000000
4903733 0xa5ce84af1a1d243974710d1bb98894341272f489b5fa821d2a33e42e39ae2a0cea99cf3d94336a996ad9e313432eda2f 0xe94f1fa4f27d9d288ffea234bb62e1fbc086ca0c 1663262858 30000000
4903734 0xb906dfaf99cf8c0323deb64f68846ab3eac6d1878ea3bce3b6f65b18979c96b5bbeb92fd94d72dbd9ef415344e24d4e4 0xe688b84b23f322a994a53dbf8e15fa82cdb71127 1665460079 30000000
4903735 0x971ac2c530ace118a894c4d743d632c89cab2d8d332de10e6a1be5cea533f8ac1c0a3e3a6e6a70ee57bda5b81b3f8629 0x388c818ca8b9251b393131c08a736a67ccb19297 1663669715 30000000
4903739 0x97f5faaf61fbbd3bd2d59468b6a2d7638263c35f6cf71d83beb41c7475910e1b74c263c9a8283a9f7dc22b2791be7894 0x0b70b578abd96aab5e80d24d1f3c28dbde14356a 1663850231 30000000
4903740 0xb80b3a110664c928fb1f505e8a6883f53f82f3e85e285e53190bdfea0c248b89f535922449f446fc0c4ea168d3e025f2 0xe688b84b23f322a994a53dbf8e15fa82cdb71127 1665459106 30000000
4903741 0x969048b7dbddd640eaef937f26353c6022c495ca9fa81780c6e36b2c08c8dd5935498187bebfa3ca7fd8974b923e6761 0x388c818ca8b9251b393131c08a736a67ccb19297 1664441908 30000000
4903742 0xb4b7ca0f9bf1d8221b2ca7a96b7eab3669727a2fdfcf0a0fc98c52e24b518ce0f3b0f713f0fa4ce89840fb7d728695cd 0x301407427168fb51bcc927b9fb76dcd88fe45681 1663245527 30000000
4903743 0x8d6efe5656b9adcf1b9ca2a4560e49cc3d38a74982a00e3774ff3e7048cd11640adb48c7fd9884efa7d3a76fade083ea 0xf22e098cb0d6c24d88ab79c1ce87b73ea26e608f 1663249341 30000000
```

## Delivered

List payloads delivered by the relay.

Use flags `--slot` or `--block` to see payload delivered for the concrete slot or block.

1. Without `--check` flag it will simply output payloads delivered by the relay

```
./boost-cli -R flashbots delivered --slot 4903741

slot    block hash                                                         builder    gas used gas limit value
4903741 0x6132a55317beed1c644d58592d0a5c2fc631b7cd3826f38d634acc440b3a6366 0x81babeec 20576237  30000000 0.07622905111
```

2. With `--check` flag it will try output additional info about delivered payload if it was committed on ethereum.

It needs Ethereum rpc available (set with `-r` flag or `$ETH_RPC_URL` env).

```
./boost-cli -R flashbots -r http://localhost:8545 delivered --block 15739594 --check  

slot    block    builder    proposer fee rec.                          block fee rec.                             claimed value  prop. diff.    value delta    rec. diff.    
4903741 15739594 0x81babeec 0x388c818ca8b9251b393131c08a736a67ccb19297 0xDAFEA492D9c6733ae3d56b7Ed1ADB60692c98Bc5 0.07622905111  0.07622905111  0              0             
```

Where
+ `proposer fee rec.` is block proposer fee recipient
Its address specified by the block proposer that should receive ETH.
+ `block fee rec.` is block fee recipient specified in the block header also known as coinbase address.
For some relays it can be the same address as in `proposer fee rec.` for others it's an address of the block builder.
+ `claimed value` values field in the delivered payload.
This value should correspond to the profit of the block proposer but relays can use their own methods on how to calculate it.
+ `prop. diff` is a balance increment of the `proposer fee rec.` during the given block. 
Usually it should be equal to the `claimed value` but it can diverge in case of relay misbehavior or if `proposer fee rec.`
transfers ETH in the given block.
+ `value delta` is `claimed value - prop. diff`
+ `rec. diff.` is a balance increment of the `block fee rec.` during the given block.
If `block fee rec.` is set to the block builder address this value can indicate block builder profit or loss in the given block.
