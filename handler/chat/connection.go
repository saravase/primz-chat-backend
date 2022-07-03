package chat

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/saravase/primz-chat-backend/apperrors"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     checkOrigin,
}

func checkOrigin(req *http.Request) bool {
	log.Printf("%s %s %s %s\n", req.Method, req.RequestURI, req.Host, req.Proto)
	return req.Method == http.MethodGet
}

func (h *Handler) UserConnection(c *gin.Context) {

	id := c.Param("user_id")

	// Create new websocket connection
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error while upgrade websocket connection. Reason: %v\n", err.Error())
		e := apperrors.NewInternal()
		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}
	defer conn.Close()

	ctx := c.Request.Context()

	h.ChatService.CreateChatConnection(ctx, conn, id)

}
