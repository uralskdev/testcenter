package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	indexhtml string
)

func main() {
	indexfile, readerror := ioutil.ReadFile("index.html")
	if readerror != nil {
		log.Fatal("ReadFile ", readerror)
	}
	indexhtml = string(indexfile)
	go http.HandleFunc("/", index)
	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		request.ParseForm()
		fmt.Println(request.Form)
	}
	fmt.Fprintf(writer, indexhtml)
}

