package main

import (
	_"fmt"
	"log"
	"net/http"
	"go_study/gopl/cp-3"
)

func main()  {
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe(":8888",nil))
}

func handler(w http.ResponseWriter,r *http.Request) {
	log.Fprintf(w,"URL.Path = %q\n",r.URL.Path)
	darw(w)
}