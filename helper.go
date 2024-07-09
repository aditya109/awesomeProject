package main

import (
	"strconv"
	"strings"
)

func joinInts(ints []int) string {
	strs := make([]string, len(ints))
	for i, val := range ints {
		strs[i] = strconv.Itoa(val)
	}
	return strings.Join(strs, " ")
}
