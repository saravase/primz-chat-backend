package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/model"
)

// Signout godoc
// @Summary      signout chat application
// @Description  signout chat application
// @Tags         auth
// @Security ApiKeyAuth
// @Success      200  {string} string "user signed out successfully!"
// @Failure      401  {object}  apperrors.Error
// @Failure      500  {object}  apperrors.Error
// @Router       /api/auth/signout [post]
func (h *Handler) Signout(c *gin.Context) {
	user := c.MustGet("user")

	ctx := c.Request.Context()
	if err := h.TokenService.Signout(ctx, user.(*model.User).UserID); err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, "user signed out successfully!")
}
