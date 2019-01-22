package main

import (
	_"bufio"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

// func main()  {
// 	counts := make(map[string]int)
// 	input := bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 		fmt.Printf("%s %d",input.Text(),counts[input.Text()])
// 	}
// 	// fmt.Printf("%s",counts)
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n",n,line)
// 		}
// 	}
// }

// func main()  {
// 	counts := make(map[string]int)
// 	files := os.Args[1:]
// 	if len(files) == 0 {
// 		countLines(os.Stdin,counts)
// 	} else {
// 		for _,arg := range files {
// 			f,err := os.Open(arg)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr,"dup err: %v\n",err)
// 				continue
// 			}
// 			countLines(f,counts)
// 			f.Close()
// 		}
// 	}
// 	for line,n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n",n,line)
// 		}
// 	}
// }

// func countLines(f *os.File ,counts map[string]int) {
// 	input := bufio.NewScanner(f)
// 	for input.Scan() {
// 		counts[input.Text()]++
// 	}
// }

func main()  {
	counts := make(map[string]int)
	filenames := make(map[string][]string)
	for _,filename := range os.Args[1:] {
		data,err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr,"dup err: %v\n",err)
			continue
		}
		for _,line := range strings.Split(string(data),"\n") {
			counts[line]++
			files := filenames[line]
			// contain(files,filename)
			filenames[line] = append(filenames[line],filename)
		}
	}
	for line,n := range counts {
		// if n > 1 {
			fmt.Printf("%d\t%s\t%s\n",n,line,filenames[line])
		// }
	}
}

