package binarysearch_test

import (
	"fmt"

	binarysearch "github.com/Arismonx/Algorithm.git/binarySearch"
)

func ExampleBinarySearch() {
	d := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarysearch.BinarySearch(d[:], 10, 0, len(d)-1))
	fmt.Println(binarysearch.BinarySearch(d[:], 11, 0, len(d)-1))
	// Output:
	// 9
	// -1
}
