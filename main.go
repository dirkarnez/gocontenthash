package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/crypto/md4"
)

func main() {
	var contentHash string
	var err error = nil
	if contentHash, err = GetContentHash("5cd1c5aa2fe84531e16ce555f52b1edf.png"); err != nil {
		log.Fatal("Program fails")
	}
	fmt.Printf("%t\n", contentHash == "5cd1c5aa2fe84531e16ce555f52b1edf")

	if contentHash, err = GetContentHash("f50a48f9b225510cf64ccef53469641f.png"); err != nil {
		log.Fatal("Program fails")
	}
	fmt.Printf("%t\n", contentHash == "f50a48f9b225510cf64ccef53469641f")
}

func GetContentHash(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	hash := md4.New()
	hash.Write(content)
	hashString := hex.EncodeToString(hash.Sum(nil)[:])
	return hashString, nil
}
