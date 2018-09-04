package transactions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

/* Takes a hex string and converts the number to a int*/
func HextoDecimal(hexinput string) (result int64, err error) {
	rmPrefix := strings.SplitAfter(hexinput, "x")

	x, err := strconv.ParseInt(rmPrefix[1], 16, 32)
	if err != nil {
		fmt.Printf("Error in ParseInt conversion: %s \n", err)
		return -1, err
	}

	result = int64(x)

	return result, nil
}

// Wrapper around the http Request
func httpRequest(method string, url string, params ParamInfo) ([]byte, error) {
	client := &http.Client{}
	var reqBody string

	jsonParams, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	reqBody = string(jsonParams)

	req, err := http.NewRequest(method, url, strings.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// Returns the latest mined block
func GetLatestBlockNumber() (int64, error) {

	var blockNumberResp BlockNumberResponse

	p1 := ParamInfo{
		Jsonrpc: "2.0",
		Method:  "eth_blockNumber",
		Params:  []interface{}{},
		Id:      83,
	}

	body1, _ := httpRequest("POST", "https://mainnet.infura.io/HozhwE8rLsXcdbLG9yL", p1)

	err := json.Unmarshal([]byte(body1), &blockNumberResp)
	if err != nil {
		fmt.Print("ERROR with the blockNumber UnmarshalL: %s ", err)
		return 0, err
	}

	blockNumberInt, err := HextoDecimal(blockNumberResp.Result)
	if err != nil {
		return 0, err
	}

	return blockNumberInt, nil

}

//Returns the Block specified by the input given
func GetFullBlock(blocknumber int64) (*BlockByNumberResponse, error) {

	var blocktimeresp *BlockByNumberResponse
	var paramconv = big.NewInt(blocknumber)
	startBlockStr := hexutil.EncodeBig(paramconv)

	p2 := ParamInfo{
		Jsonrpc: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  []interface{}{startBlockStr, true},
		Id:      1,
	}

	body1, _ := httpRequest("POST", "https://mainnet.infura.io/HozhwE8rLsXcdbLG9yL", p2)

	err := json.Unmarshal([]byte(body1), &blocktimeresp)
	if err != nil {
		fmt.Printf("ERROR with the blockByNumber Unmarshall: %s \n", err)
		return nil, err
	}
	return blocktimeresp, nil
}

// Gets the starting block adjusted within the appropriate time frame
func GetStartingBlockNumber(days int64) (int64, error) {

	const blocksPerDay int64 = 5760

	latestBlockNumber, err := GetLatestBlockNumber()
	if err != nil {
		fmt.Println("Error with retrieving the latest block number: ", err)
		return 0, err
	}

	startingBlock := latestBlockNumber - (blocksPerDay * days)

	return startingBlock, nil

}
func GetTransactions(address string, days int64) ([]TransactionBlockEtherScan, error) {

	var transactionlist Configuration
	var etherScanUrl, etherScanKey, contractAddress, startingBlockNumber, endingBlockNumber string
	etherScanKey = "" //Specific to EtherScan account - USER MUST SET
	FullListOfTransactions := []TransactionBlockEtherScan{}
	var counter int64 = 0

	const blocksPerCall int64 = 1920

	startingblock, err := GetStartingBlockNumber(days)
	if err != nil {
		fmt.Println("Error with retrieving the Starting Block Number: ", err)
		return nil, err
	}
	currBlockNumber, err := GetLatestBlockNumber()
	if err != nil {
		fmt.Println("Error retrieving the Latest Block Number: ", err)
		return nil, err
	}

	limit := ((currBlockNumber - startingblock) / blocksPerCall)

	for counter < limit {

		blknumstr := strconv.FormatInt((startingblock + (blocksPerCall * counter)), 10)
		startingBlockNumber = blknumstr
		endingBlockStr := strconv.FormatInt((startingblock)+(blocksPerCall*(counter+1)), 10)
		endingBlockNumber = endingBlockStr
		contractAddress = address
		etherScanUrl = "https://api.etherscan.io/api?module=account&action=txlist&address="
		etherScanUrl += contractAddress + "&startblock=" + startingBlockNumber + "&endblock=" + endingBlockNumber + "&sort=asc&apikey=" + etherScanKey

		res, err := http.Get(etherScanUrl)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		temp, err := ioutil.ReadAll(res.Body)

		res.Body.Close()
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		err2 := json.Unmarshal([]byte(temp), &transactionlist)
		if err2 != nil {
			fmt.Printf("Error with Unmarshall %s ", err2)
			return nil, err2
		}

		for i := 0; i < len(transactionlist.Result); i++ {
			FullListOfTransactions = append(FullListOfTransactions, transactionlist.Result[i])
		}
		transactionlist.Result = nil
		counter += 1
	}
	return FullListOfTransactions, nil
}
