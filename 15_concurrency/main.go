package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 300; i++ {
		fmt.Printf("Printing %d\n", i)
	}

}

func generateNumbers(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= n; i++ {
		fmt.Printf("Generated %d\n", i)
	}

}

func main() {
	fmt.Println(time.Now().Format("02-Jan-2006 15:04:05:123"))
	defer fmt.Println(time.Now().Format("02-Jan-2006 15:04:05:123"))
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbers(&wg)
	go generateNumbers(300, &wg)

	fmt.Println("Waitng for Go Routines to finish")
	wg.Wait()
	fmt.Println("Done")
}
