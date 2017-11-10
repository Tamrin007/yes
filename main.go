package main

import (
	"bufio"
	"os"
)

func main() {
	var expletive []byte
	if len(os.Args) == 1 {
		expletive = []byte("y\n")
	} else {
		expletive = []byte(os.Args[1] + "\n")
	}
	var buf = bufio.NewWriterSize(os.Stdout, 8096)

	for {
		buf.Write(expletive)
	}
}
