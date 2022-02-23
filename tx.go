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

const rootDataDir = "./data"

func sampleData(numTxs int, isPrv bool) error {
	randomTxs, err := getRandomTxs(numTxs, isPrv)
	if err != nil {
		return err
	}

	var dataDir string
	if isPrv {
		dataDir = fmt.Sprintf("%v/prv", rootDataDir)
	} else {
		dataDir = fmt.Sprintf("%v/token", rootDataDir)
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

func loadTxs(isPrv bool) ([]metadata.Transaction, error) {
	res := make([]metadata.Transaction, 0)

	var dataDir string
	if isPrv {
		dataDir = rootDataDir + "/prv"
	} else {
		dataDir = rootDataDir + "/token"
	}
	files, err := ioutil.ReadDir(dataDir)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		filePath := fmt.Sprintf("%v/%v", dataDir, f.Name())
		file, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}

		rawEncodedData := make([]byte, 500000)
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
