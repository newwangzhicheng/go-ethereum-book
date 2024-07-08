package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"log"
	"os"
)

func main() {
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
	file := "./wallets/UTC--2024-07-04T09-59-21.910265000Z--af41dbf0713ef6e17eddcab73fb07cd430d48813"
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
