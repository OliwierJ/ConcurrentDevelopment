// Author: Oliwier Jakubiec
// Date: 21/09/2026
// Purpose: Signalling example

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	barrier := make(chan bool)

	doStuffOne := func() bool {
		fmt.Println("StuffOne - Part A")

		// wait here for thread 2
		barrier <- true

		fmt.Println("StuffOne - PartB")
		wg.Done()
		return true
	}

	doStuffTwo := func() bool {
		time.Sleep(time.Second * 5)
		fmt.Println("StuffTwo - Part A")

		// signal thread 1
		<-barrier

		fmt.Println("StuffTwo - PartB")
		wg.Done()
		return true
	}
	wg.Add(2)
	go doStuffOne()
	go doStuffTwo()
	wg.Wait() //wait here for both threads to finish

}
