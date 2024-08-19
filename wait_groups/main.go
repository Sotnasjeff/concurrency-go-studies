package main

import (
	"fmt"
	"sync"
)

func printSomething(value string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(value)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"Alpha",
		"Beta",
		"Gamma",
		"Pi",
		"Epsilon",
	}

	wg.Add(len(words))

	for i, v := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, v), &wg)
	}

	wg.Wait()

	wg.Add(1)
	go printSomething("Hello World", &wg)
	printSomething("Hello World 2", &wg)
}

//WaitGroups.Add() the amount of calls you do in a async function you put as a parameter of Add() function, 
//if you add a value much higher than the functions are being called, so you will have a deadlock.
