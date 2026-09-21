// rendezvous.go
// MIT License
//
// Copyright (c) 2026 Oliwier Jakubiec
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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
