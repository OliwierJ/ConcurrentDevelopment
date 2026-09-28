//philosophers.go
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
// Created on 28/9/2026
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

const NumberPhilosophers = 5

var wg sync.WaitGroup

func left(i int) int { return i }

func right(i int) int {
	return (i + 1) % NumberPhilosophers
}

// Signals the footman and attempts to pick up left and right fork
func getForks(footman chan int, forks map[int]chan int, j int) {
	footman <- 0
	fmt.Printf("Phil %d waiting fork %d\n", j, left(j))
	forks[left(j)] <- 1
	fmt.Printf("Phil %d waiting fork %d\n", j, right(j))
	forks[right(j)] <- 1
}

// Sets down left and right forks and signals the footman
func putForks(footman chan int, forks map[int]chan int, j int) {
	<-forks[left(j)]
	<-forks[right(j)]
	fmt.Printf("Phil %d puts down forks %d, %d\n", j, left(j), right(j))
	<-footman
}

// Function that simulates a philosopher doing work.
// Each philosopher thinks, picks up two forks, eats, then puts the forks down.
func philosopherDoWork(goNum int, footman chan int, forks *map[int]chan int) {
	defer wg.Done()

	// Think
	fmt.Printf("Phil %d is thinking\n", goNum)
	time.Sleep(time.Second)

	// Pick up forks
	getForks(footman, *forks, goNum)

	// Eat
	fmt.Printf("Phil %d is eating\n", goNum)
	time.Sleep(time.Second)

	// Put down forks
	putForks(footman, *forks, goNum)
	fmt.Printf("Phil %d is done!\n", goNum)

}

func main() {
	wg.Add(NumberPhilosophers)

	// initialize footman
	footman := make(chan int, NumberPhilosophers-1)
	var forks = make(map[int]chan int)

	// initialize forks
	for i := range NumberPhilosophers {
		forks[i] = make(chan int, 1)
	}

	// Send start signal so can start
	footman <- 0

	// start up each philosopher
	for i := range NumberPhilosophers {
		go philosopherDoWork(i, footman, &forks)
	}

	wg.Wait()
}
