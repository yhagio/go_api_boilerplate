package hmachash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"hash"
)

// HMAC interface
type HMAC interface {
	Hash(input string) string
}

type hm struct {
	hmac hash.Hash
}

// NewHMAC instantiates HMAC utils
func NewHMAC(key string) hm {
	h := hmac.New(sha256.New, []byte(key))
	return hm{
		hmac: h,
	}
}

// Hash will hash the provided input string using HMAC with
// the secret key provided when the HMAC object was created
func (h hm) Hash(input string) string {
	h.hmac.Reset()
	h.hmac.Write([]byte(input))
	hashedData := h.hmac.Sum(nil)
	return base64.URLEncoding.EncodeToString(hashedData)
}
