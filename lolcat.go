package main

import (
	"os"
	"bytes"
	"fmt"
	"math"
)

type LolCat struct {
	input    *os.File
	output   *os.File
	tab      []string
	num      int
	freq     float64
	spread   float64
	seed     int
	animate  bool
	duration int
	speed    float64
	force    bool
}

func CreateLolCat(num int) LolCat {
	lc := LolCat{
		input:    os.Stdin,
		output:   os.Stdout,
		num:      num,
		freq:     .1,
		spread:   3.,
		seed:     42.,
		animate:  false,
		duration: 12,
		speed:    20.,
		force:    false,
	}
	lc.read()
	lc.cat()
	return lc
}

func (self *LolCat) read() {
	buff := bytes.Buffer{}
	buff.ReadFrom(self.input)
	var (
		err  error
		line []byte
		tab  []string
	)
	for err == nil {
		line, err = buff.ReadBytes('\n')
		if line != nil {
			tab = append(tab, string(line))
		}
	}
	self.tab = tab
}

func (self *LolCat) cat() {
	for _, v := range self.tab {
		self.num++
		self.printLine(v)
	}
}

func (self *LolCat) printLine(line string) {
	line = line[:len(line)-1]
	self.printLinePlain(line)
	self.output.WriteString("\n")
}

func (self *LolCat) printLinePlain(line string) {
	for i, v := range line {
		r, g, b := self.rainbow(
			self.freq,
			float64(self.num)+float64(i)/self.spread,
		)
		self.output.WriteString(self.wrap(self.ansi(r, g, b)))
		self.output.WriteString(string(v))
	}
}

func (self *LolCat) rainbow(freq, i float64) (r, g, b float64) {
	r = math.Sin(freq*i)*127 + 128
	g = math.Sin(freq*i+2*math.Pi/3)*127 + 128
	b = math.Sin(freq*i+4*math.Pi/3)*127 + 128
	return r, g, b
}

func (self *LolCat) ansi(r, g, b float64) string {
	grayPossible := true
	gray := true
	sep := 2.5

	for grayPossible {
		if r < sep || g < sep || b < sep {
			gray = r < sep && g < sep && b < sep
			grayPossible = false
		}

		sep += 42.5
	}

	var color int
	if gray {
		color = 232 + int((r+g+b)/33.)
	} else {
		color = 16 +
			int(6*r/256)*36 +
			int(6*g/256)*6 +
			int(6*b/256)*1
	}

	return fmt.Sprintf("38;5;%d", color)
}

func (self *LolCat) wrap(codes string) string {
	return fmt.Sprintf("\x1b[%sm", codes)
}
