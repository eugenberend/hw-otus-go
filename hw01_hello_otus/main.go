package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	sourceString := "Hello, OTUS!"
	resultString := stringutil.Reverse(sourceString)
	fmt.Println(resultString)
}
