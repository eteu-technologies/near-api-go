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

func buildArchivalNetworkConfig(networkID string) (ni NetworkInfo) {
	ni = buildNetworkConfig(networkID)
	ni.NodeURL = fmt.Sprintf("https://archival-rpc.%s.near.org", networkID)
	return
}

var Networks = map[string]NetworkInfo{
	"mainnet": buildNetworkConfig("mainnet"),
	"testnet": buildNetworkConfig("testnet"),
	"betanet": buildNetworkConfig("betanet"),
	"local": {
		NetworkID: "local",
		NodeURL:   "http://127.0.0.1:3030",
	},
	// From https://docs.near.org/docs/api/rpc#setup:
	// > Querying historical data (older than 5 epochs or ~2.5 days), you may get responses that the data is not available anymore.
	// > In that case, archival RPC nodes will come to your rescue
	"archival-mainnet": buildArchivalNetworkConfig("mainnet"),
	"archival-testnet": buildArchivalNetworkConfig("testnet"),
}
