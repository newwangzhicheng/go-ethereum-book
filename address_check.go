package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"regexp"
)

func main() {
	// 判断是否是一个合法的地址
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	isValid := re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")
	fmt.Printf("is valid address: %v\n", isValid)
	// 判断地址是否是一个合约
	/**
	判断在这个位置的是否有字节，如果有字节说明是一个合约
	*/
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	bytecode, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	if len(bytecode) > 0 {
		fmt.Println("this address is a contract")
	} else {
		fmt.Println("this address is not a contract, maybe this is an account")
	}
}
