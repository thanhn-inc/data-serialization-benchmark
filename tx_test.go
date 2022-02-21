package main

import (
	"fmt"
	"testing"
)

func Test_sampleData(t *testing.T) {
	err := sampleData(20)
	if err != nil {
		panic(err)
	}
}

func Test_loadTxs(t *testing.T) {
	txs, err := loadTxs()
	if err != nil {
		panic(err)
	}

	fmt.Printf("numTxs: %v\n", len(txs))
}
