package message

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/handler"
	"github.com/saravase/primz-chat-backend/model"
)

// Create godoc
// @Summary      create new message
// @Description  Create new message
// @Tags         message
// @Accept      json
// @Produce      json
// @Security ApiKeyAuth
// @Param message body MessageCreateReq true "Message Detail"
// @Success      200  {object}  handler.CreateResponse
// @Failure      415  {object}  apperrors.Error
// @Failure      400  {object}  apperrors.Error
// @Failure      409  {object}  apperrors.Error
// @Failure      500  {object}  apperrors.Error
// @Router       /api/message/ [post]
func (h *Handler) Create(c *gin.Context) {
	var (
		req MessageCreateReq
		msg model.Message
	)
	if ok := handler.BindData(c, &req); !ok {
		return
	}
	copier.Copy(&msg, &req)

	ctx := c.Request.Context()
	err := h.MessageService.Create(ctx, &msg)
	if err != nil {
		log.Printf("Failed to create message: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, &handler.CreateResponse{
		IsCreated: true,
	})
}
