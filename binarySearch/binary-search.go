package binarysearch

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
