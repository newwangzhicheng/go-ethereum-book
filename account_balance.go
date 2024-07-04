package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	// blockNumber为nil，则是最新的余额
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	blockNumber := big.NewInt(5532993)
	balanceAt, errAt := client.BalanceAt(context.Background(), address, blockNumber)
	if errAt != nil {
		log.Fatal(errAt)
	}
	fmt.Println(balanceAt)

	fBalance := new(big.Float)
	fBalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue)

	pendingBalance, pendingErr := client.PendingBalanceAt(context.Background(), address)
	if pendingErr != nil {
		log.Fatal(pendingErr)
	}
	fmt.Println(pendingBalance)
}
