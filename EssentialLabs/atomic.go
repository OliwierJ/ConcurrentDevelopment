// Author: Oliwier Jakubiec
// Date: 21/09/2026
// Purpose: Atomic int example

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Global variables shared between functions --A BAD IDEA
var wg sync.WaitGroup

// Adds n to the atomic total
func addsAtomic(n int, total *atomic.Int64) bool {
	for i := 0; i < n; i++ {
		total.Add(1) // Add one to the atomic number
	}
	wg.Done() //let waitgroup know we have finished
	return true
}

func main() {
	var total atomic.Int64

	//the waitgroup is used as a barrier
	wg.Add(10)

	//for loop using range option
	for i := range 10 {
		fmt.Println("go Routine ", i)
		go addsAtomic(1000, &total)
	}

	wg.Wait() //wait here until everyone (10 go routines) is done

	// Print the total
	fmt.Println(total.Load())

}
