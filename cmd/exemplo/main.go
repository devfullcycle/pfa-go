package main

import (
	"fmt"
	"sync"
)

func contador() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}

func main() {
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// func() {
	// 	defer wg.Done()
	// 	contador()
	// }()
	// wg.Wait()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		contador()
		wg.Done()
	}()
	go func() {
		contador()
		wg.Done()
	}()
	wg.Wait()
}
