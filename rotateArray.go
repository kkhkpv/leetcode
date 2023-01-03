package leetcode

func rotate(nums []int, k int) {
	if k == len(nums) || len(nums) == 1 || k == 0 {
		return
	}
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, i, j int) {

	for indx := 0; indx <= (j-i)/2; indx++ {
		nums[indx+i], nums[j-indx] = nums[j-indx], nums[indx+i]
	}
}
