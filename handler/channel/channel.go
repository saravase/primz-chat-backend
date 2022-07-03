package channel

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/apperrors"
)

// Channel godoc
// @Summary      get channel detail based on channel id
// @Description  Get channel detail based on channel id
// @Tags         channel
// @Produce      json
// @Security ApiKeyAuth
// @Param        channel_id   path     string  true  "Channel ID"
// @Success      200  {object}  model.Channel
// @Failure      400  {object}  apperrors.Error
// @Failure      404  {object}  apperrors.Error
// @Failure      500  {object}  apperrors.Error
// @Router       /api/channel/{channel_id} [get]
func (h *Handler) Channel(c *gin.Context) {
	id := c.Param("channel_id")
	ctx := c.Request.Context()
	org, err := h.ChannelService.Get(ctx, id)
	if err != nil {
		log.Printf("Failed to get channel: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, org)
}
