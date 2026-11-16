package main

import (
	"strconv"
	"strings"
)

func digitSum(s string, k int) string {

	for len(s) > k {
		count := 0
		sumNum := 0
		var sumStr strings.Builder

		for _, v := range s {
			count++
			sumNum += int(v - '0')
			if count == k {
				count = 0
				sumStr.WriteString(strconv.Itoa(sumNum))
				sumNum = 0
			}
		}
		if count > 0 {
			sumStr.WriteString(strconv.Itoa(sumNum))
		}
		s = sumStr.String()
	}

	return s
}
