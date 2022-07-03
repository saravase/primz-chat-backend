package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/handler"
)

// Update godoc
// @Summary      update user detail based on user id
// @Description  Update user detail based on user id
// @Tags         auth
// @Accept      json
// @Produce      json
// @Security ApiKeyAuth
// @Param        user_id   path     string  true  "User ID"
// @Param user body UserUpdateReq true "User Detail"
// @Success      200  {object}  handler.UpdateResponse
// @Failure      415  {object}  apperrors.Error
// @Failure      400  {object}  apperrors.Error
// @Failure      404  {object}  apperrors.Error
// @Failure      500  {object}  apperrors.Error
// @Router       /api/auth/user/{user_id} [put]
func (h *Handler) Update(c *gin.Context) {
	var (
		req UserUpdateReq
	)

	id := c.Param("user_id")
	if ok := handler.BindData(c, &req); !ok {
		return
	}

	ctx := c.Request.Context()
	user, err := h.UserService.Get(ctx, id)
	if err != nil {
		log.Printf("Failed to get user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	copier.CopyWithOption(user, &req, copier.Option{IgnoreEmpty: true})

	err = h.UserService.Update(ctx, id, user)
	if err != nil {
		log.Printf("Failed to update user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, &handler.UpdateResponse{
		IsUpdated: true,
	})
}
