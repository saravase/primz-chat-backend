package channel

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/handler"
)

// Update godoc
// @Summary      update channel detail based on channel id
// @Description  Update channel detail based on channel id
// @Tags         channel
// @Accept      json
// @Produce      json
// @Security ApiKeyAuth
// @Param        channel_id   path     string  true  "Channel ID"
// @Param channel body ChannelUpdateReq true "Channel Detail"
// @Success      200  {object}  handler.UpdateResponse
// @Failure      415  {object}  apperrors.Error
// @Failure      400  {object}  apperrors.Error
// @Failure      404  {object}  apperrors.Error
// @Failure      500  {object}  apperrors.Error
// @Router       /api/channel/{channel_id} [put]
func (h *Handler) Update(c *gin.Context) {
	var (
		req ChannelUpdateReq
	)

	id := c.Param("channel_id")
	if ok := handler.BindData(c, &req); !ok {
		return
	}

	ctx := c.Request.Context()
	org, err := h.ChannelService.Get(ctx, id)
	if err != nil {
		log.Printf("Failed to get channel: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	copier.CopyWithOption(org, &req, copier.Option{IgnoreEmpty: true})

	err = h.ChannelService.Update(ctx, id, org)
	if err != nil {
		log.Printf("Failed to update channel: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, &handler.UpdateResponse{
		IsUpdated: true,
	})
}
