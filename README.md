# Prepare

```
go build
```

# Usage

```
./relay-info -h

Usage of ./relay-info:
  -R flashbots
    	relay address. Either alias (e.g. flashbots) or url (e.g. https://boost-relay.flashbots.net)
  -r string
    	ethereum RPC address (or from env $ETH_RPC_URL)
Commands:
 relays
  list available relay aliases
 status
  check relay status

 validators
  list validators for current epoch

 delivered
  payloads delivered by the relay
  -block uint
    	payloads fot the specific block number
  -check
    	verify payload value delivered (needs Ethereum RPC)
  -limit int
    	limit number of payloads (max: 200)
  -slot uint
    	payloads for the specific slot number
```

## Relays

List relays

```
./relay-info relays

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
./relay-info -R flashbots status

OK, time: 241.549696ms
```

## Validators
List validators for the current epoch

```
./relay-info -R flashbots validators

slot    proposer pubkey                                                                                    fee recepient                              timestamp  gaslimit 
4903680 0x8d2b505bbf55185201c2df4863b553fe2305a924668e1383a4010a677205c4be6ba4ef221a23b6f425e1fcd5b81b6e4b 0x388c818ca8b9251b393131c08a736a67ccb19297 1665496511 30000000
4903681 0xb613811f915a207f30a076fd4f3c4d33e5ee6e44eefacd2867d73090e31596c7fba75e27381eaff3276d4a56f5620ce7 0xebec795c9c8bbd61ffc14a6662944748f299cacf 1663240912 30000000
4903682 0xb57e034bc138c79738e1e2df18c021de5f923d99991c4b0c801ad22896fd6196310f219f0155a59b438e94bea9461107 0xe688b84b23f322a994a53dbf8e15fa82cdb71127 1663645375 30000000
...
```

## Delivered

List payloads delivered by the relay.

Use flags `--slot` or `--block` to see payload delivered for the concrete slot or block.

1. Without `--check` flag it will simply output payloads delivered by the relay

```
./relay-info -R flashbots delivered --slot 4903741

slot    block hash                                                         builder    gas used gas limit value
4903741 0x6132a55317beed1c644d58592d0a5c2fc631b7cd3826f38d634acc440b3a6366 0x81babeec 20576237  30000000 0.07622905111
```

2. With `--check` flag it will try output additional info about delivered payload if it was committed on ethereum.

It needs Ethereum rpc available (set with `-r` flag or `$ETH_RPC_URL` env).

```
./relay-info -R flashbots -r http://localhost:8545 delivered --block 15739594 --check  

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
