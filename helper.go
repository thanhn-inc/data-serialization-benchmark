package main

import (
	"fmt"
	"github.com/incognitochain/go-incognito-sdk-v2/rpchandler"
)

func getTxsByBlockNumber(blkHeight uint64, shardID byte) ([]string, error) {
	type RPCResult struct {
		TxHashes []string
	}

	rpcMethod := "retrieveblockbyheight"
	responseInBytes, err := ic.NewRPCCall("", rpcMethod, []interface{}{blkHeight, shardID, "0"}, 1)
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
