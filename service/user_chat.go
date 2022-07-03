package service

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/saravase/primz-chat-backend/model"
)

type UserChat struct {
	Channels   *ChatChannel
	UserID     string
	Connection *websocket.Conn
}

func NewUserChat(channels *ChatChannel, userID string, conn *websocket.Conn) *UserChat {
	return &UserChat{
		Channels:   channels,
		UserID:     userID,
		Connection: conn,
	}
}

func (u *UserChat) OnlineListen() {
	for {
		if _, message, err := u.Connection.ReadMessage(); err != nil {
			log.Printf("Error while reading messge from connection. Reason: %v\n", err.Error())
			break
		} else {
			msg := new(model.Message)
			if err := json.Unmarshal(message, msg); err != nil {
				log.Printf("Error while unmarshaling received msg. Reason: %v\n", err.Error())
			} else {
				u.Channels.messageChannel <- msg
			}
		}
	}

	u.Channels.leaveChannel <- u
}

func (u *UserChat) SendMessageToClient(message *model.Message) error {
	message.MsgID = RandomID()
	if data, err := json.Marshal(message); err != nil {
		return fmt.Errorf("error, While marshaling message. Reason : %v", err.Error())
	} else {
		err := u.Connection.WriteMessage(websocket.TextMessage, data)
		log.Printf("Message send from %s to %s \n", message.UserID, message.ChannelID)
		return err
	}
}
