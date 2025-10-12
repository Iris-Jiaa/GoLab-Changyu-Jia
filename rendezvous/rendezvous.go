package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

//Global variables shared between functions --A BAD IDEA

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, count *int, theLock *sync.Mutex, threadCount int, barrier chan bool) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("PartA", Num)
	//Rendezvous here

	theLock.Lock()
	*count++
	currentCount := *count

	if currentCount == threadCount {
		theLock.Unlock()
		close(barrier)
	} else {
		theLock.Unlock()
		<-barrier
	}

	fmt.Println("PartB", Num)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	barrier := make(chan bool)
	threadCount := 5
	count := 0
	var theLock sync.Mutex

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, &count, &theLock, threadCount, barrier)
	}
	wg.Wait()

}
