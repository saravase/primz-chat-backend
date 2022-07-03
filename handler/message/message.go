package message

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/apperrors"
)

// Message godoc
// @Summary      get message detail based on message id
// @Description  Get message detail based on message id
// @Tags         message
// @Produce      json
// @Security ApiKeyAuth
// @Param        msg_id   path     string  true  "Message ID"
// @Success      200  {object}  model.Message
// @Failure      400  {object}  apperrors.Error
// @Failure      404  {object}  apperrors.Error
// @Failure      500  {object}  apperrors.Error
// @Router       /api/message/{msg_id} [get]
func (h *Handler) Message(c *gin.Context) {
	id := c.Param("msg_id")
	ctx := c.Request.Context()
	org, err := h.MessageService.Get(ctx, id)
	if err != nil {
		log.Printf("Failed to get message: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, org)
}
