package main

import (
	"fmt"
	"log"
)

type Buffer struct {
	tab []byte
}

func createBuffer(n int) Buffer {
	b := Buffer{}
	b.tab = make([]byte, n)
	return b
}

func (self *Buffer) clear() {
	for i := range self.tab {
		self.tab[i] = 0
	}
}

func (self *Buffer) print() {
	log.Println("Buffer")
	for i, v := range self.tab {
		fmt.Printf("%d) |%v|\n", i, string(v))
	}
	//fmt.Println(string(self.tab))
}
