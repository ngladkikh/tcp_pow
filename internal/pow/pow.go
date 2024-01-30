package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

type PoW interface {
	Validate(challenge string, nonce string) bool
}

type Sha256PoW struct {
	Difficulty int
}

func NewSha256PoW(difficulty int) PoW {
	return &Sha256PoW{Difficulty: difficulty}
}

func (p *Sha256PoW) Validate(challenge string, nonce string) bool {
	hash := calculateHash(challenge, nonce)
	return strings.HasPrefix(hash, strings.Repeat("0", p.Difficulty))
}

func calculateHash(data string, nonce string) string {
	input := data + nonce
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}
