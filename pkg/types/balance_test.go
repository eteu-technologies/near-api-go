package types_test

import (
	"testing"

	fuzz "github.com/google/gofuzz"

	. "github.com/eteu-technologies/near-api-go/pkg/types"
)

func TestNEARToYocto(t *testing.T) {
	var NEAR uint64 = 10

	yoctoValue := NEARToYocto(NEAR)
	orig := YoctoToNEAR(yoctoValue)

	if NEAR != orig {
		t.Errorf("expected: %d, got: %d", NEAR, orig)
	}
}

func TestNEARToYocto_Fuzz(t *testing.T) {
	f := fuzz.New()

	// TODO: ?
	var value uint16

	for i := 0; i < 1000; i++ {
		f.Fuzz(&value)
		newValue := YoctoToNEAR(NEARToYocto(uint64(value)))
		if uint64(value) != newValue {
			t.Errorf("expected: %d, got: %d", value, newValue)
		}
	}
}

// Tests for BalanceFromFloat and BalanceFromFloatNew

var valuesForTestBalanceFromFloat = map[float64]string{
	0.100000000000000000000000: "100000000000000000000000",
	0.340282366920938:          "340282366920938000000000",
	0.340282366920939:          "340282366920939000000000",
	0.340282366920940:          "340282366920940000000000",
	0.340282366920941:          "340282366920941000000000",
	340282366920938:            "340282366920938000000000000000000000000",
	3.40282366920938:           "3402823669209380000000000",
	34.0282366920938:           "34028236692093800000000000",
	340.282366920938:           "340282366920938000000000000",
	340.282:                    "340282000000000000000000000",
	340.28:                     "340280000000000000000000000",
	340.2:                      "340200000000000000000000000",
}

func TestBalanceFromFloat(t *testing.T) {
	for key, value := range valuesForTestBalanceFromFloat {
		bal := BalanceFromFloat(key)

		balString := bal.String()
		if value != balString {
			t.Errorf("Key: %.15f, expected: %s, got: %s", key, value, balString)
		}
	}
}

func TestBalanceFromFloatNew(t *testing.T) {
	for key, value := range valuesForTestBalanceFromFloat {
		bal := BalanceFromFloatNew(key)

		balString := bal.String()
		if value != balString {
			t.Errorf("Key: %.15f, expected: %s, got: %s", key, value, balString)
		}
	}
}

// Tests for BalanceFromString and BalanceFromStringNew

var valuesForTestBalanceFromString = map[string]string{
	"0.100000000000000000000000": "100000000000000000000000",
	"0.340282366920938":          "340282366920938000000000",
	"0.340282366920939":          "340282366920939000000000",
	"0.340282366920940":          "340282366920940000000000",
	"0.340282366920941":          "340282366920941000000000",
	"340282366920938":            "340282366920938000000000000000000000000",
	"3.40282366920938":           "3402823669209380000000000",
	"34.0282366920938":           "34028236692093800000000000",
	"340.282366920938":           "340282366920938000000000000",
	"340.282":                    "340282000000000000000000000",
	"340.28":                     "340280000000000000000000000",
	"340.2":                      "340200000000000000000000000",
}

func TestBalanceFromString(t *testing.T) {
	for key, value := range valuesForTestBalanceFromString {
		bal, err := BalanceFromString(key)
		if err != nil {
			t.Errorf("Key: %s, expected: %s, got: %s", key, value, err)
		}

		balString := bal.String()
		if value != balString {
			t.Errorf("Key: %s, expected: %s, got: %s", key, value, balString)
		}
	}
}

func TestBalanceFromStringNew(t *testing.T) {
	for key, value := range valuesForTestBalanceFromString {
		bal, err := BalanceFromStringNew(key)
		if err != nil {
			t.Errorf("Key: %s, expected: %s, got: %s", key, value, err)
		}

		balString := bal.String()
		if value != balString {
			t.Errorf("Key: %s, expected: %s, got: %s", key, value, balString)
		}
	}
}
