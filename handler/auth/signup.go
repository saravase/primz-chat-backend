package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/handler"
	"github.com/saravase/primz-chat-backend/model"
)

type signupReq struct {
	FirstName string `json:"first_name" validate:"required, min=2, max=100"`
	LastName  string `json:"last_name" validate:"required, min=2, max=100"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,gte=6,lte=30"`
	Role      string `json:"role" validate:"required, eq=ADMIN|eq=USER"`
}

func (h *Handler) Signup(c *gin.Context) {
	var req signupReq
	if ok := handler.BindData(c, &req); !ok {
		return
	}

	u := &model.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Role:      req.Role,
	}

	ctx := c.Request.Context()
	err := h.UserService.Signup(ctx, u)

	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
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

	c.JSON(http.StatusCreated, gin.H{
		"tokens": tokens,
	})
}
