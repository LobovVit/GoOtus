package main

import (
	flag "github.com/spf13/pflag"
	"lesson12/pkg"
	"log"
)

var firstParam string
var secondParam string
var thirdParam []string

func init() {
	flag.StringVar(&firstParam, "d", "", "Scan dir")
	flag.StringVar(&secondParam, "c", "", "Command")
	flag.StringSliceVar(&thirdParam, "p", nil, "Params (delimiter \",\")")

}

func main() {
	flag.Parse()
	//fmt.Printf("firstParam type=%T ; firstParam=%v\n", firstParam,firstParam)
	//fmt.Printf("secondParam type=%T ; secondParam=%v\n", secondParam,secondParam)
	env, err := pkg.ReadDir(firstParam)
	if err != nil {
		log.Fatal(err)
	}
	pkg.RunCmd(secondParam, thirdParam, env)
}
