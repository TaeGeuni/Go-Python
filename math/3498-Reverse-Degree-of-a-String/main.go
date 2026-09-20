package main

func reverseDegree(s string) int {
	res := 0

	alphabetNums := [...]int{26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	for i := 0; i < len(s); i++ {
		res += alphabetNums[s[i]-96] * (i + 1)
	}

	return res
}
