package main

import (
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func TestMainNtp2(t *testing.T) {
	a := random(0, 100)
	require.Greaterf(t, a, 50, "хрень")
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	if min > max {
		return min
	} else {
		return rand.Intn(max-min) + min
	}
}
