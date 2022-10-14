package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	boost "github.com/flashbots/go-boost-utils/types"
	"io"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type ErrorMessage struct {
	code    int
	message string
}

func request(relayUrl, path string, values url.Values) ([]byte, error) {
	path, err := url.JoinPath(relayUrl, path)
	if err != nil {
		return nil, err
	}
	if values != nil {
		path = fmt.Sprintf("%s?%s", path, values.Encode())
	}
	res, err := http.Get(path)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return body, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return body, fmt.Errorf("http error, code: %d", res.StatusCode)
	}

	var errorMessage ErrorMessage
	err = json.Unmarshal(body, &errorMessage)
	if err == nil {
		return body, fmt.Errorf("rpc error, code: %d, message: %s", errorMessage.code, errorMessage.message)
	}
	return body, nil
}

func Status(relayUrl string) error {
	start := time.Now()
	_, err := request(relayUrl, "/eth/v1/builder/status", nil)
	if err != nil {
		return err
	}
	log.Printf("OK, time: %s\n", time.Since(start))

	return nil
}

func Validators(relayUrl string) error {
	body, err := request(relayUrl, "/relay/v1/builder/validators", nil)

	var result []boost.BuilderGetValidatorsResponseEntry
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	log.Printf("%-7s %-98s %-42s %-10s %-8s \n", "slot", "proposer pubkey", "fee recepient", "timestamp", "gaslimit")
	for _, entry := range result {
		reg := entry.Entry.Message
		log.Printf("%d %s %s %d %d\n", entry.Slot, reg.Pubkey, reg.FeeRecipient, reg.Timestamp, reg.GasLimit)
	}

	return nil
}

func reportBidTraces(bids []boost.BidTrace) {
	log.Printf("%-7s %-66s %-10s %-8s %-8s %-8s \n", "slot", "block hash", "builder", "gas used", "gas limit", "value")
	for _, entry := range bids {
		log.Printf("%d %s %s %8d %9d %s\n", entry.Slot, entry.BlockHash, entry.BuilderPubkey.String()[:10], entry.GasUsed, entry.GasLimit, ethValue(entry.Value.BigInt()))
	}
}

func balanceDiff(client *ethclient.Client, address common.Address, block uint64) (*big.Int, error) {
	balanceBefore, err := client.BalanceAt(context.Background(), address, big.NewInt(int64(block)-1))
	if err != nil {
		return nil, err
	}

	balanceAfter, err := client.BalanceAt(context.Background(), address, big.NewInt(int64(block)))
	if err != nil {
		return nil, err
	}

	return balanceBefore.Sub(balanceAfter, balanceBefore), nil
}

func checkRelayBids(eth *ethclient.Client, bids []boost.BidTrace) {
	log.Printf("%-7s %-8s %-10s %-42s %-42s %-14s %-14s %-14s %-14s\n", "slot", "block", "builder", "proposer fee rec.", "block fee rec.", "claimed value", "prop. diff.", "value delta", "rec. diff.")
	for _, bid := range bids {
		blockData, err := eth.BlockByHash(context.Background(), common.Hash(bid.BlockHash))
		if err != nil {
			log.Println("error", err)
			continue
		}
		proposerFeeRecipient := bid.ProposerFeeRecipient
		blockFeeRecipient := blockData.Coinbase()

		propDiff, err := balanceDiff(eth, common.Address(proposerFeeRecipient), blockData.NumberU64())
		if err != nil {
			log.Println("error", err)
			continue
		}

		blockRecDiff, err := balanceDiff(eth, blockFeeRecipient, blockData.NumberU64())
		if err != nil {
			log.Println("error", err)
			continue
		}

		log.Printf("%d %d %-10s %s %s %-14s %-14s %-14s %-14s\n", bid.Slot, blockData.NumberU64(), bid.BuilderPubkey.String()[:10],
			proposerFeeRecipient.String(), blockFeeRecipient.String(), ethValue(bid.Value.BigInt()),
			ethValue(propDiff), ethValue(new(big.Int).Sub(bid.Value.BigInt(), propDiff)), ethValue(blockRecDiff))
	}
}

func PayloadDelivered(relayUrl string, slot, block uint64, limit int, check bool) error {
	v := url.Values{}
	if slot != 0 {
		v.Set("slot", strconv.FormatUint(slot, 10))
	} else if block != 0 {
		v.Set("block_number", strconv.FormatUint(block, 10))
	}

	if limit != 0 {
		if limit > 200 {
			limit = 200
		}
		v.Set("limit", strconv.Itoa(limit))
	}

	body, err := request(relayUrl, "/relay/v1/data/bidtraces/proposer_payload_delivered", v)
	if err != nil {
		return err
	}

	var result []boost.BidTrace
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	if check {
		eth, err := DialETHRpc()
		if err != nil {
			return err
		}
		defer eth.Close()

		checkRelayBids(eth, result)
	} else {
		reportBidTraces(result)
	}

	return nil
}

func ethValue(int *big.Int) string {
	val := new(big.Float).SetInt(int)
	return val.Quo(val, new(big.Float).SetInt(big.NewInt(params.Ether))).String()
}

func DialETHRpc() (*ethclient.Client, error) {
	var rawurl string
	if *ethRpc != "" {
		rawurl = *ethRpc
	} else {
		rawurl = os.Getenv("ETH_RPC_URL")
	}
	if rawurl == "" {
		return nil, errors.New("empty ethereum RPC URL")
	}
	return ethclient.Dial(rawurl)
}

var relays = map[string]string{
	"flashbots":            "https://boost-relay.flashbots.net",
	"bloxroute.max-profit": "https://bloxroute.max-profit.blxrbdn.com",
	"bloxroute.ethical":    "https://bloxroute.ethical.blxrbdn.com",
	"bloxroute.regulated":  "https://bloxroute.regulated.blxrbdn.com",
	"blocknative":          "https://builder-relay-mainnet.blocknative.com",
	"eden":                 "https://relay.edennetwork.io",
	"manifold":             "https://mainnet-relay.securerpc.com",
}

func PrintRelays() {
	log.Println("Available relays:")
	for k, v := range relays {
		log.Printf("%-20s %s", k, v)
	}
}

func GetRelayAddress(relay string) (string, error) {
	if relay == "" {
		return "", fmt.Errorf("empty relay address")
	}

	if addr, ok := relays[relay]; ok {
		return addr, nil
	}

	if _, err := url.ParseRequestURI(relay); err != nil {
		return "", fmt.Errorf("failed to parse relay url: %w", err)
	}

	return relay, nil
}

var (
	relay  = flag.String("R", "", "relay address. Either alias (e.g. `flashbots`) or url (e.g. https://boost-relay.flashbots.net)")
	ethRpc = flag.String("r", "", "ethereum RPC address (or from env $ETH_RPC_URL)")

	payloadFlags = flag.NewFlagSet("payload delivered", flag.ExitOnError)
	payloadSlot  = payloadFlags.Uint64("slot", 0, "payloads for the specific slot number")
	payloadBlock = payloadFlags.Uint64("block", 0, "payloads fot the specific block number")
	payloadLimit = payloadFlags.Int("limit", 0, "limit number of payloads (max: 200)")
	payloadCheck = payloadFlags.Bool("check", false, "verify payload value delivered (needs Ethereum RPC)")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(flag.CommandLine.Output(), "Commands:")
		fmt.Fprintln(flag.CommandLine.Output(), " relays\n  list available relay aliases")
		fmt.Fprintln(flag.CommandLine.Output(), " status\n  check relay status")
		fmt.Fprintln(flag.CommandLine.Output())
		fmt.Fprintln(flag.CommandLine.Output(), " validators\n  list validators for current epoch")
		fmt.Fprintln(flag.CommandLine.Output())
		fmt.Fprintln(flag.CommandLine.Output(), " delivered\n  payloads delivered by the relay")
		payloadFlags.PrintDefaults()
	}
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	getRelayAddress := func() string {
		relayAddress, err := GetRelayAddress(*relay)
		if err != nil {
			log.Printf("Can't get relay address: %q", err)
			PrintRelays()
			os.Exit(1)
		}
		return relayAddress
	}

	args := flag.Args()
	if len(args) == 0 {
		log.Fatalf("No subcommand specified")
	}
	cmd, args := args[0], args[1:]

	switch cmd {
	case "status":
		relayAddress := getRelayAddress()
		err := Status(relayAddress)
		if err != nil {
			log.Fatal(err)
		}
	case "relays":
		PrintRelays()
	case "validators":
		relayAddress := getRelayAddress()
		err := Validators(relayAddress)
		if err != nil {
			log.Fatal(err)
		}
	case "delivered":
		err := payloadFlags.Parse(args)
		if err != nil {
			log.Fatal(err)
		}
		relayAddress := getRelayAddress()
		err = PayloadDelivered(relayAddress, *payloadSlot, *payloadBlock, *payloadLimit, *payloadCheck)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("Incorrect subcommand")
	}
}
