// Author: Oliwier Jakubiec
// Date: 21/09/2026
// Purpose: Mutex code example

package main

import (
	"fmt"
	"sync"
)

// Global variables shared between functions --A BAD IDEA
var wg sync.WaitGroup
var total int64

func adds(n int, theLock *sync.Mutex) bool {
	for i := 0; i < n; i++ {
		// Lock and unlock before modifying the total so that no other thread can access it
		theLock.Lock()
		total++
		theLock.Unlock()
	}
	wg.Done() //let waitgroup know we have finished
	return true
}

func main() {

	// theLock will be passed by reference between go routines
	// better than using a global variable
	var theLock sync.Mutex

	total = 0
	//the waitgroup is used as a barrier
	// init it to number of go routines
	wg.Add(10)

	//for loop using range option
	for i := range 10 {
		// start each routine
		fmt.Println(i)
		go adds(1000, &theLock)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done
	fmt.Println(total)
}
