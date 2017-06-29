package main

import . "fmt"

func computedValue() int {
	return 5
}

func main() {
	var numbers3 = [5]int{1, 2, 3, 4, 5}
	slice3 := numbers3[2:len(numbers3)]
	x := computedValue()
	// length := (3)
	// capacity := (3)
	Printf("%d, %d, %d\n", len(slice3), cap(slice3), x)
}
