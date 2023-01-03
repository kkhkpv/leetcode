package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(array, k)
	fmt.Print(array)
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, i, j int) {

	for indx := 0; indx <= (j-i)/2; indx++ {
		nums[indx+i], nums[j-indx] = nums[j-indx], nums[indx]+i
	}
}
