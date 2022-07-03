package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/apperrors"
)

// Users godoc
// @Summary      get all user details
// @Description  Get all user details
// @Tags         auth
// @Produce      json
// @Security ApiKeyAuth
// @Success      200  {object}  []model.User
// @Failure      500  {object}  apperrors.Error
// @Router       /api/auth/users [get]
func (h *Handler) Users(c *gin.Context) {
	ctx := c.Request.Context()
	orgs, err := h.UserService.GetUsers(ctx)
	if err != nil {
		log.Printf("Failed to get all users: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, orgs)
}
