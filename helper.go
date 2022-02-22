package main

import (
	"fmt"
	"github.com/incognitochain/go-incognito-sdk-v2/common"
	"github.com/incognitochain/go-incognito-sdk-v2/metadata"
	"github.com/incognitochain/go-incognito-sdk-v2/rpchandler"
)

func getTxsByBlockNumber(blkHeight uint64, shardID byte) ([]string, error) {
	type RPCResult struct {
		TxHashes []string
	}

	rpcMethod := "retrieveblockbyheight"
	responseInBytes, err := ic.NewRPCCall("", rpcMethod, []interface{}{blkHeight, shardID, "1"}, 1)
	if err != nil {
		return nil, err
	}
	var tmpRes []RPCResult
	err = rpchandler.ParseResponse(responseInBytes, &tmpRes)
	if err != nil {
		return nil, err
	}

	res := make([]string, 0)
	for _, blkRes := range tmpRes {
		res = append(res, blkRes.TxHashes...)
	}

	if len(res) == 0 {
		return nil, fmt.Errorf("not txs found for height %v, shard %v", blkHeight, shardID)
	}

	return res, nil
}

func getRandomTxs(numTxs int, isPrv bool) ([]metadata.Transaction, error) {
	res := make([]metadata.Transaction, 0)
	count := 0
	for count < numTxs {
		shardID := common.RandInt() % common.MaxShardNumber
		bestBlocks, err := ic.GetBestBlock()
		if err != nil {
			return nil, err
		}
		bestBlock := bestBlocks[shardID]
		blockToSample := bestBlock - backwardBlocks + common.RandUint64()%backwardBlocks
		prefix := fmt.Sprintf("[blockToSample %v]", blockToSample)
		txHashes, err := getTxsByBlockNumber(blockToSample, byte(shardID))
		if err != nil {
			continue
		}

		txMap, err := ic.GetTxs(txHashes, false)
		if err != nil {
			//fmt.Printf("%v error: %v\n\n", prefix, err)
			continue
		}
		added := 0
		for _, tmpTx := range txMap {
			tmpIsPRV := tmpTx.GetTokenID().String() == common.PRVIDStr
			if tmpIsPRV != isPrv {
				continue
			}
			if tmpTx.GetVersion() != 2 {
				continue
			}
			switch tmpTx.GetType() {
			case "cv", "tcv":
				continue
			default:
			}

			added++
			res = append(res, tmpTx)
		}
		fmt.Printf("%v added: %v\n", prefix, added)
		count += added
		if count%10 == 0 && added > 0 {
			fmt.Printf("getRandomTxs count = %v\n", count)
		}
	}

	return res, nil
}
