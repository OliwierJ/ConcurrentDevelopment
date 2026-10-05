//main.go
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
// Created on 01/10/2026
//--------------------------------------------

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Data struct {
	item int
}

func producer(buffer chan Data) {

	for {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		item := rand.Intn(100)
		fmt.Printf("Producer item = %d\n", item)

		buffer <- Data{item}

	}
}

func consumer(goNum int, buffer chan Data) {

	for {
		item := <-buffer
		fmt.Println("Consumer ", goNum, " consumed ", item.item)

		// Process item
		time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	}

}

func main() {
	bufferTotal := 10
	wg.Add(3)
	buffer := make(chan Data, bufferTotal)

	go producer(buffer)

	go consumer(1, buffer)
	go consumer(2, buffer)
	go consumer(3, buffer)

	wg.Wait()
}
