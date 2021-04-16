// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package constants

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/ava-labs/avalanchego/ids"
)

// Const variables to be exported
const (
	MainnetID uint32 = 1
	CascadeID uint32 = 2
	DenaliID  uint32 = 3
	EverestID uint32 = 4
	FujiID    uint32 = 5
	TokenizerID uint32 = 6

	TestnetID  uint32 = FujiID
	UnitTestID uint32 = 10
	LocalID    uint32 = 12345

	MainnetName  = "mainnet"
	CascadeName  = "cascade"
	DenaliName   = "denali"
	EverestName  = "everest"
	FujiName     = "fuji"
	TestnetName  = "testnet"
	UnitTestName = "testing"
	LocalName    = "local"
	TokenizerName = "tokenizer"

	MainnetHRP  = "avax"
	CascadeHRP  = "cascade"
	DenaliHRP   = "denali"
	EverestHRP  = "everest"
	FujiHRP     = "fuji"
	UnitTestHRP = "testing"
	LocalHRP    = "local"
	FallbackHRP = "custom"
	TokenizerHRP = "tokenizer"
)

// Variables to be exported
var (
	PrimaryNetworkID = ids.Empty
	PlatformChainID  = ids.Empty

	NetworkIDToNetworkName = map[uint32]string{
		MainnetID:  MainnetName,
		CascadeID:  CascadeName,
		DenaliID:   DenaliName,
		EverestID:  EverestName,
		FujiID:     FujiName,
		UnitTestID: UnitTestName,
		LocalID:    LocalName,
		TokenizerID: TokenizerName,
	}
	NetworkNameToNetworkID = map[string]uint32{
		MainnetName:  MainnetID,
		CascadeName:  CascadeID,
		DenaliName:   DenaliID,
		EverestName:  EverestID,
		FujiName:     FujiID,
		TestnetName:  TestnetID,
		UnitTestName: UnitTestID,
		LocalName:    LocalID,
		TokenizerName: TokenizerID,
	}

	NetworkIDToHRP = map[uint32]string{
		MainnetID:  MainnetHRP,
		CascadeID:  CascadeHRP,
		DenaliID:   DenaliHRP,
		EverestID:  EverestHRP,
		FujiID:     FujiHRP,
		UnitTestID: UnitTestHRP,
		LocalID:    LocalHRP,
		TokenizerID: TokenizerHRP,
	}
	NetworkHRPToNetworkID = map[string]uint32{
		MainnetHRP:  MainnetID,
		CascadeHRP:  CascadeID,
		DenaliHRP:   DenaliID,
		EverestHRP:  EverestID,
		FujiHRP:     FujiID,
		UnitTestHRP: UnitTestID,
		LocalHRP:    LocalID,
		TokenizerHRP: TokenizerID,
	}

	ValidNetworkName = regexp.MustCompile(`network-[0-9]+`)
)

// GetHRP returns the Human-Readable-Part of bech32 addresses for a networkID
func GetHRP(networkID uint32) string {
	if hrp, ok := NetworkIDToHRP[networkID]; ok {
		return hrp
	}
	return FallbackHRP
}

// NetworkName returns a human readable name for the network with
// ID [networkID]
func NetworkName(networkID uint32) string {
	if name, exists := NetworkIDToNetworkName[networkID]; exists {
		return name
	}
	return fmt.Sprintf("network-%d", networkID)
}

// NetworkID returns the ID of the network with name [networkName]
func NetworkID(networkName string) (uint32, error) {
	networkName = strings.ToLower(networkName)
	if id, exists := NetworkNameToNetworkID[networkName]; exists {
		return id, nil
	}

	if id, err := strconv.ParseUint(networkName, 10, 0); err == nil {
		if id > math.MaxUint32 {
			return 0, fmt.Errorf("networkID %s not in [0, 2^32)", networkName)
		}
		return uint32(id), nil
	}
	if ValidNetworkName.MatchString(networkName) {
		if id, err := strconv.Atoi(networkName[8:]); err == nil {
			if id > math.MaxUint32 {
				return 0, fmt.Errorf("networkID %s not in [0, 2^32)", networkName)
			}
			return uint32(id), nil
		}
	}

	return 0, fmt.Errorf("failed to parse %s as a network name", networkName)
}
