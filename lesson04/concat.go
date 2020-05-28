package lesson04

import _ "expvar"

func Concat(inSlice [][]int) []int {
	ret := []int{}
	for _, sl := range inSlice {
		ret = append(ret, sl...)
	}
	return ret
}
