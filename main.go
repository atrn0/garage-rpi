package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	baseURL  = "http://go-rpi-gpio-api:8080"
	openPIN  = 16
	stopPIN  = 26
	closePIN = 21
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/open", func(c echo.Context) error {
		if err := push(stopPIN); err != nil {
			return err
		}

		time.Sleep(100 * time.Millisecond)

		if err := push(openPIN); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	})

	e.POST("/stop", func(c echo.Context) error {
		if err := push(stopPIN); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	})

	e.POST("/close", func(c echo.Context) error {
		if err := push(stopPIN); err != nil {
			return err
		}

		time.Sleep(100 * time.Millisecond)

		if err := push(closePIN); err != nil {
			return err
		}

		return c.NoContent(http.StatusOK)
	})

	e.Logger.Fatal(e.Start(":9000"))
}

func push(pin int) error {
	{
		res, err := http.Post(baseURL+"/pin/"+strconv.Itoa(pin)+"/0", "application/json", nil)
		if err != nil {
			return err
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		fmt.Println(string(b))
		res.Body.Close()
	}

	time.Sleep(100 * time.Millisecond)

	{
		res, err := http.Post(baseURL+"/pin/"+strconv.Itoa(pin)+"/1", "application/json", nil)
		if err != nil {
			return err
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		fmt.Println(string(b))
		res.Body.Close()
	}

	return nil
}
