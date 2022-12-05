// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minNumParticipants\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNumParticipants\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"DisputeResolutionBegan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"DisputeResolutionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slashedAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipientAddr\",\"type\":\"address\"}],\"name\":\"Slashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"AVAILABLE_NODE_TYPES\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sequencerAddr\",\"type\":\"address\"}],\"name\":\"BeginDisputeResolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_MIN_SLASH_PERCENTAGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_STAKING_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sequencerAddr\",\"type\":\"address\"}],\"name\":\"EndDisputeResolution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetAvailableNodeTypes\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"GetCurrentAccountStakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCurrentDisputeWatchtowers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCurrentParticipants\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCurrentSequencers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCurrentSequencersInProbation\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCurrentStakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCurrentStakingThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCurrentValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetCurrentWatchtowers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"watchtowerAddr\",\"type\":\"address\"}],\"name\":\"GetDisputedSequencerAddrs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sequencerAddr\",\"type\":\"address\"}],\"name\":\"GetDisputedWatchtowerAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sequencerAddr\",\"type\":\"address\"}],\"name\":\"GetIsSequencerInProbation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetMaxNumParticipants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetMaxNumSequencers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetMaxNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetMaxNumWatchtowers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetMinNumParticipants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetMinNumSequencers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetMinNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetMinNumWatchtowers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetSlashPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"IsSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"IsValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"IsWatchtower\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NODE_SEQUENCER\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NODE_VALIDATOR\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NODE_WATCHTOWER\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maximumNumParticipants\",\"type\":\"uint256\"}],\"name\":\"SetMaxNumParticipants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maximumNumSequencers\",\"type\":\"uint256\"}],\"name\":\"SetMaxNumSequencers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maximumNumValidators\",\"type\":\"uint256\"}],\"name\":\"SetMaxNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maximumNumWatchtowers\",\"type\":\"uint256\"}],\"name\":\"SetMaxNumWatchtowers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumNumParticipants\",\"type\":\"uint256\"}],\"name\":\"SetMinNumParticipants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumNumSequencers\",\"type\":\"uint256\"}],\"name\":\"SetMinNumSequencers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumNumValidators\",\"type\":\"uint256\"}],\"name\":\"SetMinNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimumNumWatchtowers\",\"type\":\"uint256\"}],\"name\":\"SetMinNumWatchtowers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPercentage\",\"type\":\"uint256\"}],\"name\":\"SetSlashPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"SetStakingMinThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressDisputedSequencerToWatchtower\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressDisputedSequencerToWatchtowerExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressDisputedWatchtowerToSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressDisputedWatchtowerToSequencerExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToDisputedWatchtowerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToIsDisputedWatchtower\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToIsParticipant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToIsSequencer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToIsSequencerInProbationAddr\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToIsValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToIsWatchtower\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToNodeType\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToParticipantIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToSequencerInProbationIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToSequencerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToSequencerToDisputedWatchtower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToStakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToValidatorIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_addressToWatchtowerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_dispute_watchtowers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_maximumNumParticipants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_maximumNumSequencers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_maximumNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_maximumNumWatchtowers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minStakingThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minimumNumParticipants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minimumNumSequencers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minimumNumValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_minimumNumWatchtowers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_participants\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_sequencers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_sequencers_in_probation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_slashPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_stakedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_watchtowers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"slashAddr\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"nodeType\",\"type\":\"string\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260405180606001604052806040518060400160405280600981526020017f73657175656e636572000000000000000000000000000000000000000000000081525081526020016040518060400160405280600a81526020017f7761746368746f7765720000000000000000000000000000000000000000000081525081526020016040518060400160405280600981526020017f76616c696461746f7200000000000000000000000000000000000000000000008152508152506000906003620000d092919062000160565b50348015620000de57600080fd5b50604051620063f8380380620063f883398181016040528101906200010491906200028d565b808211156200014a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001419062000381565b60405180910390fd5b8160098190555080600a819055505050620006fa565b828054828255906000526020600020908101928215620001ad579160200282015b82811115620001ac5782518290816200019b919062000613565b509160200191906001019062000181565b5b509050620001bc9190620001c0565b5090565b5b80821115620001e45760008181620001da9190620001e8565b50600101620001c1565b5090565b508054620001f6906200040c565b6000825580601f106200020a57506200022b565b601f0160209004906000526020600020908101906200022a91906200022e565b5b50565b5b80821115620002495760008160009055506001016200022f565b5090565b600080fd5b6000819050919050565b620002678162000252565b81146200027357600080fd5b50565b60008151905062000287816200025c565b92915050565b60008060408385031215620002a757620002a66200024d565b5b6000620002b78582860162000276565b9250506020620002ca8582860162000276565b9150509250929050565b600082825260208201905092915050565b7f4d696e207061727469636970616e7473206e756d6265722063616e206e6f742060008201527f62652067726561746572207468616e206d6178206e756d206f6620706172746960208201527f636970616e747300000000000000000000000000000000000000000000000000604082015250565b600062000369604783620002d4565b91506200037682620002e5565b606082019050919050565b600060208201905081810360008301526200039c816200035a565b9050919050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200042557607f821691505b6020821081036200043b576200043a620003dd565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620004a57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000466565b620004b1868362000466565b95508019841693508086168417925050509392505050565b6000819050919050565b6000620004f4620004ee620004e88462000252565b620004c9565b62000252565b9050919050565b6000819050919050565b6200051083620004d3565b620005286200051f82620004fb565b84845462000473565b825550505050565b600090565b6200053f62000530565b6200054c81848462000505565b505050565b5b8181101562000574576200056860008262000535565b60018101905062000552565b5050565b601f821115620005c3576200058d8162000441565b620005988462000456565b81016020851015620005a8578190505b620005c0620005b78562000456565b83018262000551565b50505b505050565b600082821c905092915050565b6000620005e860001984600802620005c8565b1980831691505092915050565b6000620006038383620005d5565b9150826002028217905092915050565b6200061e82620003a3565b67ffffffffffffffff8111156200063a5762000639620003ae565b5b6200064682546200040c565b6200065382828562000578565b600060209050601f8311600181146200068b576000841562000676578287015190505b620006828582620005f5565b865550620006f2565b601f1984166200069b8662000441565b60005b82811015620006c5578489015182556001820191506020850194506020810190506200069e565b86831015620006e55784890151620006e1601f891682620005d5565b8355505b6001600288020188555050505b505050505050565b615cee806200070a6000396000f3fe6080604052600436106104895760003560e01c8063a145012611610255578063c70966fb11610144578063dea79774116100c1578063eec9f52e11610085578063eec9f52e146113bf578063f196f824146113ea578063f76a1ee014611427578063f90ecacc14611452578063f9a43d4c1461148f578063fa70f5e5146114cc57610489565b8063dea79774146112b4578063e028aadf146112f1578063e2b5769d1461132e578063e387a7ed1461136b578063ebe4ad6f1461139657610489565b8063d147c38711610108578063d147c387146111a7578063d3be81a8146111e4578063d551832f1461120f578063d79f10d21461123a578063d80ff2b91461127757610489565b8063c70966fb146110c0578063c795c077146110fd578063c8fa5fe914611128578063c96be4cb14611153578063d14055171461117c57610489565b8063b8889f02116101d2578063c2ccacdf11610196578063c2ccacdf14610fd7578063c36352e114611002578063c402d43c1461102d578063c48c5f3714611058578063c6199c1f1461109557610489565b8063b8889f0214610eca578063bb5a26ff14610f07578063bd9ed0b314610f44578063be31929a14610f6f578063bfc5aeee14610fac57610489565b8063b17ca89c11610219578063b17ca89c14610dbd578063b1da3b1314610de8578063b41327ff14610e25578063b61b127f14610e62578063b737cd9e14610e9f57610489565b8063a145012614610cc2578063a17e8b3d14610ced578063aa90aa6614610d2a578063af6da36e14610d67578063aff1ea7614610d9257610489565b806346f45b8d1161037c5780637b339263116102f957806398abb25e116102bd57806398abb25e14610b8a57806398ff822d14610bc757806399a7c7e414610c045780639a083e8214610c2f5780639f9f710114610c5a578063a0f14bff14610c8557610489565b80637b33926314610a8f5780637dceceb814610acc578063855598b714610b095780638d08414814610b3457806390e7c86214610b5f57610489565b8063584060bb11610340578063584060bb146109965780636b2f7635146109d35780636edf2481146109fc57806372c8af2d14610a27578063760118b614610a6457610489565b806346f45b8d146108aa57806351184694146108c6578063560854f6146108f157806357a87ffd1461092e5780635804e51c1461095957610489565b80631f2743b61161040a5780632def6620116103ce5780632def6620146107d75780633603da84146107ee578063409b9ece1461081957806340b007531461084457806342ca58d51461088157610489565b80631f2743b6146106cc57806324208990146106f757806325d9b8c0146107345780632b3cfe0e146107715780632c404e9c146107ae57610489565b806319908df81161045157806319908df81461059b57806319ae6764146105d85780631b27dbc3146106155780631cb9bdca146106525780631efcf4d11461068f57610489565b806302b751991461048e578063065ae171146104cb578063068946e8146105085780630e533c8e146105335780630fb18ccc14610570575b600080fd5b34801561049a57600080fd5b506104b560048036038101906104b09190614ad2565b611509565b6040516104c29190614b18565b60405180910390f35b3480156104d757600080fd5b506104f260048036038101906104ed9190614ad2565b611521565b6040516104ff9190614b4e565b60405180910390f35b34801561051457600080fd5b5061051d611541565b60405161052a9190614b18565b60405180910390f35b34801561053f57600080fd5b5061055a60048036038101906105559190614ad2565b611547565b6040516105679190614b18565b60405180910390f35b34801561057c57600080fd5b5061058561155f565b6040516105929190614b18565b60405180910390f35b3480156105a757600080fd5b506105c260048036038101906105bd9190614b95565b611565565b6040516105cf9190614bd1565b60405180910390f35b3480156105e457600080fd5b506105ff60048036038101906105fa9190614ad2565b6115a4565b60405161060c9190614c85565b60405180910390f35b34801561062157600080fd5b5061063c60048036038101906106379190614b95565b611644565b6040516106499190614bd1565b60405180910390f35b34801561065e57600080fd5b5061067960048036038101906106749190614ad2565b611683565b6040516106869190614b18565b60405180910390f35b34801561069b57600080fd5b506106b660048036038101906106b19190614ad2565b61169b565b6040516106c39190614b4e565b60405180910390f35b3480156106d857600080fd5b506106e16116bb565b6040516106ee9190614db3565b60405180910390f35b34801561070357600080fd5b5061071e60048036038101906107199190614ad2565b611794565b60405161072b9190614bd1565b60405180910390f35b34801561074057600080fd5b5061075b60048036038101906107569190614ad2565b6117fd565b6040516107689190614b18565b60405180910390f35b34801561077d57600080fd5b5061079860048036038101906107939190614ad2565b611815565b6040516107a59190614b18565b60405180910390f35b3480156107ba57600080fd5b506107d560048036038101906107d09190614b95565b61182d565b005b3480156107e357600080fd5b506107ec611896565b005b3480156107fa57600080fd5b50610803611981565b6040516108109190614e93565b60405180910390f35b34801561082557600080fd5b5061082e611a0f565b60405161083b9190614b18565b60405180910390f35b34801561085057600080fd5b5061086b60048036038101906108669190614ad2565b611a1e565b6040516108789190614b4e565b60405180910390f35b34801561088d57600080fd5b506108a860048036038101906108a39190614b95565b611a3e565b005b6108c460048036038101906108bf9190614fea565b611aa7565b005b3480156108d257600080fd5b506108db611b12565b6040516108e89190614b18565b60405180910390f35b3480156108fd57600080fd5b5061091860048036038101906109139190614b95565b611b17565b6040516109259190614b18565b60405180910390f35b34801561093a57600080fd5b50610943611b2a565b6040516109509190614b18565b60405180910390f35b34801561096557600080fd5b50610980600480360381019061097b9190614ad2565b611b39565b60405161098d9190614bd1565b60405180910390f35b3480156109a257600080fd5b506109bd60048036038101906109b89190614b95565b611b6c565b6040516109ca9190614bd1565b60405180910390f35b3480156109df57600080fd5b506109fa60048036038101906109f59190614ad2565b611bab565b005b348015610a0857600080fd5b50610a11611c0d565b604051610a1e9190614b18565b60405180910390f35b348015610a3357600080fd5b50610a4e6004803603810190610a499190614ad2565b611c17565b604051610a5b9190614b4e565b60405180910390f35b348015610a7057600080fd5b50610a79611c37565b604051610a869190614b18565b60405180910390f35b348015610a9b57600080fd5b50610ab66004803603810190610ab19190614ad2565b611c41565b604051610ac39190614b4e565b60405180910390f35b348015610ad857600080fd5b50610af36004803603810190610aee9190614ad2565b611c61565b604051610b009190614b18565b60405180910390f35b348015610b1557600080fd5b50610b1e611c79565b604051610b2b9190614c85565b60405180910390f35b348015610b4057600080fd5b50610b49611cb2565b604051610b569190614b18565b60405180910390f35b348015610b6b57600080fd5b50610b74611cb8565b604051610b819190614b18565b60405180910390f35b348015610b9657600080fd5b50610bb16004803603810190610bac9190614b95565b611cbe565b604051610bbe9190614b18565b60405180910390f35b348015610bd357600080fd5b50610bee6004803603810190610be99190614ad2565b611cd1565b604051610bfb9190614b4e565b60405180910390f35b348015610c1057600080fd5b50610c19611d27565b604051610c269190614e93565b60405180910390f35b348015610c3b57600080fd5b50610c44611db5565b604051610c519190614c85565b60405180910390f35b348015610c6657600080fd5b50610c6f611dee565b604051610c7c9190614c85565b60405180910390f35b348015610c9157600080fd5b50610cac6004803603810190610ca79190614ad2565b611e27565b604051610cb99190614b18565b60405180910390f35b348015610cce57600080fd5b50610cd7611e3f565b604051610ce49190614b18565b60405180910390f35b348015610cf957600080fd5b50610d146004803603810190610d0f9190614ad2565b611e45565b604051610d219190614b4e565b60405180910390f35b348015610d3657600080fd5b50610d516004803603810190610d4c9190614ad2565b611e57565b604051610d5e9190614bd1565b60405180910390f35b348015610d7357600080fd5b50610d7c611ec0565b604051610d899190614b18565b60405180910390f35b348015610d9e57600080fd5b50610da7611ec6565b604051610db49190614b18565b60405180910390f35b348015610dc957600080fd5b50610dd2611ed2565b604051610ddf9190614b18565b60405180910390f35b348015610df457600080fd5b50610e0f6004803603810190610e0a9190614ad2565b611edc565b604051610e1c9190614bd1565b60405180910390f35b348015610e3157600080fd5b50610e4c6004803603810190610e479190614b95565b611f0e565b604051610e599190614c85565b60405180910390f35b348015610e6e57600080fd5b50610e896004803603810190610e849190614b95565b611fba565b604051610e969190614b18565b60405180910390f35b348015610eab57600080fd5b50610eb4611fcd565b604051610ec19190614e93565b60405180910390f35b348015610ed657600080fd5b50610ef16004803603810190610eec9190614ad2565b61205b565b604051610efe9190614b18565b60405180910390f35b348015610f1357600080fd5b50610f2e6004803603810190610f299190614ad2565b612073565b604051610f3b9190614b4e565b60405180910390f35b348015610f5057600080fd5b50610f596120c9565b604051610f669190614b18565b60405180910390f35b348015610f7b57600080fd5b50610f966004803603810190610f919190614ad2565b6120d3565b604051610fa39190614b4e565b60405180910390f35b348015610fb857600080fd5b50610fc16120f3565b604051610fce9190614b18565b60405180910390f35b348015610fe357600080fd5b50610fec6120fd565b604051610ff99190614b18565b60405180910390f35b34801561100e57600080fd5b50611017612107565b6040516110249190614e93565b60405180910390f35b34801561103957600080fd5b50611042612195565b60405161104f9190614e93565b60405180910390f35b34801561106457600080fd5b5061107f600480360381019061107a9190614b95565b612223565b60405161108c9190614bd1565b60405180910390f35b3480156110a157600080fd5b506110aa612262565b6040516110b79190614b18565b60405180910390f35b3480156110cc57600080fd5b506110e760048036038101906110e29190614ad2565b612268565b6040516110f49190614b4e565b60405180910390f35b34801561110957600080fd5b50611112612288565b60405161111f9190614b18565b60405180910390f35b34801561113457600080fd5b5061113d61228e565b60405161114a9190614b18565b60405180910390f35b34801561115f57600080fd5b5061117a60048036038101906111759190614ad2565b612298565b005b34801561118857600080fd5b506111916123d4565b60405161119e9190614e93565b60405180910390f35b3480156111b357600080fd5b506111ce60048036038101906111c99190614b95565b612462565b6040516111db9190614b18565b60405180910390f35b3480156111f057600080fd5b506111f9612475565b6040516112069190614b18565b60405180910390f35b34801561121b57600080fd5b5061122461247f565b6040516112319190614b18565b60405180910390f35b34801561124657600080fd5b50611261600480360381019061125c9190614ad2565b612485565b60405161126e9190614b18565b60405180910390f35b34801561128357600080fd5b5061129e60048036038101906112999190614b95565b6124ce565b6040516112ab9190614b18565b60405180910390f35b3480156112c057600080fd5b506112db60048036038101906112d69190614b95565b6124e1565b6040516112e89190614b18565b60405180910390f35b3480156112fd57600080fd5b5061131860048036038101906113139190614b95565b6124f4565b6040516113259190614b18565b60405180910390f35b34801561133a57600080fd5b5061135560048036038101906113509190614ad2565b612507565b6040516113629190614b4e565b60405180910390f35b34801561137757600080fd5b5061138061255d565b60405161138d9190614b18565b60405180910390f35b3480156113a257600080fd5b506113bd60048036038101906113b89190614ad2565b612563565b005b3480156113cb57600080fd5b506113d46125b2565b6040516113e19190614b18565b60405180910390f35b3480156113f657600080fd5b50611411600480360381019061140c9190614b95565b6125bc565b60405161141e9190614bd1565b60405180910390f35b34801561143357600080fd5b5061143c6125fb565b6040516114499190614b18565b60405180910390f35b34801561145e57600080fd5b5061147960048036038101906114749190614b95565b612601565b6040516114869190614bd1565b60405180910390f35b34801561149b57600080fd5b506114b660048036038101906114b19190614b95565b612640565b6040516114c39190614b18565b60405180910390f35b3480156114d857600080fd5b506114f360048036038101906114ee9190614ad2565b612653565b6040516115009190614b4e565b60405180910390f35b60246020528060005260406000206000915090505481565b60236020528060005260406000206000915054906101000a900460ff1681565b600a5481565b60166020528060005260406000206000915090505481565b600e5481565b6011818154811061157557600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600760205280600052604060002060009150905080546115c390615062565b80601f01602080910402602001604051908101604052809291908181526020018280546115ef90615062565b801561163c5780601f106116115761010080835404028352916020019161163c565b820191906000526020600020905b81548152906001019060200180831161161f57829003601f168201915b505050505081565b601a818154811061165457600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60136020528060005260406000206000915090505481565b60156020528060005260406000206000915054906101000a900460ff1681565b60606000805480602002602001604051908101604052809291908181526020016000905b8282101561178b5783829060005260206000200180546116fe90615062565b80601f016020809104026020016040519081016040528092919081815260200182805461172a90615062565b80156117775780601f1061174c57610100808354040283529160200191611777565b820191906000526020600020905b81548152906001019060200180831161175a57829003601f168201915b5050505050815260200190600101906116df565b50505050905090565b6000602060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b601c6020528060005260406000206000915090505481565b601d6020528060005260406000206000915090505481565b61184c3373ffffffffffffffffffffffffffffffffffffffff16612673565b1561188c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611883906150df565b60405180910390fd5b8060018190555050565b6118b53373ffffffffffffffffffffffffffffffffffffffff16612673565b156118f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118ec906150df565b60405180910390fd5b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205411611977576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161196e9061514b565b60405180910390fd5b61197f612696565b565b60606017805480602002602001604051908101604052809291908181526020018280548015611a0557602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116119bb575b5050505050905090565b6000611a19612869565b905090565b60186020528060005260406000206000915054906101000a900460ff1681565b611a5d3373ffffffffffffffffffffffffffffffffffffffff16612673565b15611a9d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a94906150df565b60405180910390fd5b8060028190555050565b611ac63373ffffffffffffffffffffffffffffffffffffffff16612673565b15611b06576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611afd906150df565b60405180910390fd5b611b0f8161288e565b50565b600181565b600081600c81905550600c549050919050565b6000611b34612b36565b905090565b601e6020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60148181548110611b7c57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b611bb481612b53565b611bbd33612c5a565b611bc73382612d61565b8073ffffffffffffffffffffffffffffffffffffffff167f14ebe573ff6e00d7293f5d831eb94cade703155cd0910eb5316a7b2f17b14c3560405160405180910390a250565b6000600f54905090565b60216020528060005260406000206000915054906101000a900460ff1681565b6000600b54905090565b60126020528060005260406000206000915054906101000a900460ff1681565b60056020528060005260406000206000915090505481565b6040518060400160405280600981526020017f73657175656e636572000000000000000000000000000000000000000000000081525081565b600b5481565b60105481565b6000816009819055506009549050919050565b6000602360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60606011805480602002602001604051908101604052809291908181526020018280548015611dab57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611d61575b5050505050905090565b6040518060400160405280600981526020017f76616c696461746f72000000000000000000000000000000000000000000000081525081565b6040518060400160405280600a81526020017f7761746368746f7765720000000000000000000000000000000000000000000081525081565b60196020528060005260406000206000915090505481565b60015481565b6000611e5082612f11565b9050919050565b6000601e60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b600f5481565b670de0b6b3a764000081565b6000601054905090565b602080528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008181548110611f1e57600080fd5b906000526020600020016000915090508054611f3990615062565b80601f0160208091040260200160405190810160405280929190818152602001828054611f6590615062565b8015611fb25780601f10611f8757610100808354040283529160200191611fb2565b820191906000526020600020905b815481529060010190602001808311611f9557829003601f168201915b505050505081565b600081600b81905550600b549050919050565b6060601a80548060200260200160405190810160405280929190818152602001828054801561205157602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612007575b5050505050905090565b60066020528060005260406000206000915090505481565b6000601260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6000600d54905090565b601f6020528060005260406000206000915054906101000a900460ff1681565b6000600954905090565b6000600e54905090565b6060602280548060200260200160405190810160405280929190818152602001828054801561218b57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612141575b5050505050905090565b6060601480548060200260200160405190810160405280929190818152602001828054801561221957602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116121cf575b5050505050905090565b6017818154811061223357600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60095481565b60046020528060005260406000206000915054906101000a900460ff1681565b600c5481565b6000600a54905090565b6122b73373ffffffffffffffffffffffffffffffffffffffff16612673565b156122f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122ee906150df565b60405180910390fd5b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205411612379576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123709061514b565b60405180910390fd5b6001151561238633612f67565b1515146123c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123bf906151b7565b60405180910390fd5b6123d181612fbd565b50565b6060600380548060200260200160405190810160405280929190818152602001828054801561245857602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161240e575b5050505050905090565b6000816010819055506010549050919050565b6000600854905090565b60025481565b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600081600d81905550600d549050919050565b600081600f81905550600f549050919050565b600081600a81905550600a549050919050565b6000601860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60085481565b61256c8161331c565b8073ffffffffffffffffffffffffffffffffffffffff167f7da9849dcad1bb34324464458a33a9d0820286dedf66ec9c6669e730118d1c0660405160405180910390a250565b6000600c54905090565b600381815481106125cc57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600d5481565b6022818154811061261157600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600081600e81905550600e549050919050565b601b6020528060005260406000206000915054906101000a900460ff1681565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b61269f3361361a565b6126de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016126d59061526f565b60405180910390fd5b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550806008600082825461277991906152be565b9250508190555061278933613670565b61279233612f67565b156127a1576127a03361396f565b5b6127aa33613c25565b156127b9576127b833613c7b565b5b6127c233613f31565b156127d1576127d033613f87565b5b3373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015612817573d6000803e3d6000fd5b503373ffffffffffffffffffffffffffffffffffffffff167f0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f758260405161285e9190614b18565b60405180910390a250565b600080600154111561287f57600154905061288b565b670de0b6b3a764000090505b90565b6128978161423d565b806128a757506128a6816142da565b5b806128b757506128b681614377565b5b6128f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128ed90615364565b60405180910390fd5b6128fe612869565b341015612940576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129379061541c565b60405180910390fd5b61294a3382614414565b612989576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612980906154ae565b60405180910390fd5b612993338261448d565b6129d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016129c990615540565b60405180910390fd5b6129dc3382614506565b612a1b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a12906155d2565b60405180910390fd5b3460086000828254612a2d91906155f2565b9250508190555034600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254612a8391906155f2565b92505081905550612a933361457f565b612a9c8161423d565b15612aaf57612aaa336146d8565b612ae5565b612ab8816142da565b15612acb57612ac6336147df565b612ae4565b612ad481614377565b15612ae357612ae2336148e6565b5b5b5b3373ffffffffffffffffffffffffffffffffffffffff167f9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d34604051612b2b9190614b18565b60405180910390a250565b60008060025411612b4a5760019050612b50565b60025490505b90565b6001601560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550601480549050601660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506014819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6001601b60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550601a80549050601c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550601a819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b80601e60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001601f60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555081602060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001602160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b6000601560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6000601260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000613020826064613011612b36565b61301b9190615648565b6149ed565b9050600060011515602160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515036130e357601e60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905061319f565b60011515602160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615150361319e57602060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b5b600082846131ad91906152be565b905080600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550826008600082825461320591906152be565b925050819055508173ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f19350505050158015613252573d6000803e3d6000fd5b508173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f10cc080888f97c96419786385100283aa8ae09efd09b75964b955212f42847d483866040516132b29291906156a2565b60405180910390a36132c385612f11565b15613315576132d18561331c565b8473ffffffffffffffffffffffffffffffffffffffff167f7da9849dcad1bb34324464458a33a9d0820286dedf66ec9c6669e730118d1c0660405160405180910390a25b5050505050565b61332581612f11565b613364576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161335b90615763565b60405180910390fd5b601480549050601660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054106133ea576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133e1906157f5565b60405180910390fd5b6000601660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000600160148054905061344291906152be565b90508082146135315760006014828154811061346157613460615815565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080601484815481106134a3576134a2615815565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082601660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505b6000601560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000601660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060148054806135e0576135df615844565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055505050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b600954600380549050116136b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136b09061590b565b60405180910390fd5b600380549050600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541061373f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137369061599d565b60405180910390fd5b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000600160038054905061379791906152be565b9050808214613886576000600382815481106137b6576137b5615815565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080600384815481106137f8576137f7615815565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505b6000600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600380548061393557613934615844565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055505050565b601180549050601360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054106139f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016139ec90615a2f565b60405180910390fd5b6000601360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060006001601180549050613a4d91906152be565b9050808214613b3c57600060118281548110613a6c57613a6b615815565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508060118481548110613aae57613aad615815565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082601360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505b6000601260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000601360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506011805480613beb57613bea615844565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055505050565b6000601860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b601780549050601960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410613d01576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613cf890615ac1565b60405180910390fd5b6000601960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060006001601780549050613d5991906152be565b9050808214613e4857600060178281548110613d7857613d77615815565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508060178481548110613dba57613db9615815565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082601960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505b6000601860008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000601960008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506017805480613ef757613ef6615844565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055505050565b6000602360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b602280549050602460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541061400d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161400490615b53565b60405180910390fd5b6000602460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000600160228054905061406591906152be565b90508082146141545760006022828154811061408457614083615815565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080602284815481106140c6576140c5615815565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082602460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505b6000602360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000602460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550602280548061420357614202615844565b5b6001900381819060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690559055505050565b6000816040516020016142509190615baf565b604051602081830303815290604052805190602001206040518060400160405280600981526020017f73657175656e63657200000000000000000000000000000000000000000000008152506040516020016142ac9190615baf565b60405160208183030381529060405280519060200120036142d057600190506142d5565b600090505b919050565b6000816040516020016142ed9190615baf565b604051602081830303815290604052805190602001206040518060400160405280600a81526020017f7761746368746f776572000000000000000000000000000000000000000000008152506040516020016143499190615baf565b604051602081830303815290604052805190602001200361436d5760019050614372565b600090505b919050565b60008160405160200161438a9190615baf565b604051602081830303815290604052805190602001206040518060400160405280600981526020017f76616c696461746f7200000000000000000000000000000000000000000000008152506040516020016143e69190615baf565b604051602081830303815290604052805190602001200361440a576001905061440f565b600090505b919050565b600061441f8261423d565b80156144745750601260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff165b156144825760009050614487565b600190505b92915050565b6000614498826142da565b80156144ed5750601860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff165b156144fb5760009050614500565b600190505b92915050565b600061451182614377565b80156145665750602360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff165b156145745760009050614579565b600190505b92915050565b600460008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166146d5576001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600380549050600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506003819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b6001601260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550601180549050601360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506011819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6001601860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550601780549050601960008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506017819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6001602360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550602280549050602460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506022819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600061271082846149fe9190615648565b1015614a3f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614a3690615c38565b60405180910390fd5b6127108284614a4e9190615648565b614a589190615c87565b905092915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000614a9f82614a74565b9050919050565b614aaf81614a94565b8114614aba57600080fd5b50565b600081359050614acc81614aa6565b92915050565b600060208284031215614ae857614ae7614a6a565b5b6000614af684828501614abd565b91505092915050565b6000819050919050565b614b1281614aff565b82525050565b6000602082019050614b2d6000830184614b09565b92915050565b60008115159050919050565b614b4881614b33565b82525050565b6000602082019050614b636000830184614b3f565b92915050565b614b7281614aff565b8114614b7d57600080fd5b50565b600081359050614b8f81614b69565b92915050565b600060208284031215614bab57614baa614a6a565b5b6000614bb984828501614b80565b91505092915050565b614bcb81614a94565b82525050565b6000602082019050614be66000830184614bc2565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015614c26578082015181840152602081019050614c0b565b83811115614c35576000848401525b50505050565b6000601f19601f8301169050919050565b6000614c5782614bec565b614c618185614bf7565b9350614c71818560208601614c08565b614c7a81614c3b565b840191505092915050565b60006020820190508181036000830152614c9f8184614c4c565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b6000614cef82614bec565b614cf98185614cd3565b9350614d09818560208601614c08565b614d1281614c3b565b840191505092915050565b6000614d298383614ce4565b905092915050565b6000602082019050919050565b6000614d4982614ca7565b614d538185614cb2565b935083602082028501614d6585614cc3565b8060005b85811015614da15784840389528151614d828582614d1d565b9450614d8d83614d31565b925060208a01995050600181019050614d69565b50829750879550505050505092915050565b60006020820190508181036000830152614dcd8184614d3e565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b614e0a81614a94565b82525050565b6000614e1c8383614e01565b60208301905092915050565b6000602082019050919050565b6000614e4082614dd5565b614e4a8185614de0565b9350614e5583614df1565b8060005b83811015614e86578151614e6d8882614e10565b9750614e7883614e28565b925050600181019050614e59565b5085935050505092915050565b60006020820190508181036000830152614ead8184614e35565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b614ef782614c3b565b810181811067ffffffffffffffff82111715614f1657614f15614ebf565b5b80604052505050565b6000614f29614a60565b9050614f358282614eee565b919050565b600067ffffffffffffffff821115614f5557614f54614ebf565b5b614f5e82614c3b565b9050602081019050919050565b82818337600083830152505050565b6000614f8d614f8884614f3a565b614f1f565b905082815260208101848484011115614fa957614fa8614eba565b5b614fb4848285614f6b565b509392505050565b600082601f830112614fd157614fd0614eb5565b5b8135614fe1848260208601614f7a565b91505092915050565b60006020828403121561500057614fff614a6a565b5b600082013567ffffffffffffffff81111561501e5761501d614a6f565b5b61502a84828501614fbc565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061507a57607f821691505b60208210810361508d5761508c615033565b5b50919050565b7f4f6e6c7920454f412063616e2063616c6c2066756e6374696f6e000000000000600082015250565b60006150c9601a83614bf7565b91506150d482615093565b602082019050919050565b600060208201905081810360008301526150f8816150bc565b9050919050565b7f4f6e6c79207374616b65722063616e2063616c6c2066756e6374696f6e000000600082015250565b6000615135601d83614bf7565b9150615140826150ff565b602082019050919050565b6000602082019050818103600083015261516481615128565b9050919050565b7f4f6e6c792073657175656e6365722063616e2063616c6c2066756e6374696f6e600082015250565b60006151a1602083614bf7565b91506151ac8261516b565b602082019050919050565b600060208201905081810360008301526151d081615194565b9050919050565b7f53656e6465722068617320746f2062652070617274206f66207374616b696e6760008201527f20706f6c6c20696e206f7264657220746f20756e7374616b652069747320736860208201527f6172652e20556e7374616b652072656a65637465642e00000000000000000000604082015250565b6000615259605683614bf7565b9150615264826151d7565b606082019050919050565b600060208201905081810360008301526152888161524c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006152c982614aff565b91506152d483614aff565b9250828210156152e7576152e661528f565b5b828203905092915050565b7f50726f7669646564206e6f646520747970652068617320746f206d617463682060008201527f617661696c61626c65206e6f6465207479706573000000000000000000000000602082015250565b600061534e603483614bf7565b9150615359826152f2565b604082019050919050565b6000602082019050818103600083015261537d81615341565b9050919050565b7f496e737566696369656e74207374616b696e6720616d6f756e742070726f766960008201527f6465642e2048617320746f206265206c6172676572207468616e207374616b6960208201527f6e67206d696e696d756d207468726573686f6c64000000000000000000000000604082015250565b6000615406605483614bf7565b915061541182615384565b606082019050919050565b60006020820190508181036000830152615435816153f9565b9050919050565b7f53656e64657220697320616c72656164792073657175656e6365722e2052656a60008201527f656374696e67207374616b6520726571756573742e0000000000000000000000602082015250565b6000615498603583614bf7565b91506154a38261543c565b604082019050919050565b600060208201905081810360008301526154c78161548b565b9050919050565b7f53656e64657220697320616c7265616479207761746368746f7765722e20526560008201527f6a656374696e67207374616b6520726571756573742e00000000000000000000602082015250565b600061552a603683614bf7565b9150615535826154ce565b604082019050919050565b600060208201905081810360008301526155598161551d565b9050919050565b7f53656e64657220697320616c72656164792076616c696461746f722e2052656a60008201527f656374696e67207374616b6520726571756573742e0000000000000000000000602082015250565b60006155bc603583614bf7565b91506155c782615560565b604082019050919050565b600060208201905081810360008301526155eb816155af565b9050919050565b60006155fd82614aff565b915061560883614aff565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561563d5761563c61528f565b5b828201905092915050565b600061565382614aff565b915061565e83614aff565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156156975761569661528f565b5b828202905092915050565b60006040820190506156b76000830185614b09565b6156c46020830184614b09565b9392505050565b7f416464726573732068617320746f20626520696e2070726f626174696f6e206960008201527f6e206f7264657220746f2064656c6574652069742066726f6d2074686520707260208201527f6f626174696f6e2073657175656e63657273206c6973742e0000000000000000604082015250565b600061574d605883614bf7565b9150615758826156cb565b606082019050919050565b6000602082019050818103600083015261577c81615740565b9050919050565b7f6d616c6963696f7573207061727469636970616e7420696e646578206f75742060008201527f6f662072616e676520696e206d617070696e6700000000000000000000000000602082015250565b60006157df603383614bf7565b91506157ea82615783565b604082019050919050565b6000602082019050818103600083015261580e816157d2565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f5374616b696e67207061727469636970616e74732063616e2774206265206c6560008201527f7373207468616e20746865206d696e696d756d2072657175697265642070617260208201527f7469636970616e74206e756d0000000000000000000000000000000000000000604082015250565b60006158f5604c83614bf7565b915061590082615873565b606082019050919050565b60006020820190508181036000830152615924816158e8565b9050919050565b7f7061727469636970616e7420696e646578206f7574206f662072616e6765206960008201527f6e206d617070696e670000000000000000000000000000000000000000000000602082015250565b6000615987602983614bf7565b91506159928261592b565b604082019050919050565b600060208201905081810360008301526159b68161597a565b9050919050565b7f73657175656e63657220696e646578206f7574206f662072616e676520696e2060008201527f6d617070696e6700000000000000000000000000000000000000000000000000602082015250565b6000615a19602783614bf7565b9150615a24826159bd565b604082019050919050565b60006020820190508181036000830152615a4881615a0c565b9050919050565b7f7761746368746f77657220696e646578206f7574206f662072616e676520696e60008201527f206d617070696e67000000000000000000000000000000000000000000000000602082015250565b6000615aab602883614bf7565b9150615ab682615a4f565b604082019050919050565b60006020820190508181036000830152615ada81615a9e565b9050919050565b7f76616c696461746f7220696e646578206f7574206f662072616e676520696e2060008201527f6d617070696e6700000000000000000000000000000000000000000000000000602082015250565b6000615b3d602783614bf7565b9150615b4882615ae1565b604082019050919050565b60006020820190508181036000830152615b6c81615b30565b9050919050565b600081905092915050565b6000615b8982614bec565b615b938185615b73565b9350615ba3818560208601614c08565b80840191505092915050565b6000615bbb8284615b7e565b915081905092915050565b7f416d6f756e7420697320746f6f206c6f7720746f2063616c63756c617465206660008201527f6565000000000000000000000000000000000000000000000000000000000000602082015250565b6000615c22602283614bf7565b9150615c2d82615bc6565b604082019050919050565b60006020820190508181036000830152615c5181615c15565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000615c9282614aff565b9150615c9d83614aff565b925082615cad57615cac615c58565b5b82820490509291505056fea26469706673582212202364ec7796dc5de9de6e35ec35dccdcff5724d03f3202bd71c511101a0ca69a164736f6c634300080f0033",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// StakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingMetaData.Bin instead.
var StakingBin = StakingMetaData.Bin

// DeployStaking deploys a new Ethereum contract, binding an instance of Staking to it.
func DeployStaking(auth *bind.TransactOpts, backend bind.ContractBackend, minNumParticipants *big.Int, maxNumParticipants *big.Int) (common.Address, *types.Transaction, *Staking, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingBin), backend, minNumParticipants, maxNumParticipants)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// AVAILABLENODETYPES is a free data retrieval call binding the contract method 0xb41327ff.
//
// Solidity: function AVAILABLE_NODE_TYPES(uint256 ) view returns(string)
func (_Staking *StakingCaller) AVAILABLENODETYPES(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "AVAILABLE_NODE_TYPES", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AVAILABLENODETYPES is a free data retrieval call binding the contract method 0xb41327ff.
//
// Solidity: function AVAILABLE_NODE_TYPES(uint256 ) view returns(string)
func (_Staking *StakingSession) AVAILABLENODETYPES(arg0 *big.Int) (string, error) {
	return _Staking.Contract.AVAILABLENODETYPES(&_Staking.CallOpts, arg0)
}

// AVAILABLENODETYPES is a free data retrieval call binding the contract method 0xb41327ff.
//
// Solidity: function AVAILABLE_NODE_TYPES(uint256 ) view returns(string)
func (_Staking *StakingCallerSession) AVAILABLENODETYPES(arg0 *big.Int) (string, error) {
	return _Staking.Contract.AVAILABLENODETYPES(&_Staking.CallOpts, arg0)
}

// DEFAULTMINSLASHPERCENTAGE is a free data retrieval call binding the contract method 0x51184694.
//
// Solidity: function DEFAULT_MIN_SLASH_PERCENTAGE() view returns(uint256)
func (_Staking *StakingCaller) DEFAULTMINSLASHPERCENTAGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "DEFAULT_MIN_SLASH_PERCENTAGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTMINSLASHPERCENTAGE is a free data retrieval call binding the contract method 0x51184694.
//
// Solidity: function DEFAULT_MIN_SLASH_PERCENTAGE() view returns(uint256)
func (_Staking *StakingSession) DEFAULTMINSLASHPERCENTAGE() (*big.Int, error) {
	return _Staking.Contract.DEFAULTMINSLASHPERCENTAGE(&_Staking.CallOpts)
}

// DEFAULTMINSLASHPERCENTAGE is a free data retrieval call binding the contract method 0x51184694.
//
// Solidity: function DEFAULT_MIN_SLASH_PERCENTAGE() view returns(uint256)
func (_Staking *StakingCallerSession) DEFAULTMINSLASHPERCENTAGE() (*big.Int, error) {
	return _Staking.Contract.DEFAULTMINSLASHPERCENTAGE(&_Staking.CallOpts)
}

// DEFAULTSTAKINGTHRESHOLD is a free data retrieval call binding the contract method 0xaff1ea76.
//
// Solidity: function DEFAULT_STAKING_THRESHOLD() view returns(uint256)
func (_Staking *StakingCaller) DEFAULTSTAKINGTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "DEFAULT_STAKING_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTSTAKINGTHRESHOLD is a free data retrieval call binding the contract method 0xaff1ea76.
//
// Solidity: function DEFAULT_STAKING_THRESHOLD() view returns(uint256)
func (_Staking *StakingSession) DEFAULTSTAKINGTHRESHOLD() (*big.Int, error) {
	return _Staking.Contract.DEFAULTSTAKINGTHRESHOLD(&_Staking.CallOpts)
}

// DEFAULTSTAKINGTHRESHOLD is a free data retrieval call binding the contract method 0xaff1ea76.
//
// Solidity: function DEFAULT_STAKING_THRESHOLD() view returns(uint256)
func (_Staking *StakingCallerSession) DEFAULTSTAKINGTHRESHOLD() (*big.Int, error) {
	return _Staking.Contract.DEFAULTSTAKINGTHRESHOLD(&_Staking.CallOpts)
}

// GetAvailableNodeTypes is a free data retrieval call binding the contract method 0x1f2743b6.
//
// Solidity: function GetAvailableNodeTypes() view returns(string[])
func (_Staking *StakingCaller) GetAvailableNodeTypes(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetAvailableNodeTypes")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAvailableNodeTypes is a free data retrieval call binding the contract method 0x1f2743b6.
//
// Solidity: function GetAvailableNodeTypes() view returns(string[])
func (_Staking *StakingSession) GetAvailableNodeTypes() ([]string, error) {
	return _Staking.Contract.GetAvailableNodeTypes(&_Staking.CallOpts)
}

// GetAvailableNodeTypes is a free data retrieval call binding the contract method 0x1f2743b6.
//
// Solidity: function GetAvailableNodeTypes() view returns(string[])
func (_Staking *StakingCallerSession) GetAvailableNodeTypes() ([]string, error) {
	return _Staking.Contract.GetAvailableNodeTypes(&_Staking.CallOpts)
}

// GetCurrentAccountStakedAmount is a free data retrieval call binding the contract method 0xd79f10d2.
//
// Solidity: function GetCurrentAccountStakedAmount(address addr) view returns(uint256)
func (_Staking *StakingCaller) GetCurrentAccountStakedAmount(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetCurrentAccountStakedAmount", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentAccountStakedAmount is a free data retrieval call binding the contract method 0xd79f10d2.
//
// Solidity: function GetCurrentAccountStakedAmount(address addr) view returns(uint256)
func (_Staking *StakingSession) GetCurrentAccountStakedAmount(addr common.Address) (*big.Int, error) {
	return _Staking.Contract.GetCurrentAccountStakedAmount(&_Staking.CallOpts, addr)
}

// GetCurrentAccountStakedAmount is a free data retrieval call binding the contract method 0xd79f10d2.
//
// Solidity: function GetCurrentAccountStakedAmount(address addr) view returns(uint256)
func (_Staking *StakingCallerSession) GetCurrentAccountStakedAmount(addr common.Address) (*big.Int, error) {
	return _Staking.Contract.GetCurrentAccountStakedAmount(&_Staking.CallOpts, addr)
}

// GetCurrentDisputeWatchtowers is a free data retrieval call binding the contract method 0xb737cd9e.
//
// Solidity: function GetCurrentDisputeWatchtowers() view returns(address[])
func (_Staking *StakingCaller) GetCurrentDisputeWatchtowers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetCurrentDisputeWatchtowers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentDisputeWatchtowers is a free data retrieval call binding the contract method 0xb737cd9e.
//
// Solidity: function GetCurrentDisputeWatchtowers() view returns(address[])
func (_Staking *StakingSession) GetCurrentDisputeWatchtowers() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentDisputeWatchtowers(&_Staking.CallOpts)
}

// GetCurrentDisputeWatchtowers is a free data retrieval call binding the contract method 0xb737cd9e.
//
// Solidity: function GetCurrentDisputeWatchtowers() view returns(address[])
func (_Staking *StakingCallerSession) GetCurrentDisputeWatchtowers() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentDisputeWatchtowers(&_Staking.CallOpts)
}

// GetCurrentParticipants is a free data retrieval call binding the contract method 0xd1405517.
//
// Solidity: function GetCurrentParticipants() view returns(address[])
func (_Staking *StakingCaller) GetCurrentParticipants(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetCurrentParticipants")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentParticipants is a free data retrieval call binding the contract method 0xd1405517.
//
// Solidity: function GetCurrentParticipants() view returns(address[])
func (_Staking *StakingSession) GetCurrentParticipants() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentParticipants(&_Staking.CallOpts)
}

// GetCurrentParticipants is a free data retrieval call binding the contract method 0xd1405517.
//
// Solidity: function GetCurrentParticipants() view returns(address[])
func (_Staking *StakingCallerSession) GetCurrentParticipants() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentParticipants(&_Staking.CallOpts)
}

// GetCurrentSequencers is a free data retrieval call binding the contract method 0x99a7c7e4.
//
// Solidity: function GetCurrentSequencers() view returns(address[])
func (_Staking *StakingCaller) GetCurrentSequencers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetCurrentSequencers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentSequencers is a free data retrieval call binding the contract method 0x99a7c7e4.
//
// Solidity: function GetCurrentSequencers() view returns(address[])
func (_Staking *StakingSession) GetCurrentSequencers() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentSequencers(&_Staking.CallOpts)
}

// GetCurrentSequencers is a free data retrieval call binding the contract method 0x99a7c7e4.
//
// Solidity: function GetCurrentSequencers() view returns(address[])
func (_Staking *StakingCallerSession) GetCurrentSequencers() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentSequencers(&_Staking.CallOpts)
}

// GetCurrentSequencersInProbation is a free data retrieval call binding the contract method 0xc402d43c.
//
// Solidity: function GetCurrentSequencersInProbation() view returns(address[])
func (_Staking *StakingCaller) GetCurrentSequencersInProbation(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetCurrentSequencersInProbation")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentSequencersInProbation is a free data retrieval call binding the contract method 0xc402d43c.
//
// Solidity: function GetCurrentSequencersInProbation() view returns(address[])
func (_Staking *StakingSession) GetCurrentSequencersInProbation() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentSequencersInProbation(&_Staking.CallOpts)
}

// GetCurrentSequencersInProbation is a free data retrieval call binding the contract method 0xc402d43c.
//
// Solidity: function GetCurrentSequencersInProbation() view returns(address[])
func (_Staking *StakingCallerSession) GetCurrentSequencersInProbation() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentSequencersInProbation(&_Staking.CallOpts)
}

// GetCurrentStakedAmount is a free data retrieval call binding the contract method 0xd3be81a8.
//
// Solidity: function GetCurrentStakedAmount() view returns(uint256)
func (_Staking *StakingCaller) GetCurrentStakedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetCurrentStakedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentStakedAmount is a free data retrieval call binding the contract method 0xd3be81a8.
//
// Solidity: function GetCurrentStakedAmount() view returns(uint256)
func (_Staking *StakingSession) GetCurrentStakedAmount() (*big.Int, error) {
	return _Staking.Contract.GetCurrentStakedAmount(&_Staking.CallOpts)
}

// GetCurrentStakedAmount is a free data retrieval call binding the contract method 0xd3be81a8.
//
// Solidity: function GetCurrentStakedAmount() view returns(uint256)
func (_Staking *StakingCallerSession) GetCurrentStakedAmount() (*big.Int, error) {
	return _Staking.Contract.GetCurrentStakedAmount(&_Staking.CallOpts)
}

// GetCurrentStakingThreshold is a free data retrieval call binding the contract method 0x409b9ece.
//
// Solidity: function GetCurrentStakingThreshold() view returns(uint256)
func (_Staking *StakingCaller) GetCurrentStakingThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetCurrentStakingThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentStakingThreshold is a free data retrieval call binding the contract method 0x409b9ece.
//
// Solidity: function GetCurrentStakingThreshold() view returns(uint256)
func (_Staking *StakingSession) GetCurrentStakingThreshold() (*big.Int, error) {
	return _Staking.Contract.GetCurrentStakingThreshold(&_Staking.CallOpts)
}

// GetCurrentStakingThreshold is a free data retrieval call binding the contract method 0x409b9ece.
//
// Solidity: function GetCurrentStakingThreshold() view returns(uint256)
func (_Staking *StakingCallerSession) GetCurrentStakingThreshold() (*big.Int, error) {
	return _Staking.Contract.GetCurrentStakingThreshold(&_Staking.CallOpts)
}

// GetCurrentValidators is a free data retrieval call binding the contract method 0xc36352e1.
//
// Solidity: function GetCurrentValidators() view returns(address[])
func (_Staking *StakingCaller) GetCurrentValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetCurrentValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentValidators is a free data retrieval call binding the contract method 0xc36352e1.
//
// Solidity: function GetCurrentValidators() view returns(address[])
func (_Staking *StakingSession) GetCurrentValidators() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentValidators(&_Staking.CallOpts)
}

// GetCurrentValidators is a free data retrieval call binding the contract method 0xc36352e1.
//
// Solidity: function GetCurrentValidators() view returns(address[])
func (_Staking *StakingCallerSession) GetCurrentValidators() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentValidators(&_Staking.CallOpts)
}

// GetCurrentWatchtowers is a free data retrieval call binding the contract method 0x3603da84.
//
// Solidity: function GetCurrentWatchtowers() view returns(address[])
func (_Staking *StakingCaller) GetCurrentWatchtowers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetCurrentWatchtowers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentWatchtowers is a free data retrieval call binding the contract method 0x3603da84.
//
// Solidity: function GetCurrentWatchtowers() view returns(address[])
func (_Staking *StakingSession) GetCurrentWatchtowers() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentWatchtowers(&_Staking.CallOpts)
}

// GetCurrentWatchtowers is a free data retrieval call binding the contract method 0x3603da84.
//
// Solidity: function GetCurrentWatchtowers() view returns(address[])
func (_Staking *StakingCallerSession) GetCurrentWatchtowers() ([]common.Address, error) {
	return _Staking.Contract.GetCurrentWatchtowers(&_Staking.CallOpts)
}

// GetDisputedSequencerAddrs is a free data retrieval call binding the contract method 0xaa90aa66.
//
// Solidity: function GetDisputedSequencerAddrs(address watchtowerAddr) view returns(address)
func (_Staking *StakingCaller) GetDisputedSequencerAddrs(opts *bind.CallOpts, watchtowerAddr common.Address) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetDisputedSequencerAddrs", watchtowerAddr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDisputedSequencerAddrs is a free data retrieval call binding the contract method 0xaa90aa66.
//
// Solidity: function GetDisputedSequencerAddrs(address watchtowerAddr) view returns(address)
func (_Staking *StakingSession) GetDisputedSequencerAddrs(watchtowerAddr common.Address) (common.Address, error) {
	return _Staking.Contract.GetDisputedSequencerAddrs(&_Staking.CallOpts, watchtowerAddr)
}

// GetDisputedSequencerAddrs is a free data retrieval call binding the contract method 0xaa90aa66.
//
// Solidity: function GetDisputedSequencerAddrs(address watchtowerAddr) view returns(address)
func (_Staking *StakingCallerSession) GetDisputedSequencerAddrs(watchtowerAddr common.Address) (common.Address, error) {
	return _Staking.Contract.GetDisputedSequencerAddrs(&_Staking.CallOpts, watchtowerAddr)
}

// GetDisputedWatchtowerAddr is a free data retrieval call binding the contract method 0x24208990.
//
// Solidity: function GetDisputedWatchtowerAddr(address sequencerAddr) view returns(address)
func (_Staking *StakingCaller) GetDisputedWatchtowerAddr(opts *bind.CallOpts, sequencerAddr common.Address) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetDisputedWatchtowerAddr", sequencerAddr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDisputedWatchtowerAddr is a free data retrieval call binding the contract method 0x24208990.
//
// Solidity: function GetDisputedWatchtowerAddr(address sequencerAddr) view returns(address)
func (_Staking *StakingSession) GetDisputedWatchtowerAddr(sequencerAddr common.Address) (common.Address, error) {
	return _Staking.Contract.GetDisputedWatchtowerAddr(&_Staking.CallOpts, sequencerAddr)
}

// GetDisputedWatchtowerAddr is a free data retrieval call binding the contract method 0x24208990.
//
// Solidity: function GetDisputedWatchtowerAddr(address sequencerAddr) view returns(address)
func (_Staking *StakingCallerSession) GetDisputedWatchtowerAddr(sequencerAddr common.Address) (common.Address, error) {
	return _Staking.Contract.GetDisputedWatchtowerAddr(&_Staking.CallOpts, sequencerAddr)
}

// GetIsSequencerInProbation is a free data retrieval call binding the contract method 0xa17e8b3d.
//
// Solidity: function GetIsSequencerInProbation(address sequencerAddr) view returns(bool)
func (_Staking *StakingCaller) GetIsSequencerInProbation(opts *bind.CallOpts, sequencerAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetIsSequencerInProbation", sequencerAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsSequencerInProbation is a free data retrieval call binding the contract method 0xa17e8b3d.
//
// Solidity: function GetIsSequencerInProbation(address sequencerAddr) view returns(bool)
func (_Staking *StakingSession) GetIsSequencerInProbation(sequencerAddr common.Address) (bool, error) {
	return _Staking.Contract.GetIsSequencerInProbation(&_Staking.CallOpts, sequencerAddr)
}

// GetIsSequencerInProbation is a free data retrieval call binding the contract method 0xa17e8b3d.
//
// Solidity: function GetIsSequencerInProbation(address sequencerAddr) view returns(bool)
func (_Staking *StakingCallerSession) GetIsSequencerInProbation(sequencerAddr common.Address) (bool, error) {
	return _Staking.Contract.GetIsSequencerInProbation(&_Staking.CallOpts, sequencerAddr)
}

// GetMaxNumParticipants is a free data retrieval call binding the contract method 0xc8fa5fe9.
//
// Solidity: function GetMaxNumParticipants() view returns(uint256)
func (_Staking *StakingCaller) GetMaxNumParticipants(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetMaxNumParticipants")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxNumParticipants is a free data retrieval call binding the contract method 0xc8fa5fe9.
//
// Solidity: function GetMaxNumParticipants() view returns(uint256)
func (_Staking *StakingSession) GetMaxNumParticipants() (*big.Int, error) {
	return _Staking.Contract.GetMaxNumParticipants(&_Staking.CallOpts)
}

// GetMaxNumParticipants is a free data retrieval call binding the contract method 0xc8fa5fe9.
//
// Solidity: function GetMaxNumParticipants() view returns(uint256)
func (_Staking *StakingCallerSession) GetMaxNumParticipants() (*big.Int, error) {
	return _Staking.Contract.GetMaxNumParticipants(&_Staking.CallOpts)
}

// GetMaxNumSequencers is a free data retrieval call binding the contract method 0xc2ccacdf.
//
// Solidity: function GetMaxNumSequencers() view returns(uint256)
func (_Staking *StakingCaller) GetMaxNumSequencers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetMaxNumSequencers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxNumSequencers is a free data retrieval call binding the contract method 0xc2ccacdf.
//
// Solidity: function GetMaxNumSequencers() view returns(uint256)
func (_Staking *StakingSession) GetMaxNumSequencers() (*big.Int, error) {
	return _Staking.Contract.GetMaxNumSequencers(&_Staking.CallOpts)
}

// GetMaxNumSequencers is a free data retrieval call binding the contract method 0xc2ccacdf.
//
// Solidity: function GetMaxNumSequencers() view returns(uint256)
func (_Staking *StakingCallerSession) GetMaxNumSequencers() (*big.Int, error) {
	return _Staking.Contract.GetMaxNumSequencers(&_Staking.CallOpts)
}

// GetMaxNumValidators is a free data retrieval call binding the contract method 0x6edf2481.
//
// Solidity: function GetMaxNumValidators() view returns(uint256)
func (_Staking *StakingCaller) GetMaxNumValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetMaxNumValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxNumValidators is a free data retrieval call binding the contract method 0x6edf2481.
//
// Solidity: function GetMaxNumValidators() view returns(uint256)
func (_Staking *StakingSession) GetMaxNumValidators() (*big.Int, error) {
	return _Staking.Contract.GetMaxNumValidators(&_Staking.CallOpts)
}

// GetMaxNumValidators is a free data retrieval call binding the contract method 0x6edf2481.
//
// Solidity: function GetMaxNumValidators() view returns(uint256)
func (_Staking *StakingCallerSession) GetMaxNumValidators() (*big.Int, error) {
	return _Staking.Contract.GetMaxNumValidators(&_Staking.CallOpts)
}

// GetMaxNumWatchtowers is a free data retrieval call binding the contract method 0xb17ca89c.
//
// Solidity: function GetMaxNumWatchtowers() view returns(uint256)
func (_Staking *StakingCaller) GetMaxNumWatchtowers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetMaxNumWatchtowers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxNumWatchtowers is a free data retrieval call binding the contract method 0xb17ca89c.
//
// Solidity: function GetMaxNumWatchtowers() view returns(uint256)
func (_Staking *StakingSession) GetMaxNumWatchtowers() (*big.Int, error) {
	return _Staking.Contract.GetMaxNumWatchtowers(&_Staking.CallOpts)
}

// GetMaxNumWatchtowers is a free data retrieval call binding the contract method 0xb17ca89c.
//
// Solidity: function GetMaxNumWatchtowers() view returns(uint256)
func (_Staking *StakingCallerSession) GetMaxNumWatchtowers() (*big.Int, error) {
	return _Staking.Contract.GetMaxNumWatchtowers(&_Staking.CallOpts)
}

// GetMinNumParticipants is a free data retrieval call binding the contract method 0xbfc5aeee.
//
// Solidity: function GetMinNumParticipants() view returns(uint256)
func (_Staking *StakingCaller) GetMinNumParticipants(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetMinNumParticipants")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinNumParticipants is a free data retrieval call binding the contract method 0xbfc5aeee.
//
// Solidity: function GetMinNumParticipants() view returns(uint256)
func (_Staking *StakingSession) GetMinNumParticipants() (*big.Int, error) {
	return _Staking.Contract.GetMinNumParticipants(&_Staking.CallOpts)
}

// GetMinNumParticipants is a free data retrieval call binding the contract method 0xbfc5aeee.
//
// Solidity: function GetMinNumParticipants() view returns(uint256)
func (_Staking *StakingCallerSession) GetMinNumParticipants() (*big.Int, error) {
	return _Staking.Contract.GetMinNumParticipants(&_Staking.CallOpts)
}

// GetMinNumSequencers is a free data retrieval call binding the contract method 0x760118b6.
//
// Solidity: function GetMinNumSequencers() view returns(uint256)
func (_Staking *StakingCaller) GetMinNumSequencers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetMinNumSequencers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinNumSequencers is a free data retrieval call binding the contract method 0x760118b6.
//
// Solidity: function GetMinNumSequencers() view returns(uint256)
func (_Staking *StakingSession) GetMinNumSequencers() (*big.Int, error) {
	return _Staking.Contract.GetMinNumSequencers(&_Staking.CallOpts)
}

// GetMinNumSequencers is a free data retrieval call binding the contract method 0x760118b6.
//
// Solidity: function GetMinNumSequencers() view returns(uint256)
func (_Staking *StakingCallerSession) GetMinNumSequencers() (*big.Int, error) {
	return _Staking.Contract.GetMinNumSequencers(&_Staking.CallOpts)
}

// GetMinNumValidators is a free data retrieval call binding the contract method 0xeec9f52e.
//
// Solidity: function GetMinNumValidators() view returns(uint256)
func (_Staking *StakingCaller) GetMinNumValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetMinNumValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinNumValidators is a free data retrieval call binding the contract method 0xeec9f52e.
//
// Solidity: function GetMinNumValidators() view returns(uint256)
func (_Staking *StakingSession) GetMinNumValidators() (*big.Int, error) {
	return _Staking.Contract.GetMinNumValidators(&_Staking.CallOpts)
}

// GetMinNumValidators is a free data retrieval call binding the contract method 0xeec9f52e.
//
// Solidity: function GetMinNumValidators() view returns(uint256)
func (_Staking *StakingCallerSession) GetMinNumValidators() (*big.Int, error) {
	return _Staking.Contract.GetMinNumValidators(&_Staking.CallOpts)
}

// GetMinNumWatchtowers is a free data retrieval call binding the contract method 0xbd9ed0b3.
//
// Solidity: function GetMinNumWatchtowers() view returns(uint256)
func (_Staking *StakingCaller) GetMinNumWatchtowers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetMinNumWatchtowers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinNumWatchtowers is a free data retrieval call binding the contract method 0xbd9ed0b3.
//
// Solidity: function GetMinNumWatchtowers() view returns(uint256)
func (_Staking *StakingSession) GetMinNumWatchtowers() (*big.Int, error) {
	return _Staking.Contract.GetMinNumWatchtowers(&_Staking.CallOpts)
}

// GetMinNumWatchtowers is a free data retrieval call binding the contract method 0xbd9ed0b3.
//
// Solidity: function GetMinNumWatchtowers() view returns(uint256)
func (_Staking *StakingCallerSession) GetMinNumWatchtowers() (*big.Int, error) {
	return _Staking.Contract.GetMinNumWatchtowers(&_Staking.CallOpts)
}

// GetSlashPercentage is a free data retrieval call binding the contract method 0x57a87ffd.
//
// Solidity: function GetSlashPercentage() view returns(uint256)
func (_Staking *StakingCaller) GetSlashPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "GetSlashPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSlashPercentage is a free data retrieval call binding the contract method 0x57a87ffd.
//
// Solidity: function GetSlashPercentage() view returns(uint256)
func (_Staking *StakingSession) GetSlashPercentage() (*big.Int, error) {
	return _Staking.Contract.GetSlashPercentage(&_Staking.CallOpts)
}

// GetSlashPercentage is a free data retrieval call binding the contract method 0x57a87ffd.
//
// Solidity: function GetSlashPercentage() view returns(uint256)
func (_Staking *StakingCallerSession) GetSlashPercentage() (*big.Int, error) {
	return _Staking.Contract.GetSlashPercentage(&_Staking.CallOpts)
}

// IsSequencer is a free data retrieval call binding the contract method 0xbb5a26ff.
//
// Solidity: function IsSequencer(address addr) view returns(bool)
func (_Staking *StakingCaller) IsSequencer(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "IsSequencer", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequencer is a free data retrieval call binding the contract method 0xbb5a26ff.
//
// Solidity: function IsSequencer(address addr) view returns(bool)
func (_Staking *StakingSession) IsSequencer(addr common.Address) (bool, error) {
	return _Staking.Contract.IsSequencer(&_Staking.CallOpts, addr)
}

// IsSequencer is a free data retrieval call binding the contract method 0xbb5a26ff.
//
// Solidity: function IsSequencer(address addr) view returns(bool)
func (_Staking *StakingCallerSession) IsSequencer(addr common.Address) (bool, error) {
	return _Staking.Contract.IsSequencer(&_Staking.CallOpts, addr)
}

// IsValidator is a free data retrieval call binding the contract method 0x98ff822d.
//
// Solidity: function IsValidator(address addr) view returns(bool)
func (_Staking *StakingCaller) IsValidator(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "IsValidator", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0x98ff822d.
//
// Solidity: function IsValidator(address addr) view returns(bool)
func (_Staking *StakingSession) IsValidator(addr common.Address) (bool, error) {
	return _Staking.Contract.IsValidator(&_Staking.CallOpts, addr)
}

// IsValidator is a free data retrieval call binding the contract method 0x98ff822d.
//
// Solidity: function IsValidator(address addr) view returns(bool)
func (_Staking *StakingCallerSession) IsValidator(addr common.Address) (bool, error) {
	return _Staking.Contract.IsValidator(&_Staking.CallOpts, addr)
}

// IsWatchtower is a free data retrieval call binding the contract method 0xe2b5769d.
//
// Solidity: function IsWatchtower(address addr) view returns(bool)
func (_Staking *StakingCaller) IsWatchtower(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "IsWatchtower", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWatchtower is a free data retrieval call binding the contract method 0xe2b5769d.
//
// Solidity: function IsWatchtower(address addr) view returns(bool)
func (_Staking *StakingSession) IsWatchtower(addr common.Address) (bool, error) {
	return _Staking.Contract.IsWatchtower(&_Staking.CallOpts, addr)
}

// IsWatchtower is a free data retrieval call binding the contract method 0xe2b5769d.
//
// Solidity: function IsWatchtower(address addr) view returns(bool)
func (_Staking *StakingCallerSession) IsWatchtower(addr common.Address) (bool, error) {
	return _Staking.Contract.IsWatchtower(&_Staking.CallOpts, addr)
}

// NODESEQUENCER is a free data retrieval call binding the contract method 0x855598b7.
//
// Solidity: function NODE_SEQUENCER() view returns(string)
func (_Staking *StakingCaller) NODESEQUENCER(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "NODE_SEQUENCER")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NODESEQUENCER is a free data retrieval call binding the contract method 0x855598b7.
//
// Solidity: function NODE_SEQUENCER() view returns(string)
func (_Staking *StakingSession) NODESEQUENCER() (string, error) {
	return _Staking.Contract.NODESEQUENCER(&_Staking.CallOpts)
}

// NODESEQUENCER is a free data retrieval call binding the contract method 0x855598b7.
//
// Solidity: function NODE_SEQUENCER() view returns(string)
func (_Staking *StakingCallerSession) NODESEQUENCER() (string, error) {
	return _Staking.Contract.NODESEQUENCER(&_Staking.CallOpts)
}

// NODEVALIDATOR is a free data retrieval call binding the contract method 0x9a083e82.
//
// Solidity: function NODE_VALIDATOR() view returns(string)
func (_Staking *StakingCaller) NODEVALIDATOR(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "NODE_VALIDATOR")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NODEVALIDATOR is a free data retrieval call binding the contract method 0x9a083e82.
//
// Solidity: function NODE_VALIDATOR() view returns(string)
func (_Staking *StakingSession) NODEVALIDATOR() (string, error) {
	return _Staking.Contract.NODEVALIDATOR(&_Staking.CallOpts)
}

// NODEVALIDATOR is a free data retrieval call binding the contract method 0x9a083e82.
//
// Solidity: function NODE_VALIDATOR() view returns(string)
func (_Staking *StakingCallerSession) NODEVALIDATOR() (string, error) {
	return _Staking.Contract.NODEVALIDATOR(&_Staking.CallOpts)
}

// NODEWATCHTOWER is a free data retrieval call binding the contract method 0x9f9f7101.
//
// Solidity: function NODE_WATCHTOWER() view returns(string)
func (_Staking *StakingCaller) NODEWATCHTOWER(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "NODE_WATCHTOWER")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NODEWATCHTOWER is a free data retrieval call binding the contract method 0x9f9f7101.
//
// Solidity: function NODE_WATCHTOWER() view returns(string)
func (_Staking *StakingSession) NODEWATCHTOWER() (string, error) {
	return _Staking.Contract.NODEWATCHTOWER(&_Staking.CallOpts)
}

// NODEWATCHTOWER is a free data retrieval call binding the contract method 0x9f9f7101.
//
// Solidity: function NODE_WATCHTOWER() view returns(string)
func (_Staking *StakingCallerSession) NODEWATCHTOWER() (string, error) {
	return _Staking.Contract.NODEWATCHTOWER(&_Staking.CallOpts)
}

// AddressDisputedSequencerToWatchtower is a free data retrieval call binding the contract method 0xb1da3b13.
//
// Solidity: function _addressDisputedSequencerToWatchtower(address ) view returns(address)
func (_Staking *StakingCaller) AddressDisputedSequencerToWatchtower(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressDisputedSequencerToWatchtower", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressDisputedSequencerToWatchtower is a free data retrieval call binding the contract method 0xb1da3b13.
//
// Solidity: function _addressDisputedSequencerToWatchtower(address ) view returns(address)
func (_Staking *StakingSession) AddressDisputedSequencerToWatchtower(arg0 common.Address) (common.Address, error) {
	return _Staking.Contract.AddressDisputedSequencerToWatchtower(&_Staking.CallOpts, arg0)
}

// AddressDisputedSequencerToWatchtower is a free data retrieval call binding the contract method 0xb1da3b13.
//
// Solidity: function _addressDisputedSequencerToWatchtower(address ) view returns(address)
func (_Staking *StakingCallerSession) AddressDisputedSequencerToWatchtower(arg0 common.Address) (common.Address, error) {
	return _Staking.Contract.AddressDisputedSequencerToWatchtower(&_Staking.CallOpts, arg0)
}

// AddressDisputedSequencerToWatchtowerExists is a free data retrieval call binding the contract method 0x72c8af2d.
//
// Solidity: function _addressDisputedSequencerToWatchtowerExists(address ) view returns(bool)
func (_Staking *StakingCaller) AddressDisputedSequencerToWatchtowerExists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressDisputedSequencerToWatchtowerExists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddressDisputedSequencerToWatchtowerExists is a free data retrieval call binding the contract method 0x72c8af2d.
//
// Solidity: function _addressDisputedSequencerToWatchtowerExists(address ) view returns(bool)
func (_Staking *StakingSession) AddressDisputedSequencerToWatchtowerExists(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressDisputedSequencerToWatchtowerExists(&_Staking.CallOpts, arg0)
}

// AddressDisputedSequencerToWatchtowerExists is a free data retrieval call binding the contract method 0x72c8af2d.
//
// Solidity: function _addressDisputedSequencerToWatchtowerExists(address ) view returns(bool)
func (_Staking *StakingCallerSession) AddressDisputedSequencerToWatchtowerExists(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressDisputedSequencerToWatchtowerExists(&_Staking.CallOpts, arg0)
}

// AddressDisputedWatchtowerToSequencer is a free data retrieval call binding the contract method 0x5804e51c.
//
// Solidity: function _addressDisputedWatchtowerToSequencer(address ) view returns(address)
func (_Staking *StakingCaller) AddressDisputedWatchtowerToSequencer(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressDisputedWatchtowerToSequencer", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressDisputedWatchtowerToSequencer is a free data retrieval call binding the contract method 0x5804e51c.
//
// Solidity: function _addressDisputedWatchtowerToSequencer(address ) view returns(address)
func (_Staking *StakingSession) AddressDisputedWatchtowerToSequencer(arg0 common.Address) (common.Address, error) {
	return _Staking.Contract.AddressDisputedWatchtowerToSequencer(&_Staking.CallOpts, arg0)
}

// AddressDisputedWatchtowerToSequencer is a free data retrieval call binding the contract method 0x5804e51c.
//
// Solidity: function _addressDisputedWatchtowerToSequencer(address ) view returns(address)
func (_Staking *StakingCallerSession) AddressDisputedWatchtowerToSequencer(arg0 common.Address) (common.Address, error) {
	return _Staking.Contract.AddressDisputedWatchtowerToSequencer(&_Staking.CallOpts, arg0)
}

// AddressDisputedWatchtowerToSequencerExists is a free data retrieval call binding the contract method 0xbe31929a.
//
// Solidity: function _addressDisputedWatchtowerToSequencerExists(address ) view returns(bool)
func (_Staking *StakingCaller) AddressDisputedWatchtowerToSequencerExists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressDisputedWatchtowerToSequencerExists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddressDisputedWatchtowerToSequencerExists is a free data retrieval call binding the contract method 0xbe31929a.
//
// Solidity: function _addressDisputedWatchtowerToSequencerExists(address ) view returns(bool)
func (_Staking *StakingSession) AddressDisputedWatchtowerToSequencerExists(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressDisputedWatchtowerToSequencerExists(&_Staking.CallOpts, arg0)
}

// AddressDisputedWatchtowerToSequencerExists is a free data retrieval call binding the contract method 0xbe31929a.
//
// Solidity: function _addressDisputedWatchtowerToSequencerExists(address ) view returns(bool)
func (_Staking *StakingCallerSession) AddressDisputedWatchtowerToSequencerExists(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressDisputedWatchtowerToSequencerExists(&_Staking.CallOpts, arg0)
}

// AddressToDisputedWatchtowerIndex is a free data retrieval call binding the contract method 0x25d9b8c0.
//
// Solidity: function _addressToDisputedWatchtowerIndex(address ) view returns(uint256)
func (_Staking *StakingCaller) AddressToDisputedWatchtowerIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToDisputedWatchtowerIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToDisputedWatchtowerIndex is a free data retrieval call binding the contract method 0x25d9b8c0.
//
// Solidity: function _addressToDisputedWatchtowerIndex(address ) view returns(uint256)
func (_Staking *StakingSession) AddressToDisputedWatchtowerIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToDisputedWatchtowerIndex(&_Staking.CallOpts, arg0)
}

// AddressToDisputedWatchtowerIndex is a free data retrieval call binding the contract method 0x25d9b8c0.
//
// Solidity: function _addressToDisputedWatchtowerIndex(address ) view returns(uint256)
func (_Staking *StakingCallerSession) AddressToDisputedWatchtowerIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToDisputedWatchtowerIndex(&_Staking.CallOpts, arg0)
}

// AddressToIsDisputedWatchtower is a free data retrieval call binding the contract method 0xfa70f5e5.
//
// Solidity: function _addressToIsDisputedWatchtower(address ) view returns(bool)
func (_Staking *StakingCaller) AddressToIsDisputedWatchtower(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToIsDisputedWatchtower", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddressToIsDisputedWatchtower is a free data retrieval call binding the contract method 0xfa70f5e5.
//
// Solidity: function _addressToIsDisputedWatchtower(address ) view returns(bool)
func (_Staking *StakingSession) AddressToIsDisputedWatchtower(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsDisputedWatchtower(&_Staking.CallOpts, arg0)
}

// AddressToIsDisputedWatchtower is a free data retrieval call binding the contract method 0xfa70f5e5.
//
// Solidity: function _addressToIsDisputedWatchtower(address ) view returns(bool)
func (_Staking *StakingCallerSession) AddressToIsDisputedWatchtower(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsDisputedWatchtower(&_Staking.CallOpts, arg0)
}

// AddressToIsParticipant is a free data retrieval call binding the contract method 0xc70966fb.
//
// Solidity: function _addressToIsParticipant(address ) view returns(bool)
func (_Staking *StakingCaller) AddressToIsParticipant(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToIsParticipant", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddressToIsParticipant is a free data retrieval call binding the contract method 0xc70966fb.
//
// Solidity: function _addressToIsParticipant(address ) view returns(bool)
func (_Staking *StakingSession) AddressToIsParticipant(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsParticipant(&_Staking.CallOpts, arg0)
}

// AddressToIsParticipant is a free data retrieval call binding the contract method 0xc70966fb.
//
// Solidity: function _addressToIsParticipant(address ) view returns(bool)
func (_Staking *StakingCallerSession) AddressToIsParticipant(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsParticipant(&_Staking.CallOpts, arg0)
}

// AddressToIsSequencer is a free data retrieval call binding the contract method 0x7b339263.
//
// Solidity: function _addressToIsSequencer(address ) view returns(bool)
func (_Staking *StakingCaller) AddressToIsSequencer(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToIsSequencer", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddressToIsSequencer is a free data retrieval call binding the contract method 0x7b339263.
//
// Solidity: function _addressToIsSequencer(address ) view returns(bool)
func (_Staking *StakingSession) AddressToIsSequencer(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsSequencer(&_Staking.CallOpts, arg0)
}

// AddressToIsSequencer is a free data retrieval call binding the contract method 0x7b339263.
//
// Solidity: function _addressToIsSequencer(address ) view returns(bool)
func (_Staking *StakingCallerSession) AddressToIsSequencer(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsSequencer(&_Staking.CallOpts, arg0)
}

// AddressToIsSequencerInProbationAddr is a free data retrieval call binding the contract method 0x1efcf4d1.
//
// Solidity: function _addressToIsSequencerInProbationAddr(address ) view returns(bool)
func (_Staking *StakingCaller) AddressToIsSequencerInProbationAddr(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToIsSequencerInProbationAddr", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddressToIsSequencerInProbationAddr is a free data retrieval call binding the contract method 0x1efcf4d1.
//
// Solidity: function _addressToIsSequencerInProbationAddr(address ) view returns(bool)
func (_Staking *StakingSession) AddressToIsSequencerInProbationAddr(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsSequencerInProbationAddr(&_Staking.CallOpts, arg0)
}

// AddressToIsSequencerInProbationAddr is a free data retrieval call binding the contract method 0x1efcf4d1.
//
// Solidity: function _addressToIsSequencerInProbationAddr(address ) view returns(bool)
func (_Staking *StakingCallerSession) AddressToIsSequencerInProbationAddr(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsSequencerInProbationAddr(&_Staking.CallOpts, arg0)
}

// AddressToIsValidator is a free data retrieval call binding the contract method 0x065ae171.
//
// Solidity: function _addressToIsValidator(address ) view returns(bool)
func (_Staking *StakingCaller) AddressToIsValidator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToIsValidator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddressToIsValidator is a free data retrieval call binding the contract method 0x065ae171.
//
// Solidity: function _addressToIsValidator(address ) view returns(bool)
func (_Staking *StakingSession) AddressToIsValidator(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsValidator(&_Staking.CallOpts, arg0)
}

// AddressToIsValidator is a free data retrieval call binding the contract method 0x065ae171.
//
// Solidity: function _addressToIsValidator(address ) view returns(bool)
func (_Staking *StakingCallerSession) AddressToIsValidator(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsValidator(&_Staking.CallOpts, arg0)
}

// AddressToIsWatchtower is a free data retrieval call binding the contract method 0x40b00753.
//
// Solidity: function _addressToIsWatchtower(address ) view returns(bool)
func (_Staking *StakingCaller) AddressToIsWatchtower(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToIsWatchtower", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddressToIsWatchtower is a free data retrieval call binding the contract method 0x40b00753.
//
// Solidity: function _addressToIsWatchtower(address ) view returns(bool)
func (_Staking *StakingSession) AddressToIsWatchtower(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsWatchtower(&_Staking.CallOpts, arg0)
}

// AddressToIsWatchtower is a free data retrieval call binding the contract method 0x40b00753.
//
// Solidity: function _addressToIsWatchtower(address ) view returns(bool)
func (_Staking *StakingCallerSession) AddressToIsWatchtower(arg0 common.Address) (bool, error) {
	return _Staking.Contract.AddressToIsWatchtower(&_Staking.CallOpts, arg0)
}

// AddressToNodeType is a free data retrieval call binding the contract method 0x19ae6764.
//
// Solidity: function _addressToNodeType(address ) view returns(string)
func (_Staking *StakingCaller) AddressToNodeType(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToNodeType", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AddressToNodeType is a free data retrieval call binding the contract method 0x19ae6764.
//
// Solidity: function _addressToNodeType(address ) view returns(string)
func (_Staking *StakingSession) AddressToNodeType(arg0 common.Address) (string, error) {
	return _Staking.Contract.AddressToNodeType(&_Staking.CallOpts, arg0)
}

// AddressToNodeType is a free data retrieval call binding the contract method 0x19ae6764.
//
// Solidity: function _addressToNodeType(address ) view returns(string)
func (_Staking *StakingCallerSession) AddressToNodeType(arg0 common.Address) (string, error) {
	return _Staking.Contract.AddressToNodeType(&_Staking.CallOpts, arg0)
}

// AddressToParticipantIndex is a free data retrieval call binding the contract method 0xb8889f02.
//
// Solidity: function _addressToParticipantIndex(address ) view returns(uint256)
func (_Staking *StakingCaller) AddressToParticipantIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToParticipantIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToParticipantIndex is a free data retrieval call binding the contract method 0xb8889f02.
//
// Solidity: function _addressToParticipantIndex(address ) view returns(uint256)
func (_Staking *StakingSession) AddressToParticipantIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToParticipantIndex(&_Staking.CallOpts, arg0)
}

// AddressToParticipantIndex is a free data retrieval call binding the contract method 0xb8889f02.
//
// Solidity: function _addressToParticipantIndex(address ) view returns(uint256)
func (_Staking *StakingCallerSession) AddressToParticipantIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToParticipantIndex(&_Staking.CallOpts, arg0)
}

// AddressToSequencerInProbationIndex is a free data retrieval call binding the contract method 0x0e533c8e.
//
// Solidity: function _addressToSequencerInProbationIndex(address ) view returns(uint256)
func (_Staking *StakingCaller) AddressToSequencerInProbationIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToSequencerInProbationIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToSequencerInProbationIndex is a free data retrieval call binding the contract method 0x0e533c8e.
//
// Solidity: function _addressToSequencerInProbationIndex(address ) view returns(uint256)
func (_Staking *StakingSession) AddressToSequencerInProbationIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToSequencerInProbationIndex(&_Staking.CallOpts, arg0)
}

// AddressToSequencerInProbationIndex is a free data retrieval call binding the contract method 0x0e533c8e.
//
// Solidity: function _addressToSequencerInProbationIndex(address ) view returns(uint256)
func (_Staking *StakingCallerSession) AddressToSequencerInProbationIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToSequencerInProbationIndex(&_Staking.CallOpts, arg0)
}

// AddressToSequencerIndex is a free data retrieval call binding the contract method 0x1cb9bdca.
//
// Solidity: function _addressToSequencerIndex(address ) view returns(uint256)
func (_Staking *StakingCaller) AddressToSequencerIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToSequencerIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToSequencerIndex is a free data retrieval call binding the contract method 0x1cb9bdca.
//
// Solidity: function _addressToSequencerIndex(address ) view returns(uint256)
func (_Staking *StakingSession) AddressToSequencerIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToSequencerIndex(&_Staking.CallOpts, arg0)
}

// AddressToSequencerIndex is a free data retrieval call binding the contract method 0x1cb9bdca.
//
// Solidity: function _addressToSequencerIndex(address ) view returns(uint256)
func (_Staking *StakingCallerSession) AddressToSequencerIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToSequencerIndex(&_Staking.CallOpts, arg0)
}

// AddressToSequencerToDisputedWatchtower is a free data retrieval call binding the contract method 0x2b3cfe0e.
//
// Solidity: function _addressToSequencerToDisputedWatchtower(address ) view returns(uint256)
func (_Staking *StakingCaller) AddressToSequencerToDisputedWatchtower(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToSequencerToDisputedWatchtower", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToSequencerToDisputedWatchtower is a free data retrieval call binding the contract method 0x2b3cfe0e.
//
// Solidity: function _addressToSequencerToDisputedWatchtower(address ) view returns(uint256)
func (_Staking *StakingSession) AddressToSequencerToDisputedWatchtower(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToSequencerToDisputedWatchtower(&_Staking.CallOpts, arg0)
}

// AddressToSequencerToDisputedWatchtower is a free data retrieval call binding the contract method 0x2b3cfe0e.
//
// Solidity: function _addressToSequencerToDisputedWatchtower(address ) view returns(uint256)
func (_Staking *StakingCallerSession) AddressToSequencerToDisputedWatchtower(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToSequencerToDisputedWatchtower(&_Staking.CallOpts, arg0)
}

// AddressToStakedAmount is a free data retrieval call binding the contract method 0x7dceceb8.
//
// Solidity: function _addressToStakedAmount(address ) view returns(uint256)
func (_Staking *StakingCaller) AddressToStakedAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToStakedAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToStakedAmount is a free data retrieval call binding the contract method 0x7dceceb8.
//
// Solidity: function _addressToStakedAmount(address ) view returns(uint256)
func (_Staking *StakingSession) AddressToStakedAmount(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToStakedAmount(&_Staking.CallOpts, arg0)
}

// AddressToStakedAmount is a free data retrieval call binding the contract method 0x7dceceb8.
//
// Solidity: function _addressToStakedAmount(address ) view returns(uint256)
func (_Staking *StakingCallerSession) AddressToStakedAmount(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToStakedAmount(&_Staking.CallOpts, arg0)
}

// AddressToValidatorIndex is a free data retrieval call binding the contract method 0x02b75199.
//
// Solidity: function _addressToValidatorIndex(address ) view returns(uint256)
func (_Staking *StakingCaller) AddressToValidatorIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToValidatorIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToValidatorIndex is a free data retrieval call binding the contract method 0x02b75199.
//
// Solidity: function _addressToValidatorIndex(address ) view returns(uint256)
func (_Staking *StakingSession) AddressToValidatorIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToValidatorIndex(&_Staking.CallOpts, arg0)
}

// AddressToValidatorIndex is a free data retrieval call binding the contract method 0x02b75199.
//
// Solidity: function _addressToValidatorIndex(address ) view returns(uint256)
func (_Staking *StakingCallerSession) AddressToValidatorIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToValidatorIndex(&_Staking.CallOpts, arg0)
}

// AddressToWatchtowerIndex is a free data retrieval call binding the contract method 0xa0f14bff.
//
// Solidity: function _addressToWatchtowerIndex(address ) view returns(uint256)
func (_Staking *StakingCaller) AddressToWatchtowerIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_addressToWatchtowerIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToWatchtowerIndex is a free data retrieval call binding the contract method 0xa0f14bff.
//
// Solidity: function _addressToWatchtowerIndex(address ) view returns(uint256)
func (_Staking *StakingSession) AddressToWatchtowerIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToWatchtowerIndex(&_Staking.CallOpts, arg0)
}

// AddressToWatchtowerIndex is a free data retrieval call binding the contract method 0xa0f14bff.
//
// Solidity: function _addressToWatchtowerIndex(address ) view returns(uint256)
func (_Staking *StakingCallerSession) AddressToWatchtowerIndex(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.AddressToWatchtowerIndex(&_Staking.CallOpts, arg0)
}

// DisputeWatchtowers is a free data retrieval call binding the contract method 0x1b27dbc3.
//
// Solidity: function _dispute_watchtowers(uint256 ) view returns(address)
func (_Staking *StakingCaller) DisputeWatchtowers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_dispute_watchtowers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DisputeWatchtowers is a free data retrieval call binding the contract method 0x1b27dbc3.
//
// Solidity: function _dispute_watchtowers(uint256 ) view returns(address)
func (_Staking *StakingSession) DisputeWatchtowers(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.DisputeWatchtowers(&_Staking.CallOpts, arg0)
}

// DisputeWatchtowers is a free data retrieval call binding the contract method 0x1b27dbc3.
//
// Solidity: function _dispute_watchtowers(uint256 ) view returns(address)
func (_Staking *StakingCallerSession) DisputeWatchtowers(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.DisputeWatchtowers(&_Staking.CallOpts, arg0)
}

// MaximumNumParticipants is a free data retrieval call binding the contract method 0x068946e8.
//
// Solidity: function _maximumNumParticipants() view returns(uint256)
func (_Staking *StakingCaller) MaximumNumParticipants(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_maximumNumParticipants")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumNumParticipants is a free data retrieval call binding the contract method 0x068946e8.
//
// Solidity: function _maximumNumParticipants() view returns(uint256)
func (_Staking *StakingSession) MaximumNumParticipants() (*big.Int, error) {
	return _Staking.Contract.MaximumNumParticipants(&_Staking.CallOpts)
}

// MaximumNumParticipants is a free data retrieval call binding the contract method 0x068946e8.
//
// Solidity: function _maximumNumParticipants() view returns(uint256)
func (_Staking *StakingCallerSession) MaximumNumParticipants() (*big.Int, error) {
	return _Staking.Contract.MaximumNumParticipants(&_Staking.CallOpts)
}

// MaximumNumSequencers is a free data retrieval call binding the contract method 0x0fb18ccc.
//
// Solidity: function _maximumNumSequencers() view returns(uint256)
func (_Staking *StakingCaller) MaximumNumSequencers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_maximumNumSequencers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumNumSequencers is a free data retrieval call binding the contract method 0x0fb18ccc.
//
// Solidity: function _maximumNumSequencers() view returns(uint256)
func (_Staking *StakingSession) MaximumNumSequencers() (*big.Int, error) {
	return _Staking.Contract.MaximumNumSequencers(&_Staking.CallOpts)
}

// MaximumNumSequencers is a free data retrieval call binding the contract method 0x0fb18ccc.
//
// Solidity: function _maximumNumSequencers() view returns(uint256)
func (_Staking *StakingCallerSession) MaximumNumSequencers() (*big.Int, error) {
	return _Staking.Contract.MaximumNumSequencers(&_Staking.CallOpts)
}

// MaximumNumValidators is a free data retrieval call binding the contract method 0xaf6da36e.
//
// Solidity: function _maximumNumValidators() view returns(uint256)
func (_Staking *StakingCaller) MaximumNumValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_maximumNumValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumNumValidators is a free data retrieval call binding the contract method 0xaf6da36e.
//
// Solidity: function _maximumNumValidators() view returns(uint256)
func (_Staking *StakingSession) MaximumNumValidators() (*big.Int, error) {
	return _Staking.Contract.MaximumNumValidators(&_Staking.CallOpts)
}

// MaximumNumValidators is a free data retrieval call binding the contract method 0xaf6da36e.
//
// Solidity: function _maximumNumValidators() view returns(uint256)
func (_Staking *StakingCallerSession) MaximumNumValidators() (*big.Int, error) {
	return _Staking.Contract.MaximumNumValidators(&_Staking.CallOpts)
}

// MaximumNumWatchtowers is a free data retrieval call binding the contract method 0x90e7c862.
//
// Solidity: function _maximumNumWatchtowers() view returns(uint256)
func (_Staking *StakingCaller) MaximumNumWatchtowers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_maximumNumWatchtowers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumNumWatchtowers is a free data retrieval call binding the contract method 0x90e7c862.
//
// Solidity: function _maximumNumWatchtowers() view returns(uint256)
func (_Staking *StakingSession) MaximumNumWatchtowers() (*big.Int, error) {
	return _Staking.Contract.MaximumNumWatchtowers(&_Staking.CallOpts)
}

// MaximumNumWatchtowers is a free data retrieval call binding the contract method 0x90e7c862.
//
// Solidity: function _maximumNumWatchtowers() view returns(uint256)
func (_Staking *StakingCallerSession) MaximumNumWatchtowers() (*big.Int, error) {
	return _Staking.Contract.MaximumNumWatchtowers(&_Staking.CallOpts)
}

// MinStakingThreshold is a free data retrieval call binding the contract method 0xa1450126.
//
// Solidity: function _minStakingThreshold() view returns(uint256)
func (_Staking *StakingCaller) MinStakingThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_minStakingThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakingThreshold is a free data retrieval call binding the contract method 0xa1450126.
//
// Solidity: function _minStakingThreshold() view returns(uint256)
func (_Staking *StakingSession) MinStakingThreshold() (*big.Int, error) {
	return _Staking.Contract.MinStakingThreshold(&_Staking.CallOpts)
}

// MinStakingThreshold is a free data retrieval call binding the contract method 0xa1450126.
//
// Solidity: function _minStakingThreshold() view returns(uint256)
func (_Staking *StakingCallerSession) MinStakingThreshold() (*big.Int, error) {
	return _Staking.Contract.MinStakingThreshold(&_Staking.CallOpts)
}

// MinimumNumParticipants is a free data retrieval call binding the contract method 0xc6199c1f.
//
// Solidity: function _minimumNumParticipants() view returns(uint256)
func (_Staking *StakingCaller) MinimumNumParticipants(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_minimumNumParticipants")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumNumParticipants is a free data retrieval call binding the contract method 0xc6199c1f.
//
// Solidity: function _minimumNumParticipants() view returns(uint256)
func (_Staking *StakingSession) MinimumNumParticipants() (*big.Int, error) {
	return _Staking.Contract.MinimumNumParticipants(&_Staking.CallOpts)
}

// MinimumNumParticipants is a free data retrieval call binding the contract method 0xc6199c1f.
//
// Solidity: function _minimumNumParticipants() view returns(uint256)
func (_Staking *StakingCallerSession) MinimumNumParticipants() (*big.Int, error) {
	return _Staking.Contract.MinimumNumParticipants(&_Staking.CallOpts)
}

// MinimumNumSequencers is a free data retrieval call binding the contract method 0x8d084148.
//
// Solidity: function _minimumNumSequencers() view returns(uint256)
func (_Staking *StakingCaller) MinimumNumSequencers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_minimumNumSequencers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumNumSequencers is a free data retrieval call binding the contract method 0x8d084148.
//
// Solidity: function _minimumNumSequencers() view returns(uint256)
func (_Staking *StakingSession) MinimumNumSequencers() (*big.Int, error) {
	return _Staking.Contract.MinimumNumSequencers(&_Staking.CallOpts)
}

// MinimumNumSequencers is a free data retrieval call binding the contract method 0x8d084148.
//
// Solidity: function _minimumNumSequencers() view returns(uint256)
func (_Staking *StakingCallerSession) MinimumNumSequencers() (*big.Int, error) {
	return _Staking.Contract.MinimumNumSequencers(&_Staking.CallOpts)
}

// MinimumNumValidators is a free data retrieval call binding the contract method 0xc795c077.
//
// Solidity: function _minimumNumValidators() view returns(uint256)
func (_Staking *StakingCaller) MinimumNumValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_minimumNumValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumNumValidators is a free data retrieval call binding the contract method 0xc795c077.
//
// Solidity: function _minimumNumValidators() view returns(uint256)
func (_Staking *StakingSession) MinimumNumValidators() (*big.Int, error) {
	return _Staking.Contract.MinimumNumValidators(&_Staking.CallOpts)
}

// MinimumNumValidators is a free data retrieval call binding the contract method 0xc795c077.
//
// Solidity: function _minimumNumValidators() view returns(uint256)
func (_Staking *StakingCallerSession) MinimumNumValidators() (*big.Int, error) {
	return _Staking.Contract.MinimumNumValidators(&_Staking.CallOpts)
}

// MinimumNumWatchtowers is a free data retrieval call binding the contract method 0xf76a1ee0.
//
// Solidity: function _minimumNumWatchtowers() view returns(uint256)
func (_Staking *StakingCaller) MinimumNumWatchtowers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_minimumNumWatchtowers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumNumWatchtowers is a free data retrieval call binding the contract method 0xf76a1ee0.
//
// Solidity: function _minimumNumWatchtowers() view returns(uint256)
func (_Staking *StakingSession) MinimumNumWatchtowers() (*big.Int, error) {
	return _Staking.Contract.MinimumNumWatchtowers(&_Staking.CallOpts)
}

// MinimumNumWatchtowers is a free data retrieval call binding the contract method 0xf76a1ee0.
//
// Solidity: function _minimumNumWatchtowers() view returns(uint256)
func (_Staking *StakingCallerSession) MinimumNumWatchtowers() (*big.Int, error) {
	return _Staking.Contract.MinimumNumWatchtowers(&_Staking.CallOpts)
}

// Participants is a free data retrieval call binding the contract method 0xf196f824.
//
// Solidity: function _participants(uint256 ) view returns(address)
func (_Staking *StakingCaller) Participants(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_participants", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Participants is a free data retrieval call binding the contract method 0xf196f824.
//
// Solidity: function _participants(uint256 ) view returns(address)
func (_Staking *StakingSession) Participants(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.Participants(&_Staking.CallOpts, arg0)
}

// Participants is a free data retrieval call binding the contract method 0xf196f824.
//
// Solidity: function _participants(uint256 ) view returns(address)
func (_Staking *StakingCallerSession) Participants(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.Participants(&_Staking.CallOpts, arg0)
}

// Sequencers is a free data retrieval call binding the contract method 0x19908df8.
//
// Solidity: function _sequencers(uint256 ) view returns(address)
func (_Staking *StakingCaller) Sequencers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_sequencers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sequencers is a free data retrieval call binding the contract method 0x19908df8.
//
// Solidity: function _sequencers(uint256 ) view returns(address)
func (_Staking *StakingSession) Sequencers(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.Sequencers(&_Staking.CallOpts, arg0)
}

// Sequencers is a free data retrieval call binding the contract method 0x19908df8.
//
// Solidity: function _sequencers(uint256 ) view returns(address)
func (_Staking *StakingCallerSession) Sequencers(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.Sequencers(&_Staking.CallOpts, arg0)
}

// SequencersInProbation is a free data retrieval call binding the contract method 0x584060bb.
//
// Solidity: function _sequencers_in_probation(uint256 ) view returns(address)
func (_Staking *StakingCaller) SequencersInProbation(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_sequencers_in_probation", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencersInProbation is a free data retrieval call binding the contract method 0x584060bb.
//
// Solidity: function _sequencers_in_probation(uint256 ) view returns(address)
func (_Staking *StakingSession) SequencersInProbation(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.SequencersInProbation(&_Staking.CallOpts, arg0)
}

// SequencersInProbation is a free data retrieval call binding the contract method 0x584060bb.
//
// Solidity: function _sequencers_in_probation(uint256 ) view returns(address)
func (_Staking *StakingCallerSession) SequencersInProbation(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.SequencersInProbation(&_Staking.CallOpts, arg0)
}

// SlashPercentage is a free data retrieval call binding the contract method 0xd551832f.
//
// Solidity: function _slashPercentage() view returns(uint256)
func (_Staking *StakingCaller) SlashPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_slashPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashPercentage is a free data retrieval call binding the contract method 0xd551832f.
//
// Solidity: function _slashPercentage() view returns(uint256)
func (_Staking *StakingSession) SlashPercentage() (*big.Int, error) {
	return _Staking.Contract.SlashPercentage(&_Staking.CallOpts)
}

// SlashPercentage is a free data retrieval call binding the contract method 0xd551832f.
//
// Solidity: function _slashPercentage() view returns(uint256)
func (_Staking *StakingCallerSession) SlashPercentage() (*big.Int, error) {
	return _Staking.Contract.SlashPercentage(&_Staking.CallOpts)
}

// StakedAmount is a free data retrieval call binding the contract method 0xe387a7ed.
//
// Solidity: function _stakedAmount() view returns(uint256)
func (_Staking *StakingCaller) StakedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_stakedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedAmount is a free data retrieval call binding the contract method 0xe387a7ed.
//
// Solidity: function _stakedAmount() view returns(uint256)
func (_Staking *StakingSession) StakedAmount() (*big.Int, error) {
	return _Staking.Contract.StakedAmount(&_Staking.CallOpts)
}

// StakedAmount is a free data retrieval call binding the contract method 0xe387a7ed.
//
// Solidity: function _stakedAmount() view returns(uint256)
func (_Staking *StakingCallerSession) StakedAmount() (*big.Int, error) {
	return _Staking.Contract.StakedAmount(&_Staking.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0xf90ecacc.
//
// Solidity: function _validators(uint256 ) view returns(address)
func (_Staking *StakingCaller) Validators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_validators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0xf90ecacc.
//
// Solidity: function _validators(uint256 ) view returns(address)
func (_Staking *StakingSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.Validators(&_Staking.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xf90ecacc.
//
// Solidity: function _validators(uint256 ) view returns(address)
func (_Staking *StakingCallerSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.Validators(&_Staking.CallOpts, arg0)
}

// Watchtowers is a free data retrieval call binding the contract method 0xc48c5f37.
//
// Solidity: function _watchtowers(uint256 ) view returns(address)
func (_Staking *StakingCaller) Watchtowers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "_watchtowers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Watchtowers is a free data retrieval call binding the contract method 0xc48c5f37.
//
// Solidity: function _watchtowers(uint256 ) view returns(address)
func (_Staking *StakingSession) Watchtowers(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.Watchtowers(&_Staking.CallOpts, arg0)
}

// Watchtowers is a free data retrieval call binding the contract method 0xc48c5f37.
//
// Solidity: function _watchtowers(uint256 ) view returns(address)
func (_Staking *StakingCallerSession) Watchtowers(arg0 *big.Int) (common.Address, error) {
	return _Staking.Contract.Watchtowers(&_Staking.CallOpts, arg0)
}

// BeginDisputeResolution is a paid mutator transaction binding the contract method 0x6b2f7635.
//
// Solidity: function BeginDisputeResolution(address sequencerAddr) returns()
func (_Staking *StakingTransactor) BeginDisputeResolution(opts *bind.TransactOpts, sequencerAddr common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "BeginDisputeResolution", sequencerAddr)
}

// BeginDisputeResolution is a paid mutator transaction binding the contract method 0x6b2f7635.
//
// Solidity: function BeginDisputeResolution(address sequencerAddr) returns()
func (_Staking *StakingSession) BeginDisputeResolution(sequencerAddr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.BeginDisputeResolution(&_Staking.TransactOpts, sequencerAddr)
}

// BeginDisputeResolution is a paid mutator transaction binding the contract method 0x6b2f7635.
//
// Solidity: function BeginDisputeResolution(address sequencerAddr) returns()
func (_Staking *StakingTransactorSession) BeginDisputeResolution(sequencerAddr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.BeginDisputeResolution(&_Staking.TransactOpts, sequencerAddr)
}

// EndDisputeResolution is a paid mutator transaction binding the contract method 0xebe4ad6f.
//
// Solidity: function EndDisputeResolution(address sequencerAddr) returns()
func (_Staking *StakingTransactor) EndDisputeResolution(opts *bind.TransactOpts, sequencerAddr common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "EndDisputeResolution", sequencerAddr)
}

// EndDisputeResolution is a paid mutator transaction binding the contract method 0xebe4ad6f.
//
// Solidity: function EndDisputeResolution(address sequencerAddr) returns()
func (_Staking *StakingSession) EndDisputeResolution(sequencerAddr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.EndDisputeResolution(&_Staking.TransactOpts, sequencerAddr)
}

// EndDisputeResolution is a paid mutator transaction binding the contract method 0xebe4ad6f.
//
// Solidity: function EndDisputeResolution(address sequencerAddr) returns()
func (_Staking *StakingTransactorSession) EndDisputeResolution(sequencerAddr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.EndDisputeResolution(&_Staking.TransactOpts, sequencerAddr)
}

// SetMaxNumParticipants is a paid mutator transaction binding the contract method 0xe028aadf.
//
// Solidity: function SetMaxNumParticipants(uint256 maximumNumParticipants) returns(uint256)
func (_Staking *StakingTransactor) SetMaxNumParticipants(opts *bind.TransactOpts, maximumNumParticipants *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "SetMaxNumParticipants", maximumNumParticipants)
}

// SetMaxNumParticipants is a paid mutator transaction binding the contract method 0xe028aadf.
//
// Solidity: function SetMaxNumParticipants(uint256 maximumNumParticipants) returns(uint256)
func (_Staking *StakingSession) SetMaxNumParticipants(maximumNumParticipants *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxNumParticipants(&_Staking.TransactOpts, maximumNumParticipants)
}

// SetMaxNumParticipants is a paid mutator transaction binding the contract method 0xe028aadf.
//
// Solidity: function SetMaxNumParticipants(uint256 maximumNumParticipants) returns(uint256)
func (_Staking *StakingTransactorSession) SetMaxNumParticipants(maximumNumParticipants *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxNumParticipants(&_Staking.TransactOpts, maximumNumParticipants)
}

// SetMaxNumSequencers is a paid mutator transaction binding the contract method 0xf9a43d4c.
//
// Solidity: function SetMaxNumSequencers(uint256 maximumNumSequencers) returns(uint256)
func (_Staking *StakingTransactor) SetMaxNumSequencers(opts *bind.TransactOpts, maximumNumSequencers *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "SetMaxNumSequencers", maximumNumSequencers)
}

// SetMaxNumSequencers is a paid mutator transaction binding the contract method 0xf9a43d4c.
//
// Solidity: function SetMaxNumSequencers(uint256 maximumNumSequencers) returns(uint256)
func (_Staking *StakingSession) SetMaxNumSequencers(maximumNumSequencers *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxNumSequencers(&_Staking.TransactOpts, maximumNumSequencers)
}

// SetMaxNumSequencers is a paid mutator transaction binding the contract method 0xf9a43d4c.
//
// Solidity: function SetMaxNumSequencers(uint256 maximumNumSequencers) returns(uint256)
func (_Staking *StakingTransactorSession) SetMaxNumSequencers(maximumNumSequencers *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxNumSequencers(&_Staking.TransactOpts, maximumNumSequencers)
}

// SetMaxNumValidators is a paid mutator transaction binding the contract method 0xdea79774.
//
// Solidity: function SetMaxNumValidators(uint256 maximumNumValidators) returns(uint256)
func (_Staking *StakingTransactor) SetMaxNumValidators(opts *bind.TransactOpts, maximumNumValidators *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "SetMaxNumValidators", maximumNumValidators)
}

// SetMaxNumValidators is a paid mutator transaction binding the contract method 0xdea79774.
//
// Solidity: function SetMaxNumValidators(uint256 maximumNumValidators) returns(uint256)
func (_Staking *StakingSession) SetMaxNumValidators(maximumNumValidators *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxNumValidators(&_Staking.TransactOpts, maximumNumValidators)
}

// SetMaxNumValidators is a paid mutator transaction binding the contract method 0xdea79774.
//
// Solidity: function SetMaxNumValidators(uint256 maximumNumValidators) returns(uint256)
func (_Staking *StakingTransactorSession) SetMaxNumValidators(maximumNumValidators *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxNumValidators(&_Staking.TransactOpts, maximumNumValidators)
}

// SetMaxNumWatchtowers is a paid mutator transaction binding the contract method 0xd147c387.
//
// Solidity: function SetMaxNumWatchtowers(uint256 maximumNumWatchtowers) returns(uint256)
func (_Staking *StakingTransactor) SetMaxNumWatchtowers(opts *bind.TransactOpts, maximumNumWatchtowers *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "SetMaxNumWatchtowers", maximumNumWatchtowers)
}

// SetMaxNumWatchtowers is a paid mutator transaction binding the contract method 0xd147c387.
//
// Solidity: function SetMaxNumWatchtowers(uint256 maximumNumWatchtowers) returns(uint256)
func (_Staking *StakingSession) SetMaxNumWatchtowers(maximumNumWatchtowers *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxNumWatchtowers(&_Staking.TransactOpts, maximumNumWatchtowers)
}

// SetMaxNumWatchtowers is a paid mutator transaction binding the contract method 0xd147c387.
//
// Solidity: function SetMaxNumWatchtowers(uint256 maximumNumWatchtowers) returns(uint256)
func (_Staking *StakingTransactorSession) SetMaxNumWatchtowers(maximumNumWatchtowers *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMaxNumWatchtowers(&_Staking.TransactOpts, maximumNumWatchtowers)
}

// SetMinNumParticipants is a paid mutator transaction binding the contract method 0x98abb25e.
//
// Solidity: function SetMinNumParticipants(uint256 minimumNumParticipants) returns(uint256)
func (_Staking *StakingTransactor) SetMinNumParticipants(opts *bind.TransactOpts, minimumNumParticipants *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "SetMinNumParticipants", minimumNumParticipants)
}

// SetMinNumParticipants is a paid mutator transaction binding the contract method 0x98abb25e.
//
// Solidity: function SetMinNumParticipants(uint256 minimumNumParticipants) returns(uint256)
func (_Staking *StakingSession) SetMinNumParticipants(minimumNumParticipants *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinNumParticipants(&_Staking.TransactOpts, minimumNumParticipants)
}

// SetMinNumParticipants is a paid mutator transaction binding the contract method 0x98abb25e.
//
// Solidity: function SetMinNumParticipants(uint256 minimumNumParticipants) returns(uint256)
func (_Staking *StakingTransactorSession) SetMinNumParticipants(minimumNumParticipants *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinNumParticipants(&_Staking.TransactOpts, minimumNumParticipants)
}

// SetMinNumSequencers is a paid mutator transaction binding the contract method 0xb61b127f.
//
// Solidity: function SetMinNumSequencers(uint256 minimumNumSequencers) returns(uint256)
func (_Staking *StakingTransactor) SetMinNumSequencers(opts *bind.TransactOpts, minimumNumSequencers *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "SetMinNumSequencers", minimumNumSequencers)
}

// SetMinNumSequencers is a paid mutator transaction binding the contract method 0xb61b127f.
//
// Solidity: function SetMinNumSequencers(uint256 minimumNumSequencers) returns(uint256)
func (_Staking *StakingSession) SetMinNumSequencers(minimumNumSequencers *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinNumSequencers(&_Staking.TransactOpts, minimumNumSequencers)
}

// SetMinNumSequencers is a paid mutator transaction binding the contract method 0xb61b127f.
//
// Solidity: function SetMinNumSequencers(uint256 minimumNumSequencers) returns(uint256)
func (_Staking *StakingTransactorSession) SetMinNumSequencers(minimumNumSequencers *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinNumSequencers(&_Staking.TransactOpts, minimumNumSequencers)
}

// SetMinNumValidators is a paid mutator transaction binding the contract method 0x560854f6.
//
// Solidity: function SetMinNumValidators(uint256 minimumNumValidators) returns(uint256)
func (_Staking *StakingTransactor) SetMinNumValidators(opts *bind.TransactOpts, minimumNumValidators *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "SetMinNumValidators", minimumNumValidators)
}

// SetMinNumValidators is a paid mutator transaction binding the contract method 0x560854f6.
//
// Solidity: function SetMinNumValidators(uint256 minimumNumValidators) returns(uint256)
func (_Staking *StakingSession) SetMinNumValidators(minimumNumValidators *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinNumValidators(&_Staking.TransactOpts, minimumNumValidators)
}

// SetMinNumValidators is a paid mutator transaction binding the contract method 0x560854f6.
//
// Solidity: function SetMinNumValidators(uint256 minimumNumValidators) returns(uint256)
func (_Staking *StakingTransactorSession) SetMinNumValidators(minimumNumValidators *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinNumValidators(&_Staking.TransactOpts, minimumNumValidators)
}

// SetMinNumWatchtowers is a paid mutator transaction binding the contract method 0xd80ff2b9.
//
// Solidity: function SetMinNumWatchtowers(uint256 minimumNumWatchtowers) returns(uint256)
func (_Staking *StakingTransactor) SetMinNumWatchtowers(opts *bind.TransactOpts, minimumNumWatchtowers *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "SetMinNumWatchtowers", minimumNumWatchtowers)
}

// SetMinNumWatchtowers is a paid mutator transaction binding the contract method 0xd80ff2b9.
//
// Solidity: function SetMinNumWatchtowers(uint256 minimumNumWatchtowers) returns(uint256)
func (_Staking *StakingSession) SetMinNumWatchtowers(minimumNumWatchtowers *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinNumWatchtowers(&_Staking.TransactOpts, minimumNumWatchtowers)
}

// SetMinNumWatchtowers is a paid mutator transaction binding the contract method 0xd80ff2b9.
//
// Solidity: function SetMinNumWatchtowers(uint256 minimumNumWatchtowers) returns(uint256)
func (_Staking *StakingTransactorSession) SetMinNumWatchtowers(minimumNumWatchtowers *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinNumWatchtowers(&_Staking.TransactOpts, minimumNumWatchtowers)
}

// SetSlashPercentage is a paid mutator transaction binding the contract method 0x42ca58d5.
//
// Solidity: function SetSlashPercentage(uint256 newPercentage) returns()
func (_Staking *StakingTransactor) SetSlashPercentage(opts *bind.TransactOpts, newPercentage *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "SetSlashPercentage", newPercentage)
}

// SetSlashPercentage is a paid mutator transaction binding the contract method 0x42ca58d5.
//
// Solidity: function SetSlashPercentage(uint256 newPercentage) returns()
func (_Staking *StakingSession) SetSlashPercentage(newPercentage *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetSlashPercentage(&_Staking.TransactOpts, newPercentage)
}

// SetSlashPercentage is a paid mutator transaction binding the contract method 0x42ca58d5.
//
// Solidity: function SetSlashPercentage(uint256 newPercentage) returns()
func (_Staking *StakingTransactorSession) SetSlashPercentage(newPercentage *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetSlashPercentage(&_Staking.TransactOpts, newPercentage)
}

// SetStakingMinThreshold is a paid mutator transaction binding the contract method 0x2c404e9c.
//
// Solidity: function SetStakingMinThreshold(uint256 newThreshold) returns()
func (_Staking *StakingTransactor) SetStakingMinThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "SetStakingMinThreshold", newThreshold)
}

// SetStakingMinThreshold is a paid mutator transaction binding the contract method 0x2c404e9c.
//
// Solidity: function SetStakingMinThreshold(uint256 newThreshold) returns()
func (_Staking *StakingSession) SetStakingMinThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetStakingMinThreshold(&_Staking.TransactOpts, newThreshold)
}

// SetStakingMinThreshold is a paid mutator transaction binding the contract method 0x2c404e9c.
//
// Solidity: function SetStakingMinThreshold(uint256 newThreshold) returns()
func (_Staking *StakingTransactorSession) SetStakingMinThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetStakingMinThreshold(&_Staking.TransactOpts, newThreshold)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address slashAddr) returns()
func (_Staking *StakingTransactor) Slash(opts *bind.TransactOpts, slashAddr common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "slash", slashAddr)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address slashAddr) returns()
func (_Staking *StakingSession) Slash(slashAddr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Slash(&_Staking.TransactOpts, slashAddr)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address slashAddr) returns()
func (_Staking *StakingTransactorSession) Slash(slashAddr common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Slash(&_Staking.TransactOpts, slashAddr)
}

// Stake is a paid mutator transaction binding the contract method 0x46f45b8d.
//
// Solidity: function stake(string nodeType) payable returns()
func (_Staking *StakingTransactor) Stake(opts *bind.TransactOpts, nodeType string) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stake", nodeType)
}

// Stake is a paid mutator transaction binding the contract method 0x46f45b8d.
//
// Solidity: function stake(string nodeType) payable returns()
func (_Staking *StakingSession) Stake(nodeType string) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, nodeType)
}

// Stake is a paid mutator transaction binding the contract method 0x46f45b8d.
//
// Solidity: function stake(string nodeType) payable returns()
func (_Staking *StakingTransactorSession) Stake(nodeType string) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, nodeType)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Staking *StakingTransactor) Unstake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "unstake")
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Staking *StakingSession) Unstake() (*types.Transaction, error) {
	return _Staking.Contract.Unstake(&_Staking.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2def6620.
//
// Solidity: function unstake() returns()
func (_Staking *StakingTransactorSession) Unstake() (*types.Transaction, error) {
	return _Staking.Contract.Unstake(&_Staking.TransactOpts)
}

// StakingDisputeResolutionBeganIterator is returned from FilterDisputeResolutionBegan and is used to iterate over the raw logs and unpacked data for DisputeResolutionBegan events raised by the Staking contract.
type StakingDisputeResolutionBeganIterator struct {
	Event *StakingDisputeResolutionBegan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingDisputeResolutionBeganIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDisputeResolutionBegan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingDisputeResolutionBegan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingDisputeResolutionBeganIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDisputeResolutionBeganIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDisputeResolutionBegan represents a DisputeResolutionBegan event raised by the Staking contract.
type StakingDisputeResolutionBegan struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDisputeResolutionBegan is a free log retrieval operation binding the contract event 0x14ebe573ff6e00d7293f5d831eb94cade703155cd0910eb5316a7b2f17b14c35.
//
// Solidity: event DisputeResolutionBegan(address indexed account)
func (_Staking *StakingFilterer) FilterDisputeResolutionBegan(opts *bind.FilterOpts, account []common.Address) (*StakingDisputeResolutionBeganIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "DisputeResolutionBegan", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakingDisputeResolutionBeganIterator{contract: _Staking.contract, event: "DisputeResolutionBegan", logs: logs, sub: sub}, nil
}

// WatchDisputeResolutionBegan is a free log subscription operation binding the contract event 0x14ebe573ff6e00d7293f5d831eb94cade703155cd0910eb5316a7b2f17b14c35.
//
// Solidity: event DisputeResolutionBegan(address indexed account)
func (_Staking *StakingFilterer) WatchDisputeResolutionBegan(opts *bind.WatchOpts, sink chan<- *StakingDisputeResolutionBegan, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "DisputeResolutionBegan", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDisputeResolutionBegan)
				if err := _Staking.contract.UnpackLog(event, "DisputeResolutionBegan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeResolutionBegan is a log parse operation binding the contract event 0x14ebe573ff6e00d7293f5d831eb94cade703155cd0910eb5316a7b2f17b14c35.
//
// Solidity: event DisputeResolutionBegan(address indexed account)
func (_Staking *StakingFilterer) ParseDisputeResolutionBegan(log types.Log) (*StakingDisputeResolutionBegan, error) {
	event := new(StakingDisputeResolutionBegan)
	if err := _Staking.contract.UnpackLog(event, "DisputeResolutionBegan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDisputeResolutionEndedIterator is returned from FilterDisputeResolutionEnded and is used to iterate over the raw logs and unpacked data for DisputeResolutionEnded events raised by the Staking contract.
type StakingDisputeResolutionEndedIterator struct {
	Event *StakingDisputeResolutionEnded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingDisputeResolutionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDisputeResolutionEnded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingDisputeResolutionEnded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingDisputeResolutionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDisputeResolutionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDisputeResolutionEnded represents a DisputeResolutionEnded event raised by the Staking contract.
type StakingDisputeResolutionEnded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDisputeResolutionEnded is a free log retrieval operation binding the contract event 0x7da9849dcad1bb34324464458a33a9d0820286dedf66ec9c6669e730118d1c06.
//
// Solidity: event DisputeResolutionEnded(address indexed account)
func (_Staking *StakingFilterer) FilterDisputeResolutionEnded(opts *bind.FilterOpts, account []common.Address) (*StakingDisputeResolutionEndedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "DisputeResolutionEnded", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakingDisputeResolutionEndedIterator{contract: _Staking.contract, event: "DisputeResolutionEnded", logs: logs, sub: sub}, nil
}

// WatchDisputeResolutionEnded is a free log subscription operation binding the contract event 0x7da9849dcad1bb34324464458a33a9d0820286dedf66ec9c6669e730118d1c06.
//
// Solidity: event DisputeResolutionEnded(address indexed account)
func (_Staking *StakingFilterer) WatchDisputeResolutionEnded(opts *bind.WatchOpts, sink chan<- *StakingDisputeResolutionEnded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "DisputeResolutionEnded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDisputeResolutionEnded)
				if err := _Staking.contract.UnpackLog(event, "DisputeResolutionEnded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeResolutionEnded is a log parse operation binding the contract event 0x7da9849dcad1bb34324464458a33a9d0820286dedf66ec9c6669e730118d1c06.
//
// Solidity: event DisputeResolutionEnded(address indexed account)
func (_Staking *StakingFilterer) ParseDisputeResolutionEnded(log types.Log) (*StakingDisputeResolutionEnded, error) {
	event := new(StakingDisputeResolutionEnded)
	if err := _Staking.contract.UnpackLog(event, "DisputeResolutionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingSlashedIterator is returned from FilterSlashed and is used to iterate over the raw logs and unpacked data for Slashed events raised by the Staking contract.
type StakingSlashedIterator struct {
	Event *StakingSlashed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingSlashed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingSlashed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingSlashed represents a Slashed event raised by the Staking contract.
type StakingSlashed struct {
	Account          common.Address
	NewAmount        *big.Int
	SlashedAmount    *big.Int
	FeeRecipientAddr common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSlashed is a free log retrieval operation binding the contract event 0x10cc080888f97c96419786385100283aa8ae09efd09b75964b955212f42847d4.
//
// Solidity: event Slashed(address indexed account, uint256 newAmount, uint256 slashedAmount, address indexed feeRecipientAddr)
func (_Staking *StakingFilterer) FilterSlashed(opts *bind.FilterOpts, account []common.Address, feeRecipientAddr []common.Address) (*StakingSlashedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	var feeRecipientAddrRule []interface{}
	for _, feeRecipientAddrItem := range feeRecipientAddr {
		feeRecipientAddrRule = append(feeRecipientAddrRule, feeRecipientAddrItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Slashed", accountRule, feeRecipientAddrRule)
	if err != nil {
		return nil, err
	}
	return &StakingSlashedIterator{contract: _Staking.contract, event: "Slashed", logs: logs, sub: sub}, nil
}

// WatchSlashed is a free log subscription operation binding the contract event 0x10cc080888f97c96419786385100283aa8ae09efd09b75964b955212f42847d4.
//
// Solidity: event Slashed(address indexed account, uint256 newAmount, uint256 slashedAmount, address indexed feeRecipientAddr)
func (_Staking *StakingFilterer) WatchSlashed(opts *bind.WatchOpts, sink chan<- *StakingSlashed, account []common.Address, feeRecipientAddr []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	var feeRecipientAddrRule []interface{}
	for _, feeRecipientAddrItem := range feeRecipientAddr {
		feeRecipientAddrRule = append(feeRecipientAddrRule, feeRecipientAddrItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Slashed", accountRule, feeRecipientAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingSlashed)
				if err := _Staking.contract.UnpackLog(event, "Slashed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlashed is a log parse operation binding the contract event 0x10cc080888f97c96419786385100283aa8ae09efd09b75964b955212f42847d4.
//
// Solidity: event Slashed(address indexed account, uint256 newAmount, uint256 slashedAmount, address indexed feeRecipientAddr)
func (_Staking *StakingFilterer) ParseSlashed(log types.Log) (*StakingSlashed, error) {
	event := new(StakingSlashed)
	if err := _Staking.contract.UnpackLog(event, "Slashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Staking contract.
type StakingStakedIterator struct {
	Event *StakingStaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingStaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStaked represents a Staked event raised by the Staking contract.
type StakingStaked struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) FilterStaked(opts *bind.FilterOpts, account []common.Address) (*StakingStakedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Staked", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakedIterator{contract: _Staking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakingStaked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Staked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStaked)
				if err := _Staking.contract.UnpackLog(event, "Staked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) ParseStaked(log types.Log) (*StakingStaked, error) {
	event := new(StakingStaked)
	if err := _Staking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Staking contract.
type StakingUnstakedIterator struct {
	Event *StakingUnstaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnstaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingUnstaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnstaked represents a Unstaked event raised by the Staking contract.
type StakingUnstaked struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) FilterUnstaked(opts *bind.FilterOpts, account []common.Address) (*StakingUnstakedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Unstaked", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakingUnstakedIterator{contract: _Staking.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *StakingUnstaked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Unstaked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnstaked)
				if err := _Staking.contract.UnpackLog(event, "Unstaked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstaked is a log parse operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed account, uint256 amount)
func (_Staking *StakingFilterer) ParseUnstaked(log types.Log) (*StakingUnstaked, error) {
	event := new(StakingUnstaked)
	if err := _Staking.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
