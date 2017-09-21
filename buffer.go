package main

import (
	"os"
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
	os.Stdout.Write(self.tab)
}
