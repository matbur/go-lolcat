package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	n := rand.Int() % 256
	CreateLolCat(n)
	fmt.Println(n)
}
