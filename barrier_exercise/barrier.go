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
//--------------------------------------------

// Implementing the barrier using mutex and semaphore
package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// Place a barrier in this function --use Mutex's and Semaphores

func doStuff(goNum int, wg *sync.WaitGroup, ctx context.Context, lock *sync.Mutex, sem *semaphore.Weighted, count *int, totalroutines int) bool {
	time.Sleep(time.Second)
	fmt.Println("PartA", goNum)

	lock.Lock()
	*count++
	currentCount := *count

	if currentCount == totalroutines {
		lock.Unlock()
		sem.Release(int64(totalroutines - 1))
	} else {
		lock.Unlock()
		sem.Acquire(ctx, 1)
	}

	fmt.Println("PartB", goNum)
	wg.Done()
	return true
}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	var count int
	var theLock sync.Mutex
	count = 0

	wg.Add(totalRoutines)
	ctx := context.TODO()

	sem := semaphore.NewWeighted(int64(totalRoutines - 1))

	if sem.Acquire(ctx, int64(totalRoutines-1)) != nil {
		fmt.Printf("Failed to acquire semaphore\n")
	}
	for i := 0; i < totalRoutines; i++ {
		go doStuff(i, &wg, ctx, &theLock, sem, &count, totalRoutines)
	}
	wg.Wait()
}
