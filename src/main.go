package main

import (
	"fmt"
	"go.api-boilerplate/infrastructure/handlers"
	"log"
)

func main() {
	app := handlers.Handler()
	cfg := "80" //config.GetConfig()
	err := app.Listen(fmt.Sprintf(":%s", cfg))
	if err != nil {
		log.Panicf("Error:%s", err.Error())
	}
}
