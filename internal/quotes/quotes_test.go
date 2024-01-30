package quotes

import (
	"testing"
)

func TestQuote(t *testing.T) {
	quote := Quote()
	if len(quote) == 0 {
		t.Fatalf("The quote '%s' is not in the known list of quotes", quote)
	}
}
