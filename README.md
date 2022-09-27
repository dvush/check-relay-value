# Prepare

Install deps
```bash
yarn install
```

# Check specific address diff during the given block

```bash
yarn run diff --provider https://mainnet.infura.io/v3/.... --address 0x003a4cc04501e9adecf850313db2d3797df801a2 --block 15583124


balance diff(eth): 0.071640660271666334
balance diff: 71640660271666334
```

# Check payload delivered by the relay

```bash
yarn run check-relay --provider https://mainnet.infura.io/v3/.... --relay https://boost-relay.flashbots.net --slot 4746672

Slot:                                 4746672
Block hash:                           0x50591139693c5e4b786c88a0ee146d6e0915e40137efbfc5547b531312502c2e
Block number:                         15583462
Proposer fee recipient:               0x388c818ca8b9251b393131c08a736a67ccb19297
Block fee recipient:                  0xdafea492d9c6733ae3d56b7ed1adb60692c98bc5
Payload value:                        0.597786452187815763
Proposer fee recipient balance diff:  0.597786452187815763
Fee received - Payload value:         0.0
Block fee recipient balance diff:     0.000132265956706032
```