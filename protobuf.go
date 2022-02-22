package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/incognitochain/go-incognito-sdk-v2/coin"
	"github.com/incognitochain/go-incognito-sdk-v2/common"
	"github.com/incognitochain/go-incognito-sdk-v2/crypto"
	"github.com/incognitochain/go-incognito-sdk-v2/metadata"
	"github.com/incognitochain/go-incognito-sdk-v2/privacy"
	"github.com/incognitochain/go-incognito-sdk-v2/privacy/v2/bulletproofs"
	"github.com/incognitochain/go-incognito-sdk-v2/transaction/tx_ver2"
	"github.com/thanhn-inc/data-serialization-benchmark/proto_test"
)

func ScalarToProtoBuf(sc *crypto.Scalar) *proto_test.PbScalar {
	if sc == nil {
		return nil
	}
	return &proto_test.PbScalar{Key: sc.ToBytesS()}
}

func ProtoBufToScalar(protoSc *proto_test.PbScalar) *crypto.Scalar {
	if protoSc == nil {
		return nil
	}
	return new(crypto.Scalar).FromBytesS(protoSc.Key)
}

func PointToProtoBuf(p *crypto.Point) *proto_test.PbPoint {
	if p == nil {
		return nil
	}
	return &proto_test.PbPoint{Key: p.ToBytesS()}
}

func ProtoBufToPoint(protoPoint *proto_test.PbPoint) *crypto.Point {
	if protoPoint == nil {
		return nil
	}
	res, _ := new(crypto.Point).FromBytesS(protoPoint.Key)
	return res
}

func CoinV2ToProtoBuf(c *coin.CoinV2) *proto_test.PbCoinV2 {
	res := new(proto_test.PbCoinV2)
	res.Version = int32(c.GetVersion())
	res.Info = c.GetInfo()
	res.PublicKey = PointToProtoBuf(c.GetPublicKey())
	res.Commitment = PointToProtoBuf(c.GetCommitment())
	res.KeyImage = PointToProtoBuf(c.GetKeyImage())
	res.SharedConcealRandom = ScalarToProtoBuf(c.GetSharedConcealRandom())
	res.SharedRandom = ScalarToProtoBuf(c.GetSharedRandom())
	if c.GetTxRandom() != nil {
		res.TxRandom = c.GetTxRandom().Bytes()
	}
	res.Mask = ScalarToProtoBuf(c.GetRandomness())
	res.Amount = ScalarToProtoBuf(c.GetAmount())
	res.AssetTag = PointToProtoBuf(c.GetAssetTag())

	return res
}

func InnerProofToProtoBuf(p *bulletproofs.InnerProductProof) *proto_test.PbInnerProductProof {
	res := new(proto_test.PbInnerProductProof)

	L := make([]*proto_test.PbPoint, 0)
	for _, lElem := range p.L() {
		L = append(L, PointToProtoBuf(lElem))
	}
	res.L = L

	R := make([]*proto_test.PbPoint, 0)
	for _, rElem := range p.R() {
		R = append(R, PointToProtoBuf(rElem))
	}
	res.R = R

	res.A = ScalarToProtoBuf(p.A())
	res.B = ScalarToProtoBuf(p.B())
	res.P = PointToProtoBuf(p.P())

	return res
}

func RangeProofToProtoBuf(p *bulletproofs.RangeProof) *proto_test.PbRangeProof {
	res := new(proto_test.PbRangeProof)

	res.CmsValues = make([]*proto_test.PbPoint, 0)
	for _, cmsValue := range p.GetCmsValues() {
		res.CmsValues = append(res.CmsValues, PointToProtoBuf(cmsValue))
	}

	res.A = PointToProtoBuf(p.A())
	res.S = PointToProtoBuf(p.S())
	res.T1 = PointToProtoBuf(p.T1())
	res.T2 = PointToProtoBuf(p.T2())
	res.TauX = ScalarToProtoBuf(p.TauX())
	res.THat = ScalarToProtoBuf(p.THat())
	res.Mu = ScalarToProtoBuf(p.Mu())
	res.InnerProof = InnerProofToProtoBuf(p.InnerProof())

	return res
}

func ProofV2ToProtoBuf(p *privacy.ProofV2) (*proto_test.PbProofV2, error) {
	res := new(proto_test.PbProofV2)

	res.Version = int32(p.Version)

	tmpRangeProof, ok := p.GetRangeProof().(*bulletproofs.RangeProof)
	if !ok {
		return nil, fmt.Errorf("not a valid bulletProofs")
	}
	res.RangeProof = RangeProofToProtoBuf(tmpRangeProof)

	res.InputCoins = make([]*proto_test.PbCoinV2, 0)
	for i, inCoin := range p.GetInputCoins() {
		inCoinV2, ok := inCoin.(*coin.CoinV2)
		if !ok {
			return nil, fmt.Errorf("input coin %v is not a CoinV2", i)
		}
		res.InputCoins = append(res.InputCoins, CoinV2ToProtoBuf(inCoinV2))
	}

	res.OutputCoins = make([]*proto_test.PbCoinV2, 0)
	for i, outCoin := range p.GetOutputCoins() {
		outCoinV2, ok := outCoin.(*coin.CoinV2)
		if !ok {
			return nil, fmt.Errorf("output coin %v is not a CoinV2", i)
		}
		res.OutputCoins = append(res.OutputCoins, CoinV2ToProtoBuf(outCoinV2))
	}

	tmpBytes := p.Bytes()
	protoBytes, _ := proto.Marshal(res)
	fmt.Printf("jsbLength: %v, protoLength: %v\n", len(tmpBytes), len(protoBytes))
	return res, nil
}

func TxToProtoBuf(tx *tx_ver2.Tx) (*proto_test.PbTxVer2, error) {
	if tx.GetType() == "cv" || tx.GetType() == "tcv" {
		return nil, fmt.Errorf("tx type %v not supported", tx.GetType())
	}

	res := new(proto_test.PbTxVer2)
	res.Version = int32(tx.Version)
	res.Type = tx.Type
	res.LockTime = tx.LockTime
	res.Fee = tx.Fee
	res.Info = tx.Info
	res.SigPubKey = tx.SigPubKey
	res.Sig = tx.Sig
	res.LastByte = int32(tx.PubKeyLastByteSender)
	if tx.GetMetadata() != nil {
		res.Metadata, _ = json.Marshal(tx.GetMetadata())
	}
	if tx.GetProof() != nil {
		res.Proof = tx.Proof.Bytes()
	}

	//proofV2, _ := tx.GetProof().(*privacy.ProofV2)
	//var err error
	//res.Proof, err = ProofV2ToProtoBuf(proofV2)
	//if err != nil {
	//	return nil, err
	//}

	return res, nil
}

func ProtoBufToTx(protoTx *proto_test.PbTxVer2) (*tx_ver2.Tx, error) {
	res := new(tx_ver2.Tx)
	var err error

	if len(protoTx.Proof) != 0 {
		proof := new(privacy.ProofV2)
		err = proof.SetBytes(protoTx.Proof)
		if err != nil {
			return nil, err
		}
		res.Proof = proof
	}

	res.Version = int8(protoTx.Version)
	res.Type = protoTx.Type
	res.LockTime = protoTx.LockTime
	res.Fee = protoTx.Fee
	res.Info = protoTx.Info
	res.SigPubKey = protoTx.SigPubKey
	res.Sig = protoTx.Sig
	res.PubKeyLastByteSender = byte(protoTx.LastByte)
	res.Metadata, err = metadata.ParseMetadata(protoTx.Metadata)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func TokenDataV2ToProtoBuf(tokenData tx_ver2.TxTokenDataVersion2) *proto_test.PbTxTokenDataVersion2 {
	res := new(proto_test.PbTxTokenDataVersion2)
	res.ID = tokenData.PropertyID.GetBytes()
	res.Name = tokenData.PropertyName
	res.Symbol = tokenData.PropertySymbol
	res.SigPubKey = tokenData.SigPubKey
	res.Sig = tokenData.Sig
	res.Proof = tokenData.Proof.Bytes()
	res.Type = int32(tokenData.Type)
	res.Mintable = tokenData.Mintable

	return res
}

func ProtoBufToTokenDataV2(protoTokenData *proto_test.PbTxTokenDataVersion2) (*tx_ver2.TxTokenDataVersion2, error) {
	res := new(tx_ver2.TxTokenDataVersion2)

	tokenID := new(common.Hash)
	err := tokenID.SetBytes(protoTokenData.ID)
	if err != nil {
		return nil, err
	}
	res.PropertyID = *tokenID

	res.PropertyName = protoTokenData.Name
	res.PropertySymbol = protoTokenData.Symbol
	res.SigPubKey = protoTokenData.SigPubKey
	res.Sig = protoTokenData.Sig
	res.Type = int(protoTokenData.Type)
	res.Mintable = protoTokenData.Mintable

	if len(protoTokenData.Proof) != 0 {
		proof := new(privacy.ProofV2)
		err = proof.SetBytes(protoTokenData.Proof)
		if err != nil {
			return nil, err
		}
		res.Proof = proof
	}

	return res, nil
}

func TxTokenToProtoBuf(tx *tx_ver2.TxToken) (*proto_test.PbTxTokenVer2, error) {
	if tx.GetType() == "cv" || tx.GetType() == "tcv" {
		return nil, fmt.Errorf("tx type %v not supported", tx.GetType())
	}

	var err error
	res := new(proto_test.PbTxTokenVer2)

	res.Tx, err = TxToProtoBuf(tx.GetTxBase().(*tx_ver2.Tx))
	if err != nil {
		return nil, err
	}

	res.TokenData = TokenDataV2ToProtoBuf(tx.TokenData)

	return res, nil
}

func ProtoBufToTxToken(protoTxToken *proto_test.PbTxTokenVer2) (*tx_ver2.TxToken, error) {
	res := new(tx_ver2.TxToken)

	tmpTx, err := ProtoBufToTx(protoTxToken.Tx)
	if err != nil {
		return nil, err
	}
	res.Tx = *tmpTx

	tmpTokenData, err := ProtoBufToTokenDataV2(protoTxToken.TokenData)
	if err != nil {
		return nil, err
	}
	res.TokenData = *tmpTokenData

	return res, nil
}
