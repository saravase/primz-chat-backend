package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/handler"
)

// Delete godoc
// @Summary      delete user detail based on user id
// @Description  Delete user detail based on user id
// @Tags         auth
// @Produce      json
// @Security ApiKeyAuth
// @Param        user_id   path     string  true  "User ID"
// @Success      200  {object}  handler.DeleteResponse
// @Failure      400  {object}  apperrors.Error
// @Failure      404  {object}  apperrors.Error
// @Failure      500  {object}  apperrors.Error
// @Router       /api/auth/user/{user_id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("user_id")
	ctx := c.Request.Context()
	_, err := h.UserService.Get(ctx, id)
	if err != nil {
		log.Printf("Failed to get user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	err = h.UserService.Delete(ctx, id)
	if err != nil {
		log.Printf("Failed to delete user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, &handler.DeleteResponse{
		IsDeleted: true,
	})
}
