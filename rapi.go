package main

import (
	"fmt"
	"sync"
	"time"
)

type Loop struct {
	num int
	sync.Mutex
}

func main() {
	var wg sync.WaitGroup
	var l Loop

	data1 := []interface{}{"coba1", "coba2", "coba3"}
	data2 := []interface{}{"bisa1", "bisa2", "bisa3"}

	wg.Add(2)
	go rapi(data1, &wg, &l)
	time.Sleep(1 * time.Second)
	go rapi(data2, &wg, &l)
	wg.Wait()

}

func rapi(data interface{}, wg *sync.WaitGroup, l *Loop) {
	defer wg.Done()
	for i := 1; i <= 4; i++ {
		l.Lock()
		// l.num++
		fmt.Println(data, i)
		l.Unlock()
		time.Sleep(3 * time.Second)
	}
}

/*
https://levelup.gitconnected.com/mutex-examples-in-go-ad7c440461a4
*/
