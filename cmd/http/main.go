package main

import "github.com/Ceazzer/personal-website-server/config"

func main() {
	// Setup environment variables
	config.DotenvSetup()

	// Setup server
	e := config.ServerSetup()

	// Start server
	e.Logger.Fatal(e.Start(":8080"))

}
