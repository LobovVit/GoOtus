package Lesson04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCnt(t *testing.T) {
	t1val := "qwe qwe qwe asd asd asd zxc zxc zxc rty rty rty1 yui1 yui yui dfgv dfg AS as AS ZX WEFER ERGRE G GFER G  ASDQWDQD "
	t1ret := [3]string{"qwe", "asd", "zxc"}

	assert.Equal(t, cnt(t1val), t1ret, "Хрень {1,2,3,4,5,6,7,8,9}")

}
