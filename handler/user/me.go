package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/model"
)

// Me godoc
// @Summary      Show user detail
// @Description  get user detail based on user id
// @Tags         auth
// @Produce      json
// @Security ApiKeyAuth
// @Success      200  {object}  model.User
// @Failure      400  {object}  apperrors.Error
// @Failure      404  {object}  apperrors.Error
// @Failure      500  {object}  apperrors.Error
// @Router       /api/auth/me [get]
func (h *Handler) Me(c *gin.Context) {
	user, exists := c.Get("user")

	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})

		return
	}

	uid := user.(*model.User).UserID

	ctx := c.Request.Context()
	u, err := h.UserService.Get(ctx, uid)

	if err != nil {
		log.Printf("Unable to find user: %v\n%v", uid, err)
		e := apperrors.NewNotFound("user", uid)

		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	c.JSON(http.StatusOK, u)
}
