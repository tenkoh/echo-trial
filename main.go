package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

var users = []string{"John, Mike"}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func showUsers(c echo.Context) error {
	return c.String(http.StatusOK, strings.Join(users, ", "))
}

func saveUser(c echo.Context) error {
	u := c.FormValue("name")
	if u == "" {
		return c.String(http.StatusBadRequest, "name must be non null")
	}
	users = append(users, u)
	return c.String(http.StatusOK, "name: "+u)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo")
	})
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.GET("/users", showUsers)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)
	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	// for Windows: sigterm: kill, sigint: ctrl+C
	// Notice: for MacOS, control[^] + C
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
