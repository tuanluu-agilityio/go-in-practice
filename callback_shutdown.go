package main

import (
	"fmt"
	"net/http"
	"os"
)

func homePage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage")
}

func shutdown(res http.ResponseWriter, req *http.Request) {
	os.Exit(0)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/shutdown", shutdown)
	http.ListenAndServe(":8080", nil)
}