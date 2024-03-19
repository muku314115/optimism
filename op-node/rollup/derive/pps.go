package derive

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hashicorp/go-multierror"
	"github.com/holiman/uint256"
)

var (
	PPSUpdatedABI    = "PPSUpdated(uint256)"
	PPSUpdatedHash   = crypto.Keccak256Hash([]byte(PPSUpdatedABI))
	PPSEventVersion0 = common.Hash{}
)

func DerivePPSUpdates(receipts []*types.Receipt, depositContractAddr common.Address) ([]hexutil.Bytes, error) {
	var out []*types.LegacyTx
	var result error
	for i, rec := range receipts {
		if rec.Status != types.ReceiptStatusSuccessful {
			continue
		}
		for j, log := range rec.Logs {
			if log.Address == depositContractAddr && len(log.Topics) > 0 && log.Topics[0] == PPSUpdatedHash {
				dep, err := UnmarshalPPSUpdatedEvent(log)
				if err != nil {
					result = multierror.Append(result, fmt.Errorf("malformatted L1 deposit log in receipt %d, log %d: %w", i, j, err))
				} else {
					out = append(out, dep)
				}
			}
		}
	}
	//TODO: CONVERT THE RETURNED LEGACY TXS TO HEXUTIL BYTES
	return out, result
}

func UnmarshalPPSUpdatedEvent(ev *types.Log) (*types.LegacyTx, error) {
	// indexed 0
	from := common.BytesToAddress(ev.Topics[1][12:])
	// indexed 1
	to := common.BytesToAddress(ev.Topics[2][12:])
	// indexed 2
	var opaqueContentOffset uint256.Int
	opaqueContentOffset.SetBytes(ev.Data[0:32])
	if !opaqueContentOffset.IsUint64() || opaqueContentOffset.Uint64() != 32 {
		return nil, fmt.Errorf("invalid opaqueData slice header offset: %d", opaqueContentOffset.Uint64())
	}
	// The next 32 bytes indicate the length of the opaqueData content.
	var opaqueContentLength uint256.Int
	opaqueContentLength.SetBytes(ev.Data[32:64])
	// Make sure the length is an uint64, it's not larger than the remaining data, and the log is using minimal padding (i.e. can't add 32 bytes without exceeding data)
	if !opaqueContentLength.IsUint64() || opaqueContentLength.Uint64() > uint64(len(ev.Data)-64) || opaqueContentLength.Uint64()+32 <= uint64(len(ev.Data)-64) {
		return nil, fmt.Errorf("invalid opaqueData slice header length: %d", opaqueContentLength.Uint64())
	}

	// TODO: USE THIS TO PARSE THE EVENT USING ABI
	event := struct {
		Amount uint32
	}{}

	//TODO: CONSTRUCT LEGACY TX TO REPRESENT THE SMART CONTRACT CALL
	dep := &types.LegacyTx{}

	return dep, nil
}
