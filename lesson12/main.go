package main

import (
	"flag"
	"fmt"
	"lesson12/pkg"
)

func init() {
	flag.StringVar()
}

func main() {
	fmt.Printf("переменная %T\n", 1000)
	pkg.ReadDir()
	pkg.RunCmd()
}
