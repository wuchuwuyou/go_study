package main

import (
	"s2/lib"
)

var name string

func init() {
	flag.StringVar(&name,"name","everyone","The greeting object.")
}

func main() {
	flag.Parse()
	lib.hello(name)	
}