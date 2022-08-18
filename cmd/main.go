package main

import (
	v1 "bingo-example/ctrl/v1"
	"bingo-example/lib"
	"github.com/xylong/bingo"
)

func main() {
	bingo.Init().
		Inject(lib.NewDB()).
		Mount("v1", v1.Controller...)().
		Lunch()
}
