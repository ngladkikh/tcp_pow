package pow

import (
	"testing"
)

func TestSha256PoW_ValidateKnownPairs(t *testing.T) {
	pow := NewSha256PoW(3)

	challenge := "secret"
	examples := []string{
		"7131",
		"18227",
		"22921",
	}

	for _, example := range examples {
		if !pow.Validate(challenge, example) {
			t.Errorf("Expected nonce %s to be valid", example)
		}
	}
}

func TestStubValidate(t *testing.T) {
	validNonce := "foo"
	pow := NewStubPow(validNonce)

	if !pow.Validate("baz", validNonce) {
		t.Errorf("Expected nonce to be valid")
	}

	if pow.Validate("baz", "fiz") {
		t.Errorf("Expected nonce invalid")
	}
}
