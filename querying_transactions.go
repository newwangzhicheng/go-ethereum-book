package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
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

		// 获取receipt
		/**
		包含交易的执行结果，交易金额等信息
		*/
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("transaction status", receipt.Status) // 1代表成功，0代表失败
		fmt.Println("transaction log", receipt.Logs)
	}

	// 通过区块的哈希和交易的索引直接获取交易信息
	blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
	tx, err := client.TransactionInBlock(context.Background(), blockHash, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("transaction hash", tx.Hash().Hex())

	// 直接通过交易的哈希获取交易信息
	transactionHash := common.HexToHash("0xbf5085065b62f89b43ce44ae1d029e9750615dd35519aa3bf367baa4f379a9b4")
	tx, isPending, err := client.TransactionByHash(context.Background(), transactionHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("transaction hash", tx.Hash().Hex())
	fmt.Println("isPending", isPending)
}
