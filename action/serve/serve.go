package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe("", nil)
}

func handDefault() {
	fmt.Println(11)
}
