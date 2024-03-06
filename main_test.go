package main

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	var tests = []struct {
		nums    []int
		val     int
		count   int
		numsout []int
	}{
		{[]int{1, 2, 4, 2, 3, 5}, 2, 4, []int{1, 4, 3, 5}},
		{[]int{1, 2, 4, 2, 3}, 88, 5, []int{1, 2, 4, 2, 3}},
		{[]int{9, 1, 56, 83, 2, 34, 50, 89, 1}, 1, 7, []int{9, 56, 83, 2, 34, 50, 89}},
	}

	for _, test := range tests {

		t.Run("testing ", func(t *testing.T) {
			k, nums := RemoveElement(test.nums, test.val)

			if k != test.count || !reflect.DeepEqual(nums, test.numsout) {
				t.Errorf(`values not equal, failed: expected number of elements to be %d, got %d; 
				expected the elements to be %v, got %v`, test.count, k, test.numsout, nums)
			}
		})
	}
}
