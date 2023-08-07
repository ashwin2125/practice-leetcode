func twoSum(nums []int, target int) []int {

	numsMap := make(map[int]int)

	for i, num := range nums {
		rem := target - num
		if index, found := numsMap[rem]; found {
			return []int{index, i}
		}
		numsMap[num] = i
	}

	return []int{}
}