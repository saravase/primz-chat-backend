package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/handler"
	"github.com/saravase/primz-chat-backend/model"
)

type signinReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
} //@name signinReq

// Signin godoc
// @Summary      signin chat application
// @Description  signin chat application
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        message body signinReq true "Signin Payload"
// @Success      200  {object}  model.TokenPair
// @Failure      500  {object}  apperrors.Error
// @Failure      415  {object}  apperrors.Error
// @Failure      400  {object}  apperrors.Error
// @Failure      404  {object}  apperrors.Error
// @Router       /api/auth/signin [post]
func (h *Handler) Signin(c *gin.Context) {
	var req signinReq

	if ok := handler.BindData(c, &req); !ok {
		return
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}

	ctx := c.Request.Context()
	err := h.UserService.Signin(ctx, u)

	if err != nil {
		log.Printf("Failed to sign in user: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	tokens, err := h.TokenService.NewPairFromUser(ctx, u, "")

	if err != nil {
		log.Printf("Failed to create tokens for user: %v\n", err.Error())

		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, tokens)
}
