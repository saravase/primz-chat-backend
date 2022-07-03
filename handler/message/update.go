package message

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/handler"
)

// Update godoc
// @Summary      update message detail based on message id
// @Description  Update message detail based on message id
// @Tags         message
// @Accept      json
// @Produce      json
// @Security ApiKeyAuth
// @Param        msg_id   path     string  true  "Message ID"
// @Param message body MessageUpdateReq true "Message Detail"
// @Success      200  {object}  handler.UpdateResponse
// @Failure      415  {object}  apperrors.Error
// @Failure      400  {object}  apperrors.Error
// @Failure      404  {object}  apperrors.Error
// @Failure      500  {object}  apperrors.Error
// @Router       /api/message/{msg_id} [put]
func (h *Handler) Update(c *gin.Context) {
	var (
		req MessageUpdateReq
	)

	id := c.Param("msg_id")
	if ok := handler.BindData(c, &req); !ok {
		return
	}

	ctx := c.Request.Context()
	org, err := h.MessageService.Get(ctx, id)
	if err != nil {
		log.Printf("Failed to get message: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	copier.CopyWithOption(org, &req, copier.Option{IgnoreEmpty: true})

	err = h.MessageService.Update(ctx, id, org)
	if err != nil {
		log.Printf("Failed to update message: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, &handler.UpdateResponse{
		IsUpdated: true,
	})
}
