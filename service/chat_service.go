package service

import (
	"context"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/saravase/primz-chat-backend/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageChannel chan *model.Message
type UserChannel chan *UserChat

type ChatChannel struct {
	messageChannel MessageChannel
	leaveChannel   UserChannel
}

type chatService struct {
	Users             map[string]*UserChat
	JoinChannel       UserChannel
	Channels          *ChatChannel
	UserRepository    model.UserRepository
	OrgRepository     model.OrgRepository
	ChennelRepository model.ChennelRepository
	MessageRepository model.MessageRepository
	ChatRepository    model.ChatRepository
}

type ChatConfig struct {
	UserRepository    model.UserRepository
	OrgRepository     model.OrgRepository
	ChennelRepository model.ChennelRepository
	MessageRepository model.MessageRepository
}

func NewChatService(c *ChatConfig) model.ChatService {
	return &chatService{
		Users:       make(map[string]*UserChat),
		JoinChannel: make(UserChannel),
		Channels: &ChatChannel{
			messageChannel: make(MessageChannel),
			leaveChannel:   make(UserChannel),
		},
		UserRepository:    c.UserRepository,
		OrgRepository:     c.OrgRepository,
		ChennelRepository: c.ChennelRepository,
		MessageRepository: c.MessageRepository,
	}
}

func (s *chatService) CreateChatConnection(ctx context.Context, conn *websocket.Conn, uid string) {
	u := NewUserChat(s.Channels, uid, conn)
	s.JoinChannel <- u
	u.OnlineListen()
}

func (s *chatService) UsersChatManager() {
	for {
		select {
		case userChat := <-s.JoinChannel:
			s.AddUser(userChat)
		case message := <-s.Channels.messageChannel:
			// Insert current message in DB
			message.ID = primitive.NewObjectID()
			message.MsgID = RandomID()
			message.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
			message.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
			s.MessageRepository.Create(context.TODO(), message)
			s.SendMessage(message)
		case userChat := <-s.Channels.leaveChannel:
			s.LeaveChat(userChat.UserID)
		}
	}
}

func (s *chatService) AddUser(userChat *UserChat) {
	if _, ok := s.Users[userChat.UserID]; !ok {
		s.Users[userChat.UserID] = userChat
		log.Printf("Added user id : %v\n", userChat.UserID)
	}
}

func (s *chatService) SendMessage(message *model.Message) {
	id := message.ChannelID
	// Fetch channel details from DB
	channel, err := s.ChennelRepository.FindByID(context.TODO(), id)
	if err != nil {
		log.Printf("Channel not found: %s\n", id)
	}

	// Send message to all the users within a channel
	for _, user_ := range channel.Users {
		if user_.ID != message.UserID {
			if user, ok := s.Users[user_.ID]; ok {
				if err := user.SendMessageToClient(message); err != nil {
					log.Printf("Error, While send messge to %s. Reason: %v\n", user_.ID, err.Error())
				}
			}
		}
	}

}

func (s *chatService) LeaveChat(uid string) {
	if user, ok := s.Users[uid]; ok {
		defer user.Connection.Close()
		delete(s.Users, uid)
		log.Printf("User: %s removed from chat.\n", uid)
	}
}
