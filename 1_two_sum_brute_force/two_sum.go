package twoSum

func twoSum(nums []int, target int) []int {

	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if target == (nums[i] + nums[j]) {
				indices := []int{i, j}
				return indices
			}
		}
	}
	return nil
}
