package main

import (
	"fmt"
)

func main() {
	//fmt.Println(string('0'+9))
	fmt.Println(itoa(99999))
	fmt.Println(itoa(123123120123))
	fmt.Println(itoa(-55111))
}

func itoa(a int) string {
	var b string
	var a2 int
	if a < 0 {
		a2 = a * -1
	}
	if a == 0 {
		return string('0')
	}
	if a > 0 {
		a2 = a
	}
	for a2 > 0 {
		b = b + string('0'+a2%10)
		a2 = a2 / 10
	}
	if a < 0 {
		b = "-" + b
	}
	return b
}
