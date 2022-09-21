const ethers = require("ethers")
const arg = require('arg');

const args = arg({
	'--provider': String,
	'--address': String,
    '--block': Number,
});

async function main() {
    const providerUrl = args['--provider']
    const address = args['--address']
    const block = args['--block']

    const provider = new ethers.providers.JsonRpcProvider(providerUrl)

    const balanceBefore = await provider.getBalance(address, block - 1)
    const balanceAfter = await provider.getBalance(address, block)
    const diff = balanceAfter.sub(balanceBefore)
    console.log("balance diff(eth):", ethers.utils.formatEther(diff))
    console.log("balance diff:", diff.toString())
}

main()
    .then(() => process.exit(0))
    .catch((err) => {
        console.error(err)
        process.exit(1)
    })