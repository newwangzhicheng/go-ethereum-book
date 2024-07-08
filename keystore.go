package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//createKs()
	importKs()
}

// 创建钱包
func createKs() {
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex())
}

func importKs() {
	file := "./wallets/UTC--2024-07-08T13-07-39.293180000Z--aff286524d8bb0557e1d9f762b62ed82255f1de4"
	ks := keystore.NewKeyStore("./temp", keystore.StandardScryptN, keystore.StandardScryptP)
	// 读取文件
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// 导入钱包
	oldPassword := "secret"
	newPassword := "secret"
	account, err := ks.Import(jsonBytes, oldPassword, newPassword)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex())

	// 删除原文件
	if err = os.Remove(file); err != nil {
		log.Fatal(err)
	}
}
