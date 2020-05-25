package itoa

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIota(t *testing.T) {
	assert.Equal(t, itoa(12313311), "12313311", "Хрень 12313311")
	assert.Equal(t, itoa(231312312312313), "231312312312313", "Хрень 231312312312313")
	assert.Equal(t, itoa(-1111), "-1111", "Хрень -11111")
	assert.Equal(t, itoa(0), "0", "Хрень -11111")
}
