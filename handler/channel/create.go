package channel

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/handler"
	"github.com/saravase/primz-chat-backend/model"
)

// Create godoc
// @Summary      create new channel
// @Description  Create new channel
// @Tags         channel
// @Accept      json
// @Produce      json
// @Security ApiKeyAuth
// @Param channel body ChannelCreateReq true "Channel Detail"
// @Success      200  {object}  handler.CreateResponse
// @Failure      415  {object}  apperrors.Error
// @Failure      400  {object}  apperrors.Error
// @Failure      409  {object}  apperrors.Error
// @Failure      500  {object}  apperrors.Error
// @Router       /api/channel/ [post]
func (h *Handler) Create(c *gin.Context) {
	var (
		req     ChannelCreateReq
		channel model.Channel
		apperr  *apperrors.Error
	)
	if ok := handler.BindData(c, &req); !ok {
		return
	}
	copier.Copy(&channel, &req)

	ctx := c.Request.Context()
	tmpChannel, err := h.ChannelService.GetByUsers(ctx, &channel.Users)
	if err != nil && !errors.As(err, &apperr) {
		log.Printf("Failed to get channel: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	if tmpChannel != nil {
		e := apperrors.NewConflict("users", fmt.Sprintf("%#v\n", channel.Users))
		log.Printf("Failed to create channel: %v\n", e.Error())
		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	err = h.ChannelService.Create(ctx, &channel)
	if err != nil {
		log.Printf("Failed to create channel: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	for _, user := range channel.Users {
		u, err := h.UserService.Get(ctx, user.ID)
		if err != nil {
			log.Printf("Unable to find user: %v\n%v", user.ID, err)
			e := apperrors.NewNotFound("user", user.ID)
			c.JSON(e.Status(), gin.H{
				"error": e,
			})
			return
		}
		if channel.GroupChannel {
			if u.PubChannelIds != nil {
				u.PubChannelIds = append(u.PubChannelIds, channel.ChannelID)
			} else {
				u.PubChannelIds = []string{channel.ChannelID}
			}
		} else {
			if u.PrivChannelIds != nil {
				u.PrivChannelIds = append(u.PrivChannelIds, channel.ChannelID)
			} else {
				u.PrivChannelIds = []string{channel.ChannelID}
			}
		}
		err = h.UserService.Update(ctx, user.ID, u)
		if err != nil {
			log.Printf("Failed to update user: %v\n", err.Error())
			c.JSON(apperrors.Status(err), gin.H{
				"error": err,
			})
			return
		}
	}

	c.JSON(http.StatusOK, &handler.CreateResponse{
		ID: channel.ChannelID,
	})
}
