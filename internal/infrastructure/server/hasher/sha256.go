package hasher

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

type Sha256Hasher struct {
	key []byte
}

func NewSha256Hasher(key string) *Sha256Hasher {
	return &Sha256Hasher{
		key: []byte(key),
	}
}

func (h *Sha256Hasher) Hash(data string) string {
	return hex.EncodeToString(h.hash(data))
}

func (h *Sha256Hasher) Equal(data string, hash string) bool {
	decoded, err := hex.DecodeString(hash)
	if err != nil {
		return false
	}
	return hmac.Equal(h.hash(data), decoded)
}

func (h *Sha256Hasher) hash(data string) []byte {
	hash := hmac.New(sha256.New, h.key)
	hash.Write([]byte(data))

	return hash.Sum(nil)
}
