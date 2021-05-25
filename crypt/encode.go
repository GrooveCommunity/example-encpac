package crypt

import (
	b64 "encoding/base64"
)

func DoEncodeFromString(token string) string {
	return b64.StdEncoding.EncodeToString([]byte(token))
}

func DoEncodeFromBytes(b []byte) string {
	return b64.URLEncoding.EncodeToString(b)
}
