func twoSum(nums []int, target int) []int {
	var indexes []int
	for i, _ := range nums {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target && i != j {
				if len(indexes) == 2 {
					break
				}
				indexes = append(indexes, i)
				indexes = append(indexes, j)
			}
		}
	}
	return indexes
}