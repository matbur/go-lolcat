package main

import (
	"os"
	"fmt"
)

func main() {
	buff := createBuffer(128)
	n, _ := os.Stdin.Read(buff.tab)
	for n > 0 {
		fmt.Println(n)
		buff.print()
		buff.clear()
		n, _ = os.Stdin.Read(buff.tab)
	}
}
