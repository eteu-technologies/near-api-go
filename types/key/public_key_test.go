package key_test

import (
	"testing"

	"github.com/eteu-technologies/near-api-go/types/key"
)

func TestED25519Key(t *testing.T) {
	expected := `ed25519:DcA2MzgpJbrUATQLLceocVckhhAqrkingax4oJ9kZ847`

	parsed, err := key.NewBase58PublicKey(expected)
	if err != nil {
		t.Errorf("failed to parse public key: %s", err)
	}

	if s := parsed.String(); s != expected {
		t.Errorf("%s != %s", s, expected)
	}
}
