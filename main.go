package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	filePath := "README.md" // Replace with the actual path to your file

	// Read the file content
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new SHA256 hash instance
	hash := sha256.New()

	// Write the file content to the hash instance
	hash.Write(content)

	// Get the computed hash sum
	hashSum := hash.Sum(nil)

	// Convert the hash sum to hexadecimal string
	hashString := hex.EncodeToString(hashSum)

	fmt.Println(hashString)
}