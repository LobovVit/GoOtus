package itoa

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIota(t *testing.T) {
	require.Equal(t, itoa(12313311), "12313311", "Хрень 12313311")
	require.Equal(t, itoa(231312312312313), "231312312312313", "Хрень 231312312312313")
	require.Equal(t, itoa(-1111), "-1111", "Хрень -1111")
}
