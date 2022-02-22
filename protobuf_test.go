package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/incognitochain/go-incognito-sdk-v2/transaction/tx_ver2"
	"github.com/thanhn-inc/data-serialization-benchmark/proto_test"
	"testing"
)

func Test_TxToProtoBuf(t *testing.T) {
	txs, err := loadTxs(true)
	if err != nil {
		panic(err)
	}

	protoTxs := make([]*proto_test.PbTxVer2, 0)
	minRate := 1000.0
	maxRate := 0.0
	totalRate := float64(0)
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
		protoTxs = append(protoTxs, protoTx)

		jsb, _ := json.Marshal(txV2)
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
	}
	fmt.Printf("minRate: %v, maxRate: %v, avgRate: %v\n", minRate, maxRate, totalRate/float64(len(protoTxs)))
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
			panic(err)
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
