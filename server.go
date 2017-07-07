package main

import (
	runner "code_runner/runner"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Code Runner!")
	})

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

	e.Logger.Fatal(e.Start(":8080"))
}
