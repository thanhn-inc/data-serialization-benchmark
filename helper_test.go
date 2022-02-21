package main

import (
	"fmt"
	"github.com/incognitochain/go-incognito-sdk-v2/common"
	"testing"
)

const (
	numTests = 10
)

func Test_getTxsByBlockNumber(t *testing.T) {
	for i := 0; i < numTests; i++ {
		shardID := common.RandInt() % common.MaxShardNumber
		bestBlocks, err := ic.GetBestBlock()
		if err != nil {
			panic(err)
		}
		bestBlock := bestBlocks[shardID]
		blockToSample := bestBlock + 1000 + common.RandUint64()%1000
		prefix := fmt.Sprintf("[i: %v, shard %v, block %v]\n", i, shardID, blockToSample)
		txs, err := getTxsByBlockNumber(blockToSample, byte(shardID))
		if err != nil {
			fmt.Printf("%v error: %v\n\n", prefix, err)
			continue
		}
		fmt.Printf("%v txs: %v\n\n", prefix, txs)
	}
}
