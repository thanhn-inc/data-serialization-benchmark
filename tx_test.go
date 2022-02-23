package main

import (
	"fmt"
	"testing"
)

func Test_sampleData(t *testing.T) {
	err := sampleData(5000, true)
	if err != nil {
		panic(err)
	}

	err = sampleData(5000, false)
	if err != nil {
		panic(err)
	}
}

func Test_loadTxs(t *testing.T) {
	txs, err := loadTxs(true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("numTxs: %v\n", len(txs))
}
