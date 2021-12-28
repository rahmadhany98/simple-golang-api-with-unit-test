package main

import (
	"github.com/rahmadhany98/simple-golang-api-with-unit-test/config"
	"github.com/rahmadhany98/simple-golang-api-with-unit-test/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
