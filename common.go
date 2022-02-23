package main

import (
	"github.com/incognitochain/go-incognito-sdk-v2/common"
	"github.com/incognitochain/go-incognito-sdk-v2/incclient"
)

var ic *incclient.IncClient

func init() {
	var err error
	ic, err = incclient.NewMainNetClient()
	if err != nil {
		panic(err)
	}

	txInfoPlaceHolder = common.SHA256([]byte("This is for the empty tx info"))[:8]
}

var txInfoPlaceHolder []byte
