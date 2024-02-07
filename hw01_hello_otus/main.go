package main

import (
	"fmt"
	"os"

	"golang.org/x/example/hello/reverse"
)

func main() {
	reversedString := reverse.String("Hello, OTUS!")

	_, err := fmt.Fprintln(os.Stdout, reversedString)
	if err != nil {
		fmt.Println(err)
	}
}
