package message

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/handler"
)

// Delete godoc
// @Summary      delete message detail based on msg id
// @Description  Delete message detail based on msg id
// @Tags         message
// @Produce      json
// @Security ApiKeyAuth
// @Param        msg_id   path     string  true  "Message ID"
// @Success      200  {object}  handler.DeleteResponse
// @Failure      400  {object}  apperrors.Error
// @Failure      404  {object}  apperrors.Error
// @Failure      500  {object}  apperrors.Error
// @Router       /api/message/{msg_id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("msg_id")
	ctx := c.Request.Context()
	_, err := h.MessageService.Get(ctx, id)
	if err != nil {
		log.Printf("Failed to get message: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	err = h.MessageService.Delete(ctx, id)
	if err != nil {
		log.Printf("Failed to delete message: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, &handler.DeleteResponse{
		IsDeleted: true,
	})
}
