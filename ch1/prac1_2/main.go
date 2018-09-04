package main

import (
	"os"
	"strconv"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		println(strconv.Itoa(i) + " " + os.Args[i])
	}
}
