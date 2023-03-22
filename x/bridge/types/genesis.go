package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		KeygenBlockList:    []KeygenBlock{},
		NodeAccountList:    []NodeAccount{},
		RegisterKeygenList: []RegisterKeygen{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in keygenBlock
	keygenBlockIndexMap := make(map[string]struct{})

	for _, elem := range gs.KeygenBlockList {
		index := string(KeygenBlockKey(elem.Index))
		if _, ok := keygenBlockIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for keygenBlock")
		}
		keygenBlockIndexMap[index] = struct{}{}
	}
	// Check for duplicated ID in nodeAccount
	nodeAccountIdMap := make(map[uint64]bool)
	nodeAccountCount := gs.GetNodeAccountCount()
	for _, elem := range gs.NodeAccountList {
		if _, ok := nodeAccountIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for nodeAccount")
		}
		if elem.Id >= nodeAccountCount {
			return fmt.Errorf("nodeAccount id should be lower or equal than the last id")
		}
		nodeAccountIdMap[elem.Id] = true
	}
	// Check for duplicated index in registerKeygen
	registerKeygenIndexMap := make(map[string]struct{})

	for _, elem := range gs.RegisterKeygenList {
		index := string(RegisterKeygenKey(elem.Index))
		if _, ok := registerKeygenIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for registerKeygen")
		}
		registerKeygenIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
