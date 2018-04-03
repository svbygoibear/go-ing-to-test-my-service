package cmd

import (
	"github.com/labstack/echo"
	"github.com/svbygoibear/go-ing-to-test-my-service/models"
	"github.com/svbygoibear/go-ing-to-test-my-service/config"
)

func Execute() {
	run()
}

func run() {
	servconf, err := config.Load("config.json")

	if(err == nil) {
		e := echo.New()
		e.GET("/user/:id", models.GetUser)

		e.Start(string(servconf.Port))
	}
}
