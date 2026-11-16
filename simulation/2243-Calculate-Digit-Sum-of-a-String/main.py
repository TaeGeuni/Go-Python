class Solution:
    def digitSum(self, s: str, k: int) -> str:

        while len(s) > k:
            count = 0
            sum_num = 0
            sum_str = []

            for value in s:
                count += 1
                sum_num += int(value)
                if count == k:
                    count = 0
                    sum_str.append(str(sum_num))
                    sum_num = 0
            if count > 0:
                sum_str.append(str(sum_num))
            s = "".join(sum_str)

        return s
