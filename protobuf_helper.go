package main

import (
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/protobuf/proto"
	incCommon "github.com/incognitochain/go-incognito-sdk-v2/common"
	"github.com/incognitochain/go-incognito-sdk-v2/metadata"
	"github.com/thanhn-inc/data-serialization-benchmark/proto_test"
)

func compactBytesToPortalV4ShieldRequest(data []byte) (metadata.Metadata, error) {
	portalShieldMdRes := new(proto_test.PbPortalShieldRequestMeta)
	err := proto.Unmarshal(data, portalShieldMdRes)
	if err == nil && portalShieldMdRes.Type == int32(metadata.PortalV4ShieldingRequestMeta) {
		md := new(metadata.PortalShieldingRequest)
		md.TokenID = portalShieldMdRes.TokenID
		md.Type = int(portalShieldMdRes.Type)
		md.IncAddressStr = portalShieldMdRes.Address
		md.ShieldingProof = base64.StdEncoding.EncodeToString(portalShieldMdRes.Proof)

		return md, nil
	}

	return nil, fmt.Errorf("not a portal shield request")
}

func compactBytesToPortalSubmitConfirmedRequest(data []byte) (metadata.Metadata, error) {
	portalSubmitMdRes := new(proto_test.PbPortalSubmitConfirmedTxMeta)
	err := proto.Unmarshal(data, portalSubmitMdRes)
	if err == nil && portalSubmitMdRes.Type == int32(metadata.PortalV4SubmitConfirmedTxMeta) {
		md := new(metadata.PortalSubmitConfirmedTxRequest)
		md.TokenID = portalSubmitMdRes.TokenID
		md.Type = int(portalSubmitMdRes.Type)
		md.BatchID = portalSubmitMdRes.BatchID
		md.UnshieldProof = base64.StdEncoding.EncodeToString(portalSubmitMdRes.Proof)

		return md, nil
	}

	return nil, fmt.Errorf("not a portal submit confirmedTx request")
}

func compactBytesToIssuingEVMRequest(data []byte) (metadata.Metadata, error) {
	issuingEVMRequest := new(proto_test.PbIssuingEVMRequest)
	err := proto.Unmarshal(data, issuingEVMRequest)
	if err == nil {
		switch issuingEVMRequest.Type {
		case int32(metadata.IssuingBSCRequestMeta), int32(metadata.IssuingETHRequestMeta),
			int32(metadata.IssuingPRVERC20RequestMeta), int32(metadata.IssuingPRVBEP20RequestMeta):
			md := new(metadata.IssuingEVMRequest)
			md.BlockHash = common.BytesToHash(issuingEVMRequest.BlockHash)
			md.TxIndex = uint(issuingEVMRequest.TxIndex)
			md.Type = int(issuingEVMRequest.Type)
			tokenID := new(incCommon.Hash)
			err = tokenID.SetBytes(issuingEVMRequest.TokenID)
			if err != nil {
				return nil, err
			}
			md.IncTokenID = *tokenID
			proofStrs := make([]string, 0)
			for _, proof := range issuingEVMRequest.Proofs {
				proofStrs = append(proofStrs, base64.StdEncoding.EncodeToString(proof))
			}
			md.ProofStrs = proofStrs

			return md, nil
		}
	}

	return nil, fmt.Errorf("not a portal submit confirmedTx request")
}
