package main

import (
	v1 "bingo-example/ctrl/v1"
	"bingo-example/lib/configuration"
	"github.com/xylong/bingo"
)

func main() {
	bingo.Init().
		Inject(configuration.NewBoot(), configuration.NewService()).
		Mount("v1", v1.Controller...)().
		Lunch()
}
