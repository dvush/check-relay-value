const ethers = require("ethers")
const arg = require('arg');
const axios = require("axios");

const args = arg({
	'--provider': String,
	'--relay': String,
    '--slot': String,
});

async function main() {
    const providerUrl = args['--provider']
    const relay = args['--relay']
    const slot = args['--slot']

    const provider = new ethers.providers.JsonRpcProvider(providerUrl)

    const payload = (await axios.get(`${relay}/relay/v1/data/bidtraces/proposer_payload_delivered?slot=${slot}`)).data[0]


    const address = payload.proposer_fee_recipient
    const value = payload.value
    const blockHeader = await provider.getBlock(payload.block_hash)
    const block = blockHeader.number

    const balanceBefore = await provider.getBalance(address, block - 1)
    const balanceAfter = await provider.getBalance(address, block)
    const diff = balanceAfter.sub(balanceBefore)
    console.log("Payload value: ", ethers.utils.formatEther(value))
    console.log("Balance diff:", ethers.utils.formatEther(diff))
    console.log("Delta:", ethers.utils.formatEther(diff.sub(value)))
}

main()
    .then(() => process.exit(0))
    .catch((err) => {
        console.error(err)
        process.exit(1)
    })