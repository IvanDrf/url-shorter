package shorter

import (
	"crypto/md5"
	"encoding/hex"
)

const shortLength = 7

func ShortenUrl(src string) string {
	hash := md5.Sum([]byte(src))
	hexHash := hex.EncodeToString(hash[:])

	res := hexHash[:shortLength]
	return res
}
