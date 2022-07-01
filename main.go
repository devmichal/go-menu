package main

import (
	"github.com/devmichal/foods/cmd"
	"log"
)

func main() {
	app := cmd.Route()
	log.Fatal(app.Listen(":8010"))
}
