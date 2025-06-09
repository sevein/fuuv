package main

import (
	"os"
	"strconv"
)

func main() {
	var code int
	if len(os.Args) > 1 {
		if i, err := strconv.Atoi(os.Args[1]); err == nil {
			code = i
		} else {
			panic(err)
		}
	}
	os.Exit(code)
}
