package routes

import (
	"16launcher/handlers"

	"github.com/labstack/echo/v4"
)

type Message struct {
	Status string `json:"message"`
}

func RegisterMojangAuthRoutes(e *echo.Echo, h *handlers.Handler) {
	group := e.Group("/")
	group.GET("check", func(context echo.Context) error {
		return context.JSON(200, Message{"green"})
	})
	group.POST("session/minecraft/profile/:uuid")
}
