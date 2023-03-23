package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	data1 := []interface{}{"bisa1", "bisa2", "bisa3"}
	data2 := []interface{}{"coba1", "coba2", "coba3"}

	wg.Add(2)
	go acak(data1, &wg, &mu)
	go acak(data2, &wg, &mu)

	wg.Wait()
}

func acak(data interface{}, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	for i := 1; i <= 4; i++ {
		// time.Sleep(1 * time.Second)
		mu.Lock()
		fmt.Println(data, i)
		mu.Unlock()
	}
}

/*
https://levelup.gitconnected.com/mutex-examples-in-go-ad7c440461a4
*/
