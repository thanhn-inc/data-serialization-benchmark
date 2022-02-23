package main

import (
	"fmt"
	"github.com/incognitochain/go-incognito-sdk-v2/common"
	"github.com/incognitochain/go-incognito-sdk-v2/common/base58"
	"github.com/incognitochain/go-incognito-sdk-v2/metadata"
	"github.com/incognitochain/go-incognito-sdk-v2/rpchandler"
	"github.com/incognitochain/go-incognito-sdk-v2/wallet"
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
	mapOTACoinLengths, err := ic.GetOTACoinLength()
	if err != nil {
		return nil, err
	}

	tokenID := common.PRVCoinID
	otaCoinLength := mapOTACoinLengths[common.PRVIDStr]
	if !isPrv {
		otaCoinLength = mapOTACoinLengths[common.ConfidentialAssetID.String()]
		tokenID = common.ConfidentialAssetID
	}

	addedTxs := make(map[string]bool)
	for count < numTxs {
		shardID := common.RandInt() % common.MaxShardNumber
		currentOTANums := otaCoinLength[byte(shardID)]
		prefix := fmt.Sprintf("[Shard %v, isPRV %v]", shardID, isPrv)

		idxList := make([]uint64, 0)
		for i := 0; i < pageSize; i++ {
			idxList = append(idxList, common.RandUint64()%currentOTANums)
		}
		otaCoins, err := ic.GetOTACoinsByIndices(byte(shardID), tokenID.String(), idxList)
		if err != nil {
			return nil, err
		}

		pkList := make([]string, 0)
		for _, otaCoin := range otaCoins {
			if wallet.IsPublicKeyBurningAddress(otaCoin.GetPublicKey().ToBytesS()) {
				continue
			}
			pkList = append(pkList, base58.Base58Check{}.Encode(otaCoin.GetPublicKey().ToBytesS(), 0))
		}

		txHashes, err := ic.GetTxHashByPublicKeys(pkList)
		if err != nil {
			return nil, err
		}
		txList := make([]string, 0)
		for _, txHashList := range txHashes {
			for _, txHash := range txHashList {
				if addedTxs[txHash] {
					continue
				}
				txList = append(txList, txHash)
				addedTxs[txHash] = true
			}
		}

		txMap, err := ic.GetTxs(txList, true)
		if err != nil {
			fmt.Printf("%v GetTxs error: %v\n", prefix, err)
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
