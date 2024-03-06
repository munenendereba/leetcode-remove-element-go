package main

import "fmt"

func main() {

	mm := []int{3, 2, 2, 3}

	count, nums := RemoveElement(mm, 3)

	fmt.Println(count, nums)
}

func RemoveElement(nums []int, val int) (int, []int) {
	k := 0
	count := len(nums)
	for i := count - 1; i >= 0; i-- {
		num := nums[i]
		if num != val {
			k++
			continue
		}
		//reassign the slice
		nums = append(nums[:i], nums[i+1:]...)
	}

	return k, nums
}
