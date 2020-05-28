package Lesson04

import (
	"sort"
	"strings"
)

func cnt(str string) [3]string {

	var all []string
	var retVal [3]string
	all = strings.Split(str, " ")
	distinctMap := make(map[string]int)
	var values int
	type line struct {
		word string
		cnt  int
	}

	for _, curWord := range all {
		values = distinctMap[curWord] + 1
		distinctMap[curWord] = values
	}

	orderedSlice := []line{}
	var t line

	for key, val := range distinctMap {
		t.word = key
		t.cnt = val
		orderedSlice = append(orderedSlice, t)
	}
	sort.Slice(orderedSlice, func(i, j int) bool {
		return orderedSlice[i].cnt > orderedSlice[j].cnt
	})

	for i, xword := range orderedSlice {
		retVal[i] = xword.word
		if i == 2 {
			break
		}
	}

	return retVal
}
