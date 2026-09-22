package main

import (
	"strconv"
	"strings"
)

func reverseBits(n int) int {
	binaryArr := [32]bool{}
	binNum := strconv.FormatInt(int64(n), 2)
	l := len(binNum)

	for i := 0; i < l; i++ {
		if binNum[l-1-i] == '1' {
			binaryArr[len(binaryArr)-1-i] = true
		} else {
			binaryArr[len(binaryArr)-1-i] = false
		}
	}

	var sb strings.Builder

	for i := 0; i < len(binaryArr); i++ {
		if binaryArr[len(binaryArr)-1-i] {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}
	}

	res, _ := strconv.ParseInt(sb.String(), 2, 0)

	return int(res)
}
