package views

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/likecodingloveproblems/jadi_jan/models"
)

func GetCounterView(c echo.Context) error {
	counter := &models.Counter{}
	ctx := context.Background()
	counter.Get(ctx)
	return c.JSON(http.StatusOK, counter)
}

func IncrCounterView(c echo.Context) error {
	counter := &models.Counter{}
	ctx := context.Background()
	counter.Incr(ctx)
	return c.JSON(http.StatusOK, counter)
}
