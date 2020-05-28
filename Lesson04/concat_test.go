package Lesson04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcat(t *testing.T) {
	t1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	t1Res := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, Concat(t1), t1Res, "Хрень {1,2,3,4,5,6,7,8,9}")
}
