package types_test

import (
	"testing"

	"bridge/x/bridge/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				KeygenBlockList: []types.KeygenBlock{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				NodeAccountList: []types.NodeAccount{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				NodeAccountCount: 2,
				RegisterKeygenList: []types.RegisterKeygen{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated keygenBlock",
			genState: &types.GenesisState{
				KeygenBlockList: []types.KeygenBlock{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated nodeAccount",
			genState: &types.GenesisState{
				NodeAccountList: []types.NodeAccount{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid nodeAccount count",
			genState: &types.GenesisState{
				NodeAccountList: []types.NodeAccount{
					{
						Id: 1,
					},
				},
				NodeAccountCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated registerKeygen",
			genState: &types.GenesisState{
				RegisterKeygenList: []types.RegisterKeygen{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
