package chat

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/saravase/primz-chat-backend/apperrors"
	"github.com/saravase/primz-chat-backend/handler"
)

// GetChat godoc
// @Summary      get chat details based on channel id's
// @Description  Get chat details based on channel id's
// @Tags         chat
// @Produce      json
// @Security ApiKeyAuth
// @Param chat_req body ChatReq true "Chat Request"
// @Success      200  {object}  []handler.Chats
// @Failure      500  {object}  apperrors.Error
// @Router       /api/chat [post]
func (h *Handler) GetChat(c *gin.Context) {
	var (
		req ChatReq
	)
	if ok := handler.BindData(c, &req); !ok {
		return
	}
	ctx := c.Request.Context()
	channels, err := h.ChannelService.GetByUserIDs(ctx, req.ChannelIDs)
	if err != nil {
		log.Printf("Failed to get channels based on channel ids: %v\n", err.Error())
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	curUserID := req.UserID
	chats := make([]handler.Chat, 0)
	for _, channel := range channels {
		var chat handler.Chat
		chatMsgs := make([]handler.ChatMessage, 0)
		userIDs := make([]string, 0)
		users := make(map[string]handler.ChatUser)
		for _, user := range channel.Users {
			var chatUser handler.ChatUser
			if user.ID != curUserID {
				userIDs = append(userIDs, user.ID)
				u, err := h.UserService.Get(ctx, user.ID)
				if err != nil {
					log.Printf("Unable to find user: %v\n%v", user.ID, err)
					e := apperrors.NewNotFound("user", user.ID)
					c.JSON(e.Status(), gin.H{
						"error": e,
					})
					return
				}
				copier.Copy(&chatUser, u)
				users[user.ID] = chatUser
			}
		}
		msgs, err := h.MessageService.GetByChennelID(ctx, channel.ChannelID)
		if err != nil {
			log.Printf("Failed to get messages based on channel id: %v\n", err.Error())
			c.JSON(apperrors.Status(err), gin.H{
				"error": err,
			})
			return
		}
		copier.Copy(&chatMsgs, msgs)
		chat.UserIDs = userIDs
		chat.Users = users
		chat.Messages = chatMsgs
		chat.ChannelId = channel.ChannelID
		chats = append(chats, chat)
	}

	c.JSON(http.StatusOK, chats)
}
