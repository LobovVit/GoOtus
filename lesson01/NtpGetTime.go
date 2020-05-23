package main

import (
	"fmt"
	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("time.nist.gov")
	if err != nil {

		fmt.Println(err)
	}
	else {
		fmt.Println(time)
	}
	fmt.Println("Finish !!!!!! ")
}
