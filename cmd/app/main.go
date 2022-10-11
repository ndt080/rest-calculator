package main

import (
	"log"
	s "rest-calculator/internal"
)

// @title          REST CALCULATOR Swagger API
// @version        1.0
// @description    Swagger API for Golang Project REST CALCULATOR.
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	server := &s.Server{}
	err := server.Run()

	if err != nil {
		log.Fatalf("Error while running server %s", err.Error())
	}
}
