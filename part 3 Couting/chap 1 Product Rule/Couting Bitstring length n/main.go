package main

import (
	"fmt"
)

func main() {

	// init bitstring length n, assume n = 16
	mTask := [16]int{}

	// assume each bit i-th have
	for m := 0; m < len(mTask); m++ {
		i := 2
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
