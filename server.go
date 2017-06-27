package main

import (
	"net/http"
	"fmt"
)

var test_int = 10000.0

func tenPercent(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "110%% = %f\n", test_int * 1.1)
	fmt.Fprintf(w, " 90%% = %f\n", test_int * 0.9)
}

func main() {
	fmt.Print(" \n ---Hello World---\n")
	http.HandleFunc("/", tenPercent)
	http.ListenAndServe(":8080", nil)
}
