package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/handler"
)

type tokensReq struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
} //@name tokensReq

// Tokens godoc
// @Summary      Generate tokens pair
// @Description  Generate idToken and refreshToken
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        message   body  tokensReq  true  "Tokens Payload"
// @Success      200  {object}  model.TokenPair
// @Failure      500  {object}  apperrors.Error
// @Failure      415  {object}  apperrors.Error
// @Failure      400  {object}  apperrors.Error
// @Router       /api/auth/tokens [post]
func (h *Handler) Tokens(c *gin.Context) {
	var req tokensReq

	if ok := handler.BindData(c, &req); !ok {
		return
	}

	ctx := c.Request.Context()

	refreshToken, err := h.TokenService.ValidateRefreshToken(req.RefreshToken)

	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	u, err := h.UserService.Get(ctx, refreshToken.UID)

	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	tokens, err := h.TokenService.NewPairFromUser(ctx, u, refreshToken.ID)

	if err != nil {
		log.Printf("Failed to create tokens for user: %+v. Error: %v\n", u, err.Error())

		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, tokens)
}
