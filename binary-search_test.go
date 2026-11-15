package main

import "testing"

func TestBinarySearch(t *testing.T) {
	// Test Case
	tests := []struct {
		name     string
		d        []int
		x        int
		left     int
		right    int
		expected int
	}{
		{
			name:     "เจอตรงกลาง",
			d:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			x:        5,
			left:     0,
			right:    9,
			expected: 4,
		},
		{
			name:     "เจอตัวแรก",
			d:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			x:        1,
			left:     0,
			right:    9,
			expected: 0,
		},
		{
			name:     "เจอตัวท้ายสุด",
			d:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			x:        10,
			left:     0,
			right:    9,
			expected: 9,
		},
		{
			name:     "ไม่เจอ",
			d:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			x:        10,
			left:     0,
			right:    9,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := BinarySearch(tt.d, tt.x, tt.left, tt.right)
			if ans != tt.expected {
				t.Errorf("BinarySearch() ผลลัพธ์ได้ %d แต่คาดหวังไว้ %d", ans, tt.expected)
			}
		})
	}
}
