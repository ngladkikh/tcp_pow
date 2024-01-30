package quotes

import (
	"math/rand"
)

var quotes = []string{
	"Foo said buz",
	"wise Foo does buz",
	"Wisdom is not buz but foo",
	"Make fiz but not buz",
	"Clever fuz has baz",
}

func Quote() string {
	randomIndex := rand.Intn(len(quotes))
	return quotes[randomIndex]
}
