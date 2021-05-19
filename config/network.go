package config

import "fmt"

type NetworkInfo struct {
	NetworkID   string
	NodeURL     string
	WalletURL   string
	HelperURL   string
	ExplorerURL string
}

func buildNetworkConfig(networkID string) NetworkInfo {
	return NetworkInfo{
		NetworkID:   networkID,
		NodeURL:     fmt.Sprintf("https://rpc.%s.near.org", networkID),
		WalletURL:   fmt.Sprintf("https://wallet.%s.near.org", networkID),
		HelperURL:   fmt.Sprintf("https://helper.%s.near.org", networkID),
		ExplorerURL: fmt.Sprintf("https://explorer.%s.near.org", networkID),
	}
}

var Networks = map[string]NetworkInfo{
	"mainnet": buildNetworkConfig("mainnet"),
	"testnet": buildNetworkConfig("testnet"),
	"betanet": buildNetworkConfig("betanet"),
	"local": {
		NetworkID: "local",
		NodeURL:   "http://127.0.0.1:3030",
	},
}
