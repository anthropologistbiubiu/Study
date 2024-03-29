package main

import (
	"fmt"
	"sync"
)

// 全局变量
var counter int

func main() {
	/*
		var wg sync.WaitGroup
		for i := 0; i < 1000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				counter++
			}()
		}

		wg.Wait()
		fmt.Println(counter)
	*/
	pant1()
}

func pant1() {
	var wg sync.WaitGroup
	var l sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			counter++
			l.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
