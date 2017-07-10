package main

import (
	runner "code_runner/runner"
	"encoding/json"
	"flag"
	"github.com/labstack/echo"
	"net/http"
)

var (
	is_secure = flag.Bool("s", false, "use https")
)

func main() {
	flag.Parse()

	e := echo.New()

	e.Static("/", "static")

	e.POST("/run", func(c echo.Context) error {
		payload := &runner.Payload{}
		if err := c.Bind(payload); err != nil {
			return c.String(http.StatusInternalServerError, "Invalid Payload!")
		}
		result := runner.Run(payload)

		result_str, err := json.Marshal(result)

		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to marshal running result!")
		}

		return c.String(http.StatusOK, string(result_str))
	})

	if *is_secure {
		e.Logger.Fatal(e.StartTLS(":8080", "cert.pem", "key.pem"))
	} else {
		e.Logger.Fatal(e.Start(":8080"))
	}

}
