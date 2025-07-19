package shorter

import (
	"crypto/md5"
	"encoding/hex"
)

type Shorter interface {
	ShortenUrl(src string) string
}

type shorten struct {
}

func NewShorten() Shorter {
	return shorten{}
}

const shortLength = 7

func (this shorten) ShortenUrl(src string) string {
	hash := md5.Sum([]byte(src))
	hexHash := hex.EncodeToString(hash[:])

	res := hexHash[:shortLength]
	return res
}
