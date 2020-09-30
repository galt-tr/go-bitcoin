package main

import (
	"log"

	"github.com/bitcoinschema/go-bitcoin"
)

func main() {

	// Start with a private key
	privateKey, err := bitcoin.CreatePrivateKeyString()
	if err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

	// Create a pubkey
	var pubKey string
	if pubKey, err = bitcoin.PubKeyFromPrivateKey(privateKey); err != nil {
		log.Fatalf("error occurred: %s", err.Error())
	}

	// Success!
	log.Printf("created pubkey: %s from private key: %s", pubKey, privateKey)
}
