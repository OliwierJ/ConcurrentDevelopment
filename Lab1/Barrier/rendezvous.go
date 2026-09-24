// rendezvous.go
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
// Date: 21/09/2026
// Title: Rendezvous example
// Purpose: Demonstrate how two channels can be used to create a rendezvous between two threads
// Help given : Mykhailo Balaker
//--------------------------------------------

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// WorkWithRendezvous simulates work and waits for the other thread to finish before proceeding.
// Rendezvous only works with two threads
func WorkWithRendezvous(wg *sync.WaitGroup, Num int, chan1 chan struct{}, chan2 chan struct{}) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) // wait random time amount
	fmt.Println("Part A", Num)

	// Rendezvous here
	if Num == 0 {
		chan1 <- struct{}{} // signal t2
		<-chan2             // wait for t2 to signal
	}
	if Num == 1 {
		chan2 <- struct{}{} // signal t1
		<-chan1             // wait for t1 to signal
	}

	fmt.Println("Part B", Num)
	wg.Done() // notify waitgroup
}

func main() {
	// Initialize waitgroup
	var wg sync.WaitGroup
	threadCount := 2

	// Create two channels
	chan1 := make(chan struct{}, 1)
	chan2 := make(chan struct{}, 1)

	// Init waitgroup to 2
	wg.Add(threadCount)

	// Run both threads
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, chan1, chan2)
	}

	wg.Wait() //wait here for everyone to finish

}
