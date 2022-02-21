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

func getRandomTxs(numTxs int) ([]metadata.Transaction, error) {
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
		txHashes, err := getTxsByBlockNumber(blockToSample, byte(shardID))
		if err != nil {
			continue
		}

		txMap, err := ic.GetTxs(txHashes, true)
		if err != nil {
			continue
		}
		for _, tmpTx := range txMap {
			res = append(res, tmpTx)
		}
		count += len(txMap)
		if count%10 == 0 && count > 0 {
			fmt.Printf("getRandomTxs count = %v\n", count)
		}
	}

	return res, nil
}
