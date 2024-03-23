package derive

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hashicorp/go-multierror"
	"math/big"
	"strings"
)

var (
	PPSUpdatedABI    = "PPSUpdated(uint256)"
	PPSUpdatedHash   = crypto.Keccak256Hash([]byte(PPSUpdatedABI))
	PPSEventVersion0 = common.Hash{}
)

var (
	PPSUpdateFnSig      = "updatePPS(uint64)"
	PPSUpdateLen        = 36
	PPSUpdateFuncBytes4 = crypto.Keccak256([]byte(PPSUpdateFnSig))[:4]
	L2PPSContract       = common.HexToAddress("CONTRACT_ADDRESS") // TODO: CONTRACT NEEDS TO BE DEPLOYED ON L2
)

func DerivePPSUpdates(receipts []*types.Receipt, depositContractAddr common.Address) ([]*types.DepositTx, error) {
	var out []*types.DepositTx
	var result error
	for i, rec := range receipts {
		if rec.Status != types.ReceiptStatusSuccessful {
			continue
		}
		for j, log := range rec.Logs {
			if log.Address == depositContractAddr && len(log.Topics) > 0 && log.Topics[0] == PPSUpdatedHash {
				fmt.Println(log)
				dep, err := UnmarshalPPSUpdatedEvent(log)
				if err != nil {
					result = multierror.Append(result, fmt.Errorf("malformatted L1 deposit log in receipt %d, log %d: %w", i, j, err))
				} else {
					out = append(out, dep)
				}
			}
		}
	}

	return out, result
}

type PPSUpdateInfo struct {
	Amount uint64
}

func (info *PPSUpdateInfo) MarshalBinary() ([]byte, error) {
	data := make([]byte, PPSUpdateLen)
	offset := 0
	copy(data[offset:4], PPSUpdateFuncBytes4)
	offset += 4
	binary.BigEndian.PutUint64(data[offset+24:offset+32], info.Amount)
	return data, nil
}

func (info *PPSUpdateInfo) UnmarshalBinary(data []byte) error {
	var padding [24]byte
	offset := 4
	info.Amount = binary.BigEndian.Uint64(data[offset+24 : offset+32])
	if !bytes.Equal(data[offset:offset+24], padding[:]) {
		return fmt.Errorf(" exceeds uint64 bounds: %x", data[offset:offset+32])
	}
	return nil
}

func PPSUpdateTxData(data []byte) (PPSUpdateInfo, error) {
	var info PPSUpdateInfo
	err := info.UnmarshalBinary(data)
	return info, err
}

func UnmarshalPPSUpdatedEvent(ev *types.Log) (*types.DepositTx, error) {

	eventAbi, _ := abi.JSON(strings.NewReader(bindings.BindingsMetaData.ABI))
	from := common.BytesToAddress(ev.Topics[1][12:])

	event := struct {
		Amount uint32
	}{}
	err := eventAbi.UnpackIntoInterface(&event, "PPSUpdated", ev.Data)
	if err != nil {
		return nil, err
	}

	infoDat := PPSUpdateInfo{
		Amount: uint64(event.Amount),
	}
	data, err := infoDat.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return &types.DepositTx{
		SourceHash:          common.Hash{},
		From:                from,
		To:                  &L2PPSContract,
		Mint:                nil,
		Value:               big.NewInt(0),
		Gas:                 150_000_000,
		IsSystemTransaction: true,
		Data:                data,
	}, nil
}

func FormatPPSBytes(out []*types.DepositTx, sysCfg eth.SystemConfig, seqNumber uint64, block eth.BlockInfo) ([]hexutil.Bytes, error) {
	source := L1InfoDepositSource{
		L1BlockHash: block.Hash(),
		SeqNumber:   seqNumber,
	}
	encodedTxs := make([]hexutil.Bytes, 0, len(out))
	for i, tx := range out {
		tx.SourceHash = source.SourceHash()
		opaqueTx, err := types.NewTx(tx).MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("failed to encode user tx %d", i)
		} else {
			encodedTxs = append(encodedTxs, opaqueTx)
		}
	}
	return encodedTxs, nil
}
