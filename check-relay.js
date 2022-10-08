const ethers = require("ethers")
const arg = require('arg');
const axios = require("axios");

const args = arg({
	'--provider': String,
	'--relay': String,
    '--slot': String,
});

async function getBalanceDiff(provider, address, block) {
    const balanceBefore = await provider.getBalance(address, block - 1)
    const balanceAfter = await provider.getBalance(address, block)
    return balanceAfter.sub(balanceBefore)
}

async function main() {
    const providerUrl = args['--provider']
    const relay = args['--relay']
    const slot = args['--slot']

    const provider = new ethers.providers.JsonRpcProvider(providerUrl)

    const payload = (await axios.get(`${relay}/relay/v1/data/bidtraces/proposer_payload_delivered?slot=${slot}`)).data[0]


    const feeRecipient = payload.proposer_fee_recipient
    const value = payload.value
    const blockHeader = await provider.getBlock(payload.block_hash)
    const builderAddress = blockHeader.miner
    const block = blockHeader.number

    const builderDiff = await getBalanceDiff(provider, builderAddress, block)
    const feeRecipientDiff = await getBalanceDiff(provider, feeRecipient, block)

    console.log("Slot:                                ", payload.slot)
    console.log("Block hash:                          ", payload.block_hash)
    console.log("Block number:                        ", block.toString())
    console.log("Proposer fee recipient:              ", feeRecipient.toLowerCase())
    console.log("Block fee recipient:                 ", builderAddress.toLowerCase())
    console.log("Claimed payload value:               ", ethers.utils.formatEther(value))
    console.log("Proposer fee recipient balance diff: ", ethers.utils.formatEther(feeRecipientDiff))
    console.log("Fee received - Payload value:        ", ethers.utils.formatEther(feeRecipientDiff.sub(value)))
    console.log("Block fee recipient balance diff:    ", ethers.utils.formatEther(builderDiff))
}

main()
    .then(() => process.exit(0))
    .catch((err) => {
        console.error(err)
        process.exit(1)
    })
