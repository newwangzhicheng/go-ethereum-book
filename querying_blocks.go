package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	// 通过HeaderByNumber查看第n个节点区块的头信息
	// blockNumber为nil则查询最新的节点
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("latest block number:", header.Number.String())

	//通过BlockByNumber查看第n个节点区块的详细信息
	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("block number:", block.Number().Uint64())                // 第几个节点
	fmt.Println("block time:", block.Time())                             // 节点时间
	fmt.Println("block hash:", block.Hash().Hex())                       // 节点地址
	fmt.Println("block difficulty", block.Difficulty().Uint64())         // 节点复杂度
	fmt.Println("block transactions number:", len(block.Transactions())) // 交易数量

	// 根据节点的哈希值直接查询交易数量
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("transaction count", count)
}
