package main

import (
	"context"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/aexrun/htmx-go-templ-echo-pg/internal/database"
	"github.com/aexrun/htmx-go-templ-echo-pg/internal/handlers"
)

func main() {
	ctx := context.Background()

	pool, err := database.Connect(ctx)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer pool.Close()

	if err := database.Migrate(ctx, pool); err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}

	h := handlers.New(pool)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", h.Home)
	e.POST("/todos", h.CreateTodo)
	e.PATCH("/todos/:id/toggle", h.ToggleTodo)
	e.DELETE("/todos/:id", h.DeleteTodo)
	e.GET("/health", h.HealthCheck)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
