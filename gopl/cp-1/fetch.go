package main

import (
	"fmt"
	_"io/ioutil"
	"io"
	"net/http"
	"os"
	"net/url"
)

func main()  {
	for _,urlString := range os.Args[1:] {
		u,e := url.Parse(urlString)
		if e != nil {
			fmt.Fprintf(os.Stderr,"url: %s  error:%v\n",urlString,e)
			os.Exit(1)
		}
		if u.Scheme == "" {
			fmt.Fprintf(os.Stdout,"url: %s\n",u.String())
			u.Scheme = "http"
			fmt.Fprintf(os.Stdout,"url: %s\n",u.String())
		}
		resp, err := http.Get(u.String())
		fmt.Fprintf(os.Stdout,"resp status: %s\n",resp.Status)

		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch :%v\n",err)
			os.Exit(1)
		}
		_,er := io.Copy(os.Stdout,resp.Body)
		if er != nil {
			fmt.Fprintf(os.Stderr,"fetch :reading %s:%v\n",urlString,er)
			os.Exit(1)
		}
		// b,err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr,"fetch :reading %s:%v\n",url,err)
		// 	os.Exit(1)
		// }

		// fmt.Printf("%s",b)
	}
}