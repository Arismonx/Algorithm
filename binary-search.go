package main

import "fmt"

func BinarySearch(d []int, x int, left int, right int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if x == d[mid] {
		return mid
	}
	if x < d[mid] {
		return BinarySearch(d, x, left, mid-1)
	} else {
		return BinarySearch(d, x, mid+1, right)
	}
}

func main() {
	d := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := BinarySearch(d[:], 10, 0, len(d)-1)
	fmt.Printf("I want search 10: Get %d index %d = d[%d]\n", d[result], result, result)
}
