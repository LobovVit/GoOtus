package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestMainNtp(t *testing.T) {
	a := random(0, 100)
	if a > 50 {
		t.Errorf("Вот и ошибка %d", a)
	} else {
		t.Logf("Вот и нет ошибки %d", a)
	}
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	if min > max {
		return min
	} else {
		return rand.Intn(max-min) + min
	}
}
