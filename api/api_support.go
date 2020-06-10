package api

import (
	"crypto/sha256"
	"encoding/base32"
	"strings"
)

func (img ImageRequest) Hash() string {
	hash := sha256.New()
	hash.Write([]byte(img.Platform))
	for _, dep := range img.Dependencies {
		hash.Write([]byte(dep))
	}
	return "v" + strings.ToLower(base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(hash.Sum(nil)))
}
