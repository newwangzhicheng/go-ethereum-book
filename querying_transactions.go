package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	// 获取transaction相关的信息
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockByNumber(context.Background(), big.NewInt(52455))
	if err != nil {
		log.Fatal(err)
	}
	for _, tx := range block.Transactions() {
		fmt.Println("transaction hash ", tx.Hash().Hex())             // 交易的哈希
		fmt.Println("transaction value ", tx.Value().String())        // 交易中所转账的金额
		fmt.Println("transaction gas ", tx.Gas())                     // 交易的gas
		fmt.Println("transaction gas price ", tx.GasPrice().Uint64()) // gas的价格
		fmt.Println("transaction nonce  ", tx.Nonce())                // 交易的随机数
		fmt.Println("transaction data ", tx.Data())                   // 交易中输入的数据
		fmt.Println("transaction receiver ", tx.To().Hex())           // 交易的对象

		// 获取交易来自哪里
		/**
		先获取chainID
		再获取sender
		*/
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		if from, err := types.Sender(types.NewLondonSigner(chainID), tx); err == nil {
			fmt.Println("transaction sender ", from.Hex())
		}
	}

}
