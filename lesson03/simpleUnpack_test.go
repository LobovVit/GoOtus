package itoa

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpack(t *testing.T) {
	assert.EqualValues(t, unpack("a4bc2d5e"), "aaaabccddddde", "Хрень aaaabccddddde")
	assert.EqualValues(t, unpack("a10bc2d5e"), "aaaaaaaaaabccddddde", "Хрень aaaaaaaaaabccddddde")
	assert.Equal(t, unpack("abcd"), "abcd", "Хрень abcd")
	assert.Equal(t, unpack("45"), "", "Хрень 45")
	assert.EqualValues(t, unpack("Ф10bc2d5e"), "ФФФФФФФФФФbccddddde", "Хрень ФФФФФФФФФФbccddddde")
	assert.EqualValues(t, unpack(`a2qwe\4`), `aaqwe4`, "Хрень qwe4")
	assert.EqualValues(t, unpack(`bqwe\45`), `bqwe44444`, "Хрень qwe44444")
	assert.EqualValues(t, unpack(`cqwe\\5qqq`), `cqwe\\\\\qqq`, `Хqweь\\\\\`)
}
