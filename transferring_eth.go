package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	privateKey, err := crypto.HexToECDSA("de9be858da4a475276426320d5e9262ecfc3ba460bfac56360bfa6c4c28b4ee0")
	if err != nil {
		log.Fatal(err)
	}
	fromAddress := common.HexToAddress("0xdD2FD4581271e230360230F9337D5c0430Bf44C0")

	// 连接区块链
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	// 生成交易
	/**
	获取账户的nonce
	创建value，gasLimit，gasPrice，toAddress
	*/
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(200000000)
	gasLimit := uint64(210000)
	var data []byte
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	gasPrice, err := client.SuggestGasPrice(context.Background()) // 可以自己定，也可以获取以太坊推荐
	if err != nil {
		log.Fatal(err)
	}

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	// 用私钥签名交易
	/**
	获取chainID
	用privateKey生成签名的交易
	*/
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 广播交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("signedTx", signedTx.Hash().Hex())
}

func generateAccount() (privateKey *ecdsa.PrivateKey, fromAddress common.Address) {
	// 创建账户
	/**
	生成私钥对象
	生成公钥对象
	生成ECDSA公钥对象
	生成地址
	*/
	privateKey, err := crypto.HexToECDSA("fad8c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("public key is not of type *ecdsa.PublicKey")
	}
	fromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	return
}
