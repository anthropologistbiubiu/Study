package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func sleepRandom(fromFunction string, ch chan int) {
	defer func() {
		fmt.Println(fromFunction, "sleepRandom complete")
	}()
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	randomNumber := r.Intn(100)
	sleeptime := randomNumber + 100
	fmt.Println(fromFunction, "Starting sleep for", sleeptime, "ms")
	time.Sleep(time.Duration(sleeptime) * time.Millisecond)
	fmt.Println(fromFunction, "Waking up, slept for ", sleeptime, "ms")
	if ch != nil {
		ch <- sleeptime
	}
}

func sleepRandomContext(ctx context.Context, ch chan bool) {

	defer func() {
		fmt.Println("sleepRandomContext complete")
		ch <- true
	}()
	sleeptimeChan := make(chan int)
	go sleepRandom("sleepRandomContext", sleeptimeChan)
	select {
	case <-ctx.Done():
		fmt.Println("sleepRandomContext: Time to return")
	case sleeptime := <-sleeptimeChan:
		fmt.Println("Slept for ", sleeptime, "ms")
	}
}

func doWorkContext(ctx context.Context) {

	ctxWithTimeout, cancelFunction := context.WithTimeout(ctx, time.Duration(150)*time.Millisecond)
	defer func() {
		fmt.Println("doWorkContext complete")
		cancelFunction()
	}()
	ch := make(chan bool)
	go sleepRandomContext(ctxWithTimeout, ch)
	select {
	case <-ctx.Done():
		fmt.Println("doWorkContext: Done")
	case <-ch:
		fmt.Println("sleepRandomContext ch")
	}
}

func main() {
	ctx := context.Background()
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)
	defer func() {
		fmt.Println("Main Defer: canceling context")
		cancelFunction()
	}()
	go func() {
		sleepRandom("Main", nil)
		//cancelFunction()
		fmt.Println("Main Sleep complete. canceling context")
	}()
	doWorkContext(ctxWithCancel)
	context.Background()
}
