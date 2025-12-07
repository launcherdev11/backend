package main

import (
	"16launcher/handlers"
	"16launcher/routes"
	"16launcher/utils"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type Message struct {
	Status string `json:"message"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to load .env: %v", err)
	}
	e := echo.New()
	e.GET("/ping", func(context echo.Context) error {
		return context.JSON(200, Message{"OK"})
	})

	minio := utils.RegisterMinio()

	db := utils.RegisterPostgres()
	handler := &handlers.Handler{DB: db, MinIOClient: minio}
	routes.RegisterMojangAuthRoutes(e, handler)
	fmt.Print("ОКАК")

	e.Start(":18205")
}
