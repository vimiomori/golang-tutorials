package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fn := os.Args[1]

	file, err := os.Open(fn)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		io.Copy(os.Stdout, file)
	}
}
