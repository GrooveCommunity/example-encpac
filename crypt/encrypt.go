package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encpack/pack"
	"math/rand"
	"strconv"
)

func GenKeyFct(fileName string) uint64 {
	var buffer bytes.Buffer
	keyFct := uint64(rand.Int63n(9999999999999))

	buffer.WriteString(strconv.FormatUint(keyFct, 10))
	pack.DoPackage(buffer, fileName)

	return keyFct
}

func DoEncrypt(dst, src, key, iv []byte) error {
	aesBlockEncrypter, err := aes.NewCipher(key)

	if err != nil {
		return err
	}

	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(dst, src)

	return nil
}
