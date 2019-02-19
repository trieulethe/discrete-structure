package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// init m tasks, assume we have 16 task
	mTask := [16]int{}

	// assume each task have i ways to do
	for m := 0; m < len(mTask); m++ {
		i := rand.Intn(7) + 1
		// fmt.Println("i:", i)
		mTask[m] = i
	}

	total := mTask[0]
	// total ways to do entire procedure:
	for m := 1; m < len(mTask); m++ {
		total = total * (mTask[m])
	}

	fmt.Println("total ways to do entire procedure: ", total)

}
