package main

import (
	"os"
	"io"
)

func main() {
	buff := createBuffer(1)
	_, err := os.Stdin.Read(buff.tab)
	for err == nil {
		buff.print()
		buff.clear()
		_, err = os.Stdin.Read(buff.tab)
	}
	if err != io.EOF {
		panic(err)
	}
}
