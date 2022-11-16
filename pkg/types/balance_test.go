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

	NEAR = 0
	yoctoValue = NEARToYocto(NEAR)
	orig = YoctoToNEAR(yoctoValue)

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

// Tests for Balance UnmarshalJSON

var valuesForTestBalanceUnmarshalJSON = map[string]string{
	`"0"`:            "0",
	`"1"`:            "1",
	`"100000000000"`: "100000000000",
	`"340282366920938000000000000000000000000"`: "340282366920938000000000000000000000000",
}

func TestBalanceUnmarshalJSON(t *testing.T) {
	for key, value := range valuesForTestBalanceUnmarshalJSON {
		var bal Balance
		err := bal.UnmarshalJSON([]byte(key))
		if err != nil {
			t.Errorf("Key: %s, expected: %s, got: %s", key, value, err)
		}

		balString := bal.String()
		if value != balString {
			t.Errorf("Key: %s, expected: %s, got: %s", key, value, balString)
		}
	}
}

var errorsForTestBalanceUnmarshalJSON = map[string]string{
	`"340.2.2"`: "unable to parse '340.2.2'",
	`"340.2"`:   "unable to parse '340.2'",
	`"abcd"`:    "unable to parse 'abcd'",
	"abcd":      "invalid character 'a' looking for beginning of value",
}

func TestBalanceUnmarshalJSONError(t *testing.T) {
	var bal Balance
	var err error

	for v, e := range errorsForTestBalanceUnmarshalJSON {
		err = bal.UnmarshalJSON([]byte(v))
		if err == nil || err.Error() != e {
			t.Errorf("Key: %s, expected: %s, got: %s", v, e, err)
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

func TestBalanceFromStringError(t *testing.T) {
	_, err := BalanceFromString("340.2.2")
	if err == nil {
		t.Errorf("expected error, got: nil")
	}
}
