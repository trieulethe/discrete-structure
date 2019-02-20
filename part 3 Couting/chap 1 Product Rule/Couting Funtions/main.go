package main

// find function
func f(value int) int {
	return value + 1
}

func main() {
	// set A
	A := [5]int{2, 1, 4, 5, 3}
	// set B
	B := [20]int{}

	for i := 0; i < len(A); i++ {
		for j := len(B) - 1; j >= 0; j-- {
			// TODO find function
		}
	}

}
