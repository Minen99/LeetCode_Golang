package twoSum

func twoSum(nums []int, target int) []int {

	numMap := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := numMap[target-num]; ok {
			indices := []int{j, i}
			return indices
		}
		numMap[num] = i
	}
	return nil
}
