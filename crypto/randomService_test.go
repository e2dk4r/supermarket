package crypto_test

import (
	"testing"

	"github.com/e2dk4r/supermarket/crypto"
)

func TestGenerateStringThatIs32CharactersLong(t *testing.T) {
	expected := 32
	rs := crypto.RandomService{}

	str, err := rs.GenerateString(expected)
	if err != nil {
		t.Fatal(err)
	}

	actual := len(str)
	if actual != expected {
		t.Fatalf("str: %s expected: %d actual: %d", str, expected, actual)
	}
}
