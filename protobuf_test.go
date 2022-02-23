package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/incognitochain/go-incognito-sdk-v2/transaction/tx_ver2"
	"github.com/thanhn-inc/data-serialization-benchmark/proto_test"
	"testing"
	"time"
)

func Test_TxToProtoBuf(t *testing.T) {
	txs, err := loadTxs(true)
	if err != nil {
		panic(err)
	}

	fmt.Println("LOAD TXS successfully!!!!")

	protoTxs := make([]*proto_test.PbTxVer2, 0)
	minRate := 1000.0
	maxRate := 0.0
	totalRate := float64(0)
	minTxHash := ""
	maxTxHash := ""
	for i, tx := range txs {
		prefix := fmt.Sprintf("[i: %v, txHash: %v]", i, tx.Hash().String()[:10])
		txV2, ok := tx.(*tx_ver2.Tx)
		if !ok {
			continue
		}
		protoTx, err := TxToProtoBuf(txV2)
		if err != nil {
			panic(fmt.Sprintf("%v error: %v", prefix, err))
		}
		protoTxs = append(protoTxs, protoTx)

		jsb, _ := json.Marshal(txV2)
		protoBytes, err := proto.Marshal(protoTx)
		rate := float64(len(protoBytes)) / float64(len(jsb))
		totalRate += rate
		if rate < minRate {
			minRate = rate
			minTxHash = tx.Hash().String()
		}
		if rate > maxRate {
			maxRate = rate
			maxTxHash = tx.Hash().String()
		}
		//fmt.Println(string(jsb))
		fmt.Printf("%v jsonLength: %v, protoLength: %v, rate: %v\n\n", prefix, len(jsb), len(protoBytes), rate)
	}
	fmt.Printf("minRate: %v(%v), maxRate: %v(%v), avgRate: %v\n", minRate, minTxHash, maxRate, maxTxHash,
		totalRate/float64(len(protoTxs)))
}

func Test_TxToCompactBytes(t *testing.T) {
	txs, err := loadTxs(true)
	if err != nil {
		panic(err)
	}

	fmt.Println("LOAD TXS successfully!!!!")

	txCount := 0
	minRate := 1000.0
	maxRate := 0.0
	totalRate := float64(0)
	minSizeTxHash := ""
	maxSizeTxHash := ""
	minTimeRate := 1000.0
	maxTimeRate := 0.0
	totalTimeRate := 0.0
	for i, tx := range txs {
		prefix := fmt.Sprintf("[i: %v, txHash: %v]", i, tx.Hash().String()[:10])
		txV2, ok := tx.(*tx_ver2.Tx)
		if !ok {
			continue
		}

		start := time.Now()
		compactBytes, err := TxToCompactBytes(txV2)
		if err != nil {
			panic(err)
		}
		protoTime := time.Since(start).Seconds()

		start = time.Now()
		jsb, _ := json.Marshal(txV2)
		jsbTime := time.Since(start).Seconds()

		timeRate := protoTime / jsbTime
		if timeRate > maxTimeRate {
			maxTimeRate = timeRate
		}
		if timeRate < minTimeRate {
			minTimeRate = timeRate
		}
		totalTimeRate += timeRate

		rate := float64(len(compactBytes)) / float64(len(jsb))
		totalRate += rate
		if rate < minRate {
			minRate = rate
			minSizeTxHash = tx.Hash().String()
		}
		if rate > maxRate {
			maxRate = rate
			maxSizeTxHash = tx.Hash().String()
		}
		//fmt.Println(string(jsb))
		fmt.Printf("%v jsonLength: %v, protoLength: %v, rate: %v\n", prefix, len(jsb), len(compactBytes), rate)
		fmt.Printf("%v jsbTime: %v, protoTime: %v, rate: %v\n\n", prefix, jsbTime, protoTime, timeRate)
		txCount++
	}
	fmt.Printf("minRate: %v(%v), maxRate: %v(%v), avgRate: %v\n", minRate, minSizeTxHash, maxRate, maxSizeTxHash,
		totalRate/float64(txCount))
	fmt.Printf("minTimeRate: %v, maxTimeRate: %v, avgTimeRate: %v\n", minTimeRate, maxTimeRate, totalTimeRate/float64(txCount))
}

func Test_CompactBytesToTx(t *testing.T) {
	txs, err := loadTxs(true)
	if err != nil {
		panic(err)
	}

	fmt.Println("LOAD TXS successfully!!!!")
	for i := 0; i < 10; i++ {
		tx := txs[i]
		prefix := fmt.Sprintf("[i: %v, txHash: %v]", i, tx.Hash().String()[:10])
		txV2, ok := tx.(*tx_ver2.Tx)
		if !ok {
			continue
		}

		start := time.Now()
		compactBytes, err := TxToCompactBytes(txV2)
		if err != nil {
			panic(err)
		}
		encodingTime := time.Since(start).Seconds()
		start = time.Now()
		newTx, err := CompactBytesToTx(compactBytes)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v encodingTime: %v decodingTime: %v\n", prefix, encodingTime, time.Since(start).Seconds())

		if newTx.Hash().String() != tx.Hash().String() {
			jsb1, _ := json.Marshal(tx)
			jsb2, _ := json.Marshal(newTx)
			fmt.Println(string(jsb1))
			fmt.Println(string(jsb2))
			panic(fmt.Sprintf("%v expected txHash %v, got %v", prefix, tx.Hash().String(), newTx.Hash().String()))
		}
	}
}

func Test_ProtoBufToTx(t *testing.T) {
	txs, err := loadTxs(true)
	if err != nil {
		panic(err)
	}

	protoTxs := make([]*proto_test.PbTxVer2, 0)
	for i, tx := range txs {
		prefix := fmt.Sprintf("[i: %v, txHash: %v]", i, tx.Hash().String()[:10])
		txV2, ok := tx.(*tx_ver2.Tx)
		if !ok {
			continue
		}
		protoTx, err := TxToProtoBuf(txV2)
		if err != nil {
			panic(err)
		}

		tmpTx, err := ProtoBufToTx(protoTx)
		if err != nil {
			panic(fmt.Sprintf("%v error: %v", prefix, err))
		}

		jsb, _ := json.Marshal(txV2)
		jsb2, _ := json.Marshal(tmpTx)
		if !bytes.Equal(jsb, jsb2) {
			fmt.Println(string(jsb))
			fmt.Println(string(jsb2))
			panic(prefix)
		}

		protoTxs = append(protoTxs, protoTx)
	}
}

func Test_TxTokenToProtoBuf(t *testing.T) {
	txs, err := loadTxs(false)
	if err != nil {
		panic(err)
	}

	protoTxs := make([]*proto_test.PbTxTokenVer2, 0)
	minRate := 1000.0
	maxRate := 0.0
	totalRate := float64(0)
	for i, tx := range txs {
		prefix := fmt.Sprintf("[i: %v, txHash: %v]", i, tx.Hash().String()[:10])
		txTokenV2, ok := tx.(*tx_ver2.TxToken)
		if !ok {
			continue
		}
		protoTx, err := TxTokenToProtoBuf(txTokenV2)
		if err != nil {
			panic(err)
		}

		jsb, _ := json.Marshal(txTokenV2)
		protoBytes, err := proto.Marshal(protoTx)
		rate := float64(len(protoBytes)) / float64(len(jsb))
		totalRate += rate
		if rate < minRate {
			minRate = rate
		}
		if rate > maxRate {
			maxRate = rate
		}

		fmt.Println(string(jsb))
		fmt.Printf("%v jsonLength: %v, protoLength: %v, rate: %v\n\n", prefix, len(jsb), len(protoBytes), rate)

		protoTxs = append(protoTxs, protoTx)
	}

	fmt.Printf("minRate: %v, maxRate: %v, avgRate: %v\n", minRate, maxRate, totalRate/float64(len(protoTxs)))
}

func Test_ProtoBufToTxToken(t *testing.T) {
	txs, err := loadTxs(false)
	if err != nil {
		panic(err)
	}

	protoTxs := make([]*proto_test.PbTxTokenVer2, 0)
	for i, tx := range txs {
		prefix := fmt.Sprintf("[i: %v, txHash: %v]", i, tx.Hash().String()[:10])
		txTokenV2, ok := tx.(*tx_ver2.TxToken)
		if !ok {
			continue
		}
		protoTx, err := TxTokenToProtoBuf(txTokenV2)
		if err != nil {
			panic(err)
		}

		tmpTx, err := ProtoBufToTxToken(protoTx)
		if err != nil {
			panic(err)
		}

		jsb, _ := json.Marshal(txTokenV2)
		jsb2, _ := json.Marshal(tmpTx)
		if !bytes.Equal(jsb, jsb2) {
			fmt.Println(string(jsb))
			fmt.Println(string(jsb2))
			panic(prefix)
		}
		protoTxs = append(protoTxs, protoTx)
	}
}
