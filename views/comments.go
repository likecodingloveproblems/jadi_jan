package views

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/likecodingloveproblems/jadi_jan/models"
)

func GetCommentsView(c echo.Context) error {
	ctx := context.Background()
	comments := &models.Comments{}
	comments.Get(ctx)
	return c.JSON(http.StatusOK, comments.Comments)
}

func CreateCommentsView(c echo.Context) error {
	ctx := context.Background()
	comment := &models.Comment{}
	err := c.Bind(comment)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	comment.Create(ctx)
	return c.JSON(http.StatusCreated, comment)
}
