package views

import (
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/likecodingloveproblems/jadi_jan/models"
	"nhooyr.io/websocket"
)

func WebSocket(c echo.Context) error {

	socket, err := websocket.Accept(c.Response(), c.Request(), &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Println(err)
		return err
	}
	counter := models.UpdatedCounter
	for {
		if models.UpdatedCounter != counter {
			socket.Write(
				c.Request().Context(),
				websocket.MessageText,
				[]byte(strconv.Itoa(int(models.UpdatedCounter))),
			)
			counter = models.UpdatedCounter
		}
	}

	socket.Close(websocket.StatusNormalClosure, "WebSocket closed")

	return nil
}

// func receiveCounterUpdate(ctx context.Context, socket *websocket.Conn, counter int) {
// }
