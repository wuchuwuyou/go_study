package main

import (
	"os"
	"flag"
	"fmt"
)

var name string

var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init()  {
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	// flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)

	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.StringVar(&name,"name","everyone","The greeting object.")
}

func main(){
	// flag.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }
	// flag.Parse()
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello, %s!\n", name)

}