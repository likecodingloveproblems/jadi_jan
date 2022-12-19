package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/likecodingloveproblems/jadi_jan/views"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("counter/", views.GetCounterView)
	e.POST("counter/", views.IncrCounterView)
	e.GET("comments/", views.GetCommentsView)
	e.POST("comments/", views.CreateCommentsView)
	e.GET("ws/", views.WebSocket)
	e.Logger.Fatal(e.Start(":1323"))
}
