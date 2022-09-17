package main

import (
	v1 "bingo-example/api/v1"
	"bingo-example/lib/core"
	"github.com/xylong/bingo"
)

func main() {
	bingo.Init("conf", "app").
		Inject(core.NewClient(), core.NewService()).
		Mount("v1", v1.Controller...)().
		Lunch()
}
