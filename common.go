package main

import "github.com/incognitochain/go-incognito-sdk-v2/incclient"

var ic *incclient.IncClient

func init() {
	var err error
	ic, err = incclient.NewMainNetClient()
	if err != nil {
		panic(err)
	}
}
