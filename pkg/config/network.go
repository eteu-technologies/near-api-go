package config

import "fmt"

type NetworkInfo struct {
	NetworkID   string
	NodeURL     string
	WalletURL   string
	HelperURL   string
	ExplorerURL string

	archive    string
	nonArchive string
}

func (n NetworkInfo) Archival() (ni NetworkInfo, ok bool) {
	ni, ok = Networks[n.archive]
	return
}

func (n NetworkInfo) NonArchival() (ni NetworkInfo, ok bool) {
	ni, ok = Networks[n.nonArchive]
	return
}

func (n NetworkInfo) IsArchival() bool {
	return n.nonArchive != ""
}

func buildNetworkConfig(networkID string, hasArchival bool) (ni NetworkInfo) {
	ni.NetworkID = networkID
	ni.NodeURL = fmt.Sprintf("https://rpc.%s.near.org", networkID)
	ni.WalletURL = fmt.Sprintf("https://wallet.%s.near.org", networkID)
	ni.HelperURL = fmt.Sprintf("https://helper.%s.near.org", networkID)
	ni.ExplorerURL = fmt.Sprintf("https://explorer.%s.near.org", networkID)
	if hasArchival {
		ni.archive = fmt.Sprintf("archival-%s", networkID)
	}
	return
}

func buildArchivalNetworkConfig(networkID string) (ni NetworkInfo) {
	ni = buildNetworkConfig(networkID, false)
	ni.NetworkID = fmt.Sprintf("archival-%s", networkID)
	ni.NodeURL = fmt.Sprintf("https://archival-rpc.%s.near.org", networkID)
	ni.nonArchive = networkID
	return
}

var Networks = map[string]NetworkInfo{
	"mainnet": buildNetworkConfig("mainnet", true),
	"testnet": buildNetworkConfig("testnet", true),
	"betanet": buildNetworkConfig("betanet", false),
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
