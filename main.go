package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		machineAccount := MachineAccount{
			Name:       "test",
			Email:      "foo@bar.com",
			IP:         "127.0.0.1",
			OsUser:     "pi",
			OsPassword: "raspberry",
		}

		return c.JSON(http.StatusOK, machineAccount)
	})

	e.Run(standard.New(":1323"))
}

type MachineAccount struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	IP         string `json:"ip"`
	OsUser     string `json:"osUser"`
	OsPassword string `json:"osPassword"`
}
