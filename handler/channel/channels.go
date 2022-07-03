package channel

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/handler"
)

// Channels godoc
// @Summary      get channels details based on channel id's
// @Description  Get channels details based on channel id's
// @Tags         channel
// @Produce      json
// @Security ApiKeyAuth
// @Param channel_id's body ChannelIDs true "Channel ID's"
// @Success      200  {object}  []model.Channel
// @Failure      500  {object}  apperrors.Error
// @Router       /api/channels [post]
func (h *Handler) Channels(c *gin.Context) {
	var (
		req ChannelIDs
	)
	if ok := handler.BindData(c, &req); !ok {
		return
	}
	ctx := c.Request.Context()
	channels, err := h.ChannelService.GetByUserIDs(ctx, req.ChannelIDs)
	if err != nil {
		log.Printf("Failed to get channels based on channel ids: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, channels)
}
