package main

import (
	"encoding/json"
	"fmt"
	"github.com/incognitochain/go-incognito-sdk-v2/common/base58"
	"github.com/incognitochain/go-incognito-sdk-v2/metadata"
	"github.com/incognitochain/go-incognito-sdk-v2/transaction"
	"io/ioutil"
	"os"
)

var dataDir = "./data"

func sampleData(numTxs int) error {
	randomTxs, err := getRandomTxs(numTxs)
	if err != nil {
		return err
	}

	for _, tx := range randomTxs {
		txHash := tx.Hash().String()
		filePath := fmt.Sprintf("%v/%v.dat", dataDir, txHash[:10])
		f, err := os.Create(filePath)
		if err != nil {
			return err
		}

		toBeWrittenBytes, _ := json.Marshal(tx)
		toBeWritten := base58.Base58Check{}.Encode(toBeWrittenBytes, 0)
		_, err = f.WriteString(toBeWritten)
		if err != nil {
			return err
		}

		_ = f.Close()
	}

	return nil
}

func loadTxs() ([]metadata.Transaction, error) {
	res := make([]metadata.Transaction, 0)

	files, err := ioutil.ReadDir("./data")
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		filePath := fmt.Sprintf("%v/%v", dataDir, f.Name())
		file, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}

		rawEncodedData := make([]byte, 100000)
		n, err := file.Read(rawEncodedData)
		if err != nil {
			return nil, err
		}
		rawTx, _, err := base58.Base58Check{}.Decode(string(rawEncodedData[:n]))
		if err != nil {
			fmt.Println(filePath, n, err)
			return nil, err
		}

		txChoice, err := transaction.DeserializeTransactionJSON(rawTx)
		if err != nil {
			return nil, err
		}

		res = append(res, txChoice.ToTx())
	}

	return res, nil
}
