package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/kardianos/osext"
)

var head string

func main() {
	path, err := osext.ExecutableFolder()
	if err != nil {
		panic(err)
	}

	headArray, err := ioutil.ReadFile( path + "/src/head.html")
	if err != nil { panic(err) }
	head = string(headArray)

	http.HandleFunc("/", mainhandle)

	http.Handle("/src/", http.StripPrefix("/src/", http.FileServer(http.Dir(path))))

	fmt.Println("Started")
	http.ListenAndServe(":80", nil)
}

func mainhandle(iWrt http.ResponseWriter, iReq *http.Request) {
	fmt.Fprintln(iWrt, head)

	fmt.Println("Connected")
	fmt.Fprintln(iWrt, "Добрый день! Так как мне лень придумывать что-то новое, я просто отошлю то, что ты прислал мне\n")
	for name, val := range iReq.Header{
		fmt.Fprintf(iWrt, "%-30s%s\n", name, val)
	}
}