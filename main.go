package main

import (
	"fmt"

	binarysearch "github.com/Arismonx/Algorithm.git/binarySearch"
)

func main() {
	d := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := binarysearch.BinarySearch(d[:], 10, 0, len(d)-1)
	fmt.Printf("I want search 10: Get %d index %d = d[%d]\n", d[result], result, result)
}
