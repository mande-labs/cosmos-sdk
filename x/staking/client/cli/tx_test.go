package cli

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
)

func TestPrepareConfigForTxCreateValidator(t *testing.T) {
	chainID := "chainID"
	ip := "1.1.1.1"
	nodeID := "nodeID"
	privKey := ed25519.GenPrivKey()
	valPubKey := privKey.PubKey()
	moniker := "DefaultMoniker"
	mkTxValCfg := func(amount string) TxCreateValidatorConfig {
		return TxCreateValidatorConfig{
			IP:                      ip,
			ChainID:                 chainID,
			NodeID:                  nodeID,
			PubKey:                  valPubKey,
			Moniker:                 moniker,
			Amount:                  amount,
		}
	}

	tests := []struct {
		name        string
		fsModify    func(fs *pflag.FlagSet)
		expectedCfg TxCreateValidatorConfig
	}{
		{
			name: "all defaults",
			fsModify: func(fs *pflag.FlagSet) {
				return
			},
			expectedCfg: mkTxValCfg(defaultAmount),
		}, {
			name: "Custom amount",
			fsModify: func(fs *pflag.FlagSet) {
				fs.Set(FlagAmount, "2000stake")
			},
			expectedCfg: mkTxValCfg("2000stake"),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			fs, _ := CreateValidatorMsgFlagSet(ip)
			fs.String(flags.FlagName, "", "name of private key with which to sign the gentx")

			tc.fsModify(fs)

			cvCfg, err := PrepareConfigForTxCreateValidator(fs, moniker, nodeID, chainID, valPubKey)
			require.NoError(t, err)

			require.Equal(t, tc.expectedCfg, cvCfg)
		})
	}
}
