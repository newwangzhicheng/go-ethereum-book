package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connect to cloudflare eth")

	_ = client
}
