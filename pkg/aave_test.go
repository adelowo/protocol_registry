package pkg

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestAave_GenerateCalldata_Forks(t *testing.T) {

	aave, err := NewAaveOperation(AaveProtocolForkAave)
	require.NoError(t, err)

	require.Equal(t, AaveV3, aave.Name())

	aave, err = NewAaveOperation(AaveProtocolForkSpark)
	require.NoError(t, err)

	require.Equal(t, SparkLend, aave.Name())
}

func TestAave_GenerateCalldata_Withdraw(t *testing.T) {
	// cast calldata "withdraw(address,uint256,address)" 0xc0ffee254729296a45a3885639AC7E10F9d54979 500000000000000000 0x0000000000000000000000000000000000000000
	// 0x69328dec000000000000000000000000c0ffee254729296a45a3885639ac7e10f9d5497900000000000000000000000000000000000000000000000006f05b59d3b200000000000000000000000000000000000000000000000000000000000000000000

	expectedCalldata := "0x69328dec000000000000000000000000c0ffee254729296a45a3885639ac7e10f9d5497900000000000000000000000000000000000000000000000006f05b59d3b200000000000000000000000000000000000000000000000000000000000000000000"

	aave, err := NewAaveOperation(AaveProtocolForkAave)
	require.NoError(t, err)

	calldata, err := aave.GenerateCalldata(LoanWithdraw, GenerateCalldataOptions{
		Asset:  common.HexToAddress("0xc0ffee254729296a45a3885639AC7E10F9d54979"),
		Amount: big.NewInt(500000000000000000),
		Sender: common.HexToAddress("0x0000000000000000000000000000000000000000"),
	})
	require.NoError(t, err)

	require.Equal(t, expectedCalldata, calldata)
}

func TestAave_GenerateCalldata_Supply(t *testing.T) {
	// cast calldata "supply(address,uint256,address,uint16)" 0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984 1000000000000000000 0x0000000000000000000000000000000000000000 10
	// 0x617ba0370000000000000000000000001f9840a85d5af5bf1d1762f925bdaddc4201f9840000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a

	expectedCalldata := "0x617ba0370000000000000000000000001f9840a85d5af5bf1d1762f925bdaddc4201f9840000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a"

	aave, err := NewAaveOperation(AaveProtocolForkAave)
	require.NoError(t, err)

	calldata, err := aave.GenerateCalldata(LoanSupply, GenerateCalldataOptions{
		Asset:       common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"),
		Amount:      big.NewInt(1000000000000000000),
		Sender:      common.HexToAddress("0x0000000000000000000000000000000000000000"),
		ReferalCode: uint16(10),
	})
	require.NoError(t, err)

	require.Equal(t, expectedCalldata, calldata)
}

func TestAave_GenerateCalldataUnspportedAction(t *testing.T) {

	aave, err := NewAaveOperation(AaveProtocolForkAave)
	require.NoError(t, err)

	_, err = aave.GenerateCalldata(NativeStake, GenerateCalldataOptions{
		Sender: common.HexToAddress("0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6"),
	})
	require.Error(t, err)
}
