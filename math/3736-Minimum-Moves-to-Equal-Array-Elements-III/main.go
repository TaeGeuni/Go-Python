package main

func minMoves(nums []int) int {
	sum := 0
	max := nums[0]

	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}

	return max*len(nums) - sum
}
