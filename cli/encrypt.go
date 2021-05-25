package main

import (
	"crypto/aes"
	"crypto/rand"
	"encpack/crypt"
	"encpack/pack"
	"flag"
	"log"
	"os"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func generateKey() []byte {
	key := make([]byte, 32)

	_, err := rand.Read(key)

	checkErr(err)

	pack.DoReorderAndPack(key, "k1", crypt.GenKeyFct("v1"))

	return key
}

func generateIV(key []byte) []byte {
	var iv = []byte(key)[:aes.BlockSize]

	pack.DoReorderAndPack(iv, "k2", crypt.GenKeyFct("v2"))

	return iv
}

func generateKeys() ([]byte, []byte) {
	key := generateKey()
	iv := generateIV(key)
	return key, iv
}

func generateToken(token string) {
	currentPath, err := os.Getwd()
	checkErr(err)

	keysPath := pack.CreateStructure("token")

	key, iv := generateKeys()

	encToken := crypt.DoEncodeFromString(token)

	tokenEncrypted := make([]byte, len(encToken))

	checkErr(crypt.DoEncrypt(tokenEncrypted, []byte(encToken), key, iv))

	pack.DoReorderAndPack(tokenEncrypted, "k3", crypt.GenKeyFct("v3"))
	pack.Compress(keysPath, currentPath, 3)
}

func generateCredentials(aid, asec, bucketName, targetName string) {
	currentPath, err := os.Getwd()
	checkErr(err)

	keysPath := pack.CreateStructure("credentials")

	key, iv := generateKeys()

	encAid := crypt.DoEncodeFromString(aid)
	encAsec := crypt.DoEncodeFromString(asec)
	encBucket := crypt.DoEncodeFromString(bucketName)
	encTarget := crypt.DoEncodeFromString(targetName)

	aidEncrypted := make([]byte, len(encAid))
	asecEncrypted := make([]byte, len(encAsec))
	bktEncrypted := make([]byte, len(encBucket))
	tgtEncrypted := make([]byte, len(encTarget))

	checkErr(crypt.DoEncrypt(aidEncrypted, []byte(encAid), key, iv))
	checkErr(crypt.DoEncrypt(asecEncrypted, []byte(encAsec), key, iv))
	checkErr(crypt.DoEncrypt(bktEncrypted, []byte(encBucket), key, iv))
	checkErr(crypt.DoEncrypt(tgtEncrypted, []byte(encTarget), key, iv))

	pack.DoReorderAndPack(aidEncrypted, "k3", crypt.GenKeyFct("v3"))
	pack.DoReorderAndPack(asecEncrypted, "k4", crypt.GenKeyFct("v4"))
	pack.DoReorderAndPack(bktEncrypted, "k5", crypt.GenKeyFct("v5"))
	pack.DoReorderAndPack(tgtEncrypted, "k6", crypt.GenKeyFct("v6"))

	pack.Compress(keysPath, currentPath, 6)
}

func main() {
	token := flag.String("token", "", "Token")

	aid := flag.String("awsid", "", "AWS access key")
	asec := flag.String("awssec", "", "AWS secret key")
	bucket := flag.String("bucket", "", "Bucket name")
	target := flag.String("target", "", "Target name")

	flag.Parse()

	if *token != "" {
		generateToken(*token)
	} else if *aid != "" && *asec != "" && *bucket != "" && *target != "" {
		generateCredentials(*aid, *asec, *bucket, *target)
	} else {
		log.Println("Você deve informar o token (-token) ou as credenciais (-awsid, -awssec e -bucket) para a geração dos arquivos")
		return
	}

	log.Println("Os dados foram criptografados com sucesso")
}
