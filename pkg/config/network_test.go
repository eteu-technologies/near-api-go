package config_test

import (
	"testing"

	"github.com/eteu-technologies/near-api-go/pkg/config"
)

func TestArchivalConfigLink(t *testing.T) {
	amni, ok := config.Networks["mainnet"].Archival()
	if !ok {
		t.Fatal("mainnet should have archival link")
	}

	mni, ok := amni.NonArchival()
	if !ok {
		t.Fatal("archival-mainnet should have non-archival link")
	}

	atni, ok := config.Networks["testnet"].Archival()
	if !ok {
		t.Fatal("testnet should have archival link")
	}

	tni, ok := atni.NonArchival()
	if !ok {
		t.Fatal("archival-testnet should have non-archival link")
	}

	_ = mni
	_ = tni
}
