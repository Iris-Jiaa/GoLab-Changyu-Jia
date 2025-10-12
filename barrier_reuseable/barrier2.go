//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Changyu Jia
// Description:
// A simple barrier implemented using mutex and unbuffered channel
// Issues:
// None I hope
//1. Change mutex to atomic variable
//2. Make it a reusable barrier
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, arrived *int32, max int, wg *sync.WaitGroup, turnstile1 chan bool, turnstile2 chan bool) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)

	//we wait here until everyone has completed part A
	count := atomic.AddInt32(arrived, 1)

	if int(count) == max { //last to arrive -signal others to go
		for i := 0; i < max-1; i++ {
			turnstile1 <- true
		}
	} else { //not all here yet we wait until signal
		<-turnstile1
	} //end of if-else
	fmt.Println("PartB", goNum)

	count = atomic.AddInt32(arrived, -1)
	if int(count) == 0 { // last one to leave
		for i := 0; i < max-1; i++ {
			turnstile2 <- true
		}
	} else {
		<-turnstile2
	}

	wg.Done()
	return true
} //end-doStuff

func main() {
	totalRoutines := 10
	var arrived int32 = 0
	var wg sync.WaitGroup
	//we will need some of these
	turnstile1 := make(chan bool)
	turnstile2 := make(chan bool)
	for j := 0; j < 2; j++ { //run the whole thing twice to show reusability
		wg.Add(totalRoutines)
		for i := 0; i < totalRoutines; i++ { //create the go Routines here
			go doStuff(i, &arrived, totalRoutines, &wg, turnstile1, turnstile2)
		}
		wg.Wait() //wait for everyone to finish before exiting
	}
} //end-main
