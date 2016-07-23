package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"

	"github.com/justphil/uph-go-hacking/models"
)

func main() {
	e := echo.New()

	e.GET("/machine-accounts", getMachineAccounts)
	e.GET("/machine-accounts/:name", getMachineAccountByName)
	e.POST("/machine-accounts", createMachineAccount)
	e.PUT("/machine-accounts/:name", updateMachineAccountByName)
	e.DELETE("/machine-accounts/:name", deleteMachineAccountByName)

	e.Run(standard.New(":1323"))
}

func getMachineAccounts(c echo.Context) error {
	machineAccount := models.MachineAccount{
		Name:       "test",
		Email:      "foo@bar.com",
		IP:         "127.0.0.1",
		OsUser:     "pi",
		OsPassword: "raspberry",
	}

	return c.JSON(http.StatusOK, machineAccount)
}

func getMachineAccountByName(c echo.Context) error {
	return nil
}

func createMachineAccount(c echo.Context) error {
	return nil
}

func updateMachineAccountByName(c echo.Context) error {
	return nil
}

func deleteMachineAccountByName(c echo.Context) error {
	return nil
}
