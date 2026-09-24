//Barrier.go
//Copyright (C) 2026 Oliwier Jakubiec

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
// Author: Oliwier Jakubiec
// Created on 21/09/2026
// Title: Barrier example
// Purpose: Demonstrate how a channels can be used to create a barrier between `n` threads
// Help Received : Mykhailo Balaker
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

// A function that does work, creates a barrier, and completes its work
func doStuff(goNum int, wg *sync.WaitGroup, channel chan struct{}, lock *sync.Mutex, counter *int) {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	// Wait here until everyone has completed part A

	lock.Lock()
	*counter-- // decrement counter
	// If every thread has passed through
	if *counter == 0 {
		lock.Unlock()
		fmt.Println("All threads have reached barrier")
		close(channel) // close the channel to release all the threads
	} else {
		lock.Unlock()
		<-channel // Wait for the channel to close
	}
	time.Sleep(time.Second)

	fmt.Println("Part B", goNum)
	wg.Done() // Notify the waitgroup
}

func main() {
	// Initialize waitgroup
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)

	// Create mutex lock
	var theLock sync.Mutex

	channel := make(chan struct{}) // Channel for waiting
	counter := totalRoutines       // Initialize counter

	// Launch go routines
	for i := range totalRoutines {
		go doStuff(i, &wg, channel, &theLock, &counter)
	}

	wg.Wait() // Wait for everyone to finish
}
