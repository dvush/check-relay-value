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

Payload value:  0.597786452187815763
Balance diff: 0.597786452187815763
Delta: 0.0
```

Negative delta -> balance diff is less than relay payload value