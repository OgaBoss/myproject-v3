package main

import (
	"github.com/OgaBoss/myproject-v3/http/route"
	_ "github.com/OgaBoss/myproject-v3/http/validation"
	"os"

	"github.com/System-Glitch/goyave/v3"
	// Import the appropriate GORM dialect for the database you're using.
	_ "github.com/System-Glitch/goyave/v3/database/dialect/mysql"
)


func main() {
	// This is the entry point of your application.
	if err := goyave.Start(route.Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}
}
