package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/jessevdk/go-flags"
)

type Opts struct {
	Spread   float64 `short:"p" long:"spread"   default:"3."  value-name:"SPREAD"   description:"Rainbow spread"`
	Freq     float64 `short:"F" long:"freq"     default:".1"  value-name:"FREQ"     description:"Rainbow frequency"`
	Seed     int     `short:"S" long:"seed"     default:"0"   value-name:"SEED"     description:"Rainbow seed"`
	Animate  bool    `short:"a" long:"animate"                value-name:"ANIMATE"  description:"Enable psychedelics"`
	Duration int     `short:"d" long:"duration" default:"12"  value-name:"DURATION" description:"Animation duration"`
	Speed    float64 `short:"s" long:"speed"    default:"20." value-name:"SPEED"    description:"Animation speed"`
	Force    bool    `short:"f" long:"force"                  value-name:"FORCE"    description:"Force colour even when stdout is not a tty"`
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	var opts Opts
	args, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		os.Exit(1)
	}

	log.Println(args)
	CreateLolCat(opts)
}
