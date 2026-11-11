class Solution:
    def minMoves(self, nums: List[int]) -> int:
        sum_nums = 0
        max_num = nums[0]
        for v in nums:
            sum_nums += v
            if v > max_num:
                max_num = v

        return max_num * len(nums) - sum_nums
