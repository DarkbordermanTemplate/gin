package main

import (
	"practice/connection"
	"practice/endpoints"
	"practice/lib"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger

// @contact.name darkborderman
// @contact.url https://darkborderman.github.io/

// @license.name BSD 2-Clause License
// @license.url https://github.com/Darkborderman/gin-practice/blob/master/LICENSE

// @host localhost:8080
// schemes http
func main() {
	connection.Setup()
	router := endpoints.Create_router()
	router.Run(lib.CONFIG.Address + ":" + lib.CONFIG.Port)
}
