package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	msg = "Hello World"

	wg.Add(1)
	go updateMessage("Hello Cosmos", &wg)
	wg.Wait()

	printMessage()
	
	wg.Add(1)
	go updateMessage("Hello Gamma", &wg)
	wg.Wait()

	printMessage()

	wg.Add(1)
	go updateMessage("Hello Zeta", &wg)
	wg.Wait()

	printMessage()
}
