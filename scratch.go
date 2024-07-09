package main

import (
	"fmt"
)

func helloWorld() {
	fmt.Println("Hello World")
}

func main() {
	go helloWorld()
	fmt.Println("首先执行")
}
