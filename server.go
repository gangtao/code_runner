package main

import (
	runner "code_runner/runner"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
)

var (
	is_secure = flag.Bool("s", false, "use https")
)

const (
	CODE_BASE_DIR = "code"
	CODE_POST_FIX = "txt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func get_code(language *string, id *string) (error, string) {
	path := fmt.Sprintf("./%s/%s/%s.%s", CODE_BASE_DIR, *language, *id, CODE_POST_FIX)
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return err, "Failed to read file"
	}
	return nil, string(dat)
}

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

	e.GET("/run/:language/:id", func(c echo.Context) error {
		language := c.Param("language")
		id := c.Param("id")

		err, code := get_code(&language, &id)
		if err != nil {
			return c.String(http.StatusInternalServerError, code)
		}
		return c.String(http.StatusOK, code)
	})

	if *is_secure {
		e.Logger.Fatal(e.StartTLS(":8080", "cert.pem", "key.pem"))
	} else {
		e.Logger.Fatal(e.Start(":8080"))
	}

}
