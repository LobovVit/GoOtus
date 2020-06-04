package lesson08

import (
	"fmt"
	"testing"
)

func TestGRun(t *testing.T) {

	b1 := func() error {
		fmt.Println("1")
		return fmt.Errorf("XX")
	}
	b2 := func() error {
		fmt.Println("2")
		return fmt.Errorf("XX")
	}
	b3 := func() error {
		fmt.Println("3")
		return fmt.Errorf("XX")
	}
	b4 := func() error {
		fmt.Println("4")
		return fmt.Errorf("XX")
	}
	b5 := func() error {
		fmt.Println("5")
		return fmt.Errorf("XX")
	}
	b6 := func() error {
		fmt.Println("6")
		return fmt.Errorf("XX")
	}
	b7 := func() error {
		fmt.Println("7")
		return fmt.Errorf("XX")
	}
	b8 := func() error {
		fmt.Println("8")
		return fmt.Errorf("XX")
	}
	b9 := func() error {
		fmt.Println("9")
		return fmt.Errorf("XX")
	}
	b10 := func() error {
		fmt.Println("10")
		return fmt.Errorf("XX")
	}
	b11 := func() error {
		fmt.Println("11")
		return fmt.Errorf("XX")
	}

	sl := []func() error{b1, b2, b3, b4, b5, b6, b7, b8, b9, b10, b11}
	GRun(sl)
	//t.Error("Вот ошибки")
	t.Logf("Вот ошибки")
}
