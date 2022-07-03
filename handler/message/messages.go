package message

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/apperrors"
)

// Messages godoc
// @Summary      get messages based on channel id
// @Description  Get messages based on channel id
// @Tags         message
// @Produce      json
// @Security ApiKeyAuth
// @Param        channel_id   path     string  true  "Channel ID"
// @Success      200  {object}  []model.Channel
// @Failure      500  {object}  apperrors.Error
// @Router       /api/messages/{channel_id} [get]
func (h *Handler) Messages(c *gin.Context) {
	id := c.Param("channel_id")
	ctx := c.Request.Context()
	channels, err := h.MessageService.GetByChennelID(ctx, id)
	if err != nil {
		log.Printf("Failed to get messages based on channel id: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, channels)
}
