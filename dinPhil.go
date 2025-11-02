// Dining Philosophers Template Code
// Author: Joseph Kehoe
// Created: 21/10/24
//GPL Licence
// MISSING:
// 1. Readme
// 2. Full licence info.
// 3. Comments
// 4. It can Deadlock!

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

var semaphore = make(chan bool, 4)

func think(index int) {
	X := time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was thinking")
}

func eat(index int) {
	X := time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was eating")
}

func getForks(index int, forks map[int]chan bool) {
	if index == 0 {
		forks[(index+1)%5] <- true
		forks[index] <- true
	} else {
		forks[index] <- true
		forks[(index+1)%5] <- true
	}

}

func putForks(index int, forks map[int]chan bool) {
	if index == 0 {
		<-forks[(index+1)%5]
		<-forks[index]
	} else {
		<-forks[index]
		<-forks[(index+1)%5]
	}
}

func doPhilStuff(index int, forks map[int]chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		think(index)
		semaphore <- true
		getForks(index, forks)
		eat(index)
		putForks(index, forks)
		<-semaphore
	}
}

func main() {
	var wg sync.WaitGroup
	philCount := 5
	wg.Add(philCount)

	forks := make(map[int]chan bool)
	for k := range philCount {
		forks[k] = make(chan bool, 1)
	} //set up forks
	for N := range philCount {
		go doPhilStuff(N, forks, &wg)
	} //start philosophers
	wg.Wait() //wait here until everyone (10 go routines) is done

} //main
