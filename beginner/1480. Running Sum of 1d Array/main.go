func runningSum(nums []int) []int {
	for x := 1; x < len(nums); x++ {
		if x < len(nums) {
			nums[x] = nums[x] + nums[x-1]
		}
	}
	return nums
}