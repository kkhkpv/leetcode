package main

import "fmt"

func main() {
	nums := []int{5, 25, 75}
	target := 100
	fmt.Printf("twoSum(nums, target): %v\n", twoSum(nums, target))
}

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		tempSum := numbers[left] + numbers[right]
		if tempSum == target {
			return []int{left + 1, right + 1}
		} else if tempSum > target {
			right--
		} else {
			left++
		}
	}
	return []int{-1}
}
