package main

import (
	"github.com/labstack/echo/v4"
	"github.com/likecodingloveproblems/jadi_jan/views"
)

func main() {
	e := echo.New()
	e.GET("counter/", views.GetCounterView)
	e.POST("counter/", views.IncrCounterView)
	e.GET("comments/", views.GetCommentsView)
	e.POST("comments/", views.CreateCommentsView)
	e.Logger.Fatal(e.Start(":1323"))
}
