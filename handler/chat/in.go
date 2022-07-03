package chat

type ChannelUser struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type ConnReq struct {
	Users        []ChannelUser `json:"users" validate:"required,gt=1,dive"`
	Name         string        `json:"name" validate:"omitempty,min=2,max=100"`
	GroupChannel *bool         `json:"group_channel" validate:"required"`
	ChannelOwner ChannelUser   `json:"channel_owner" validate:"required,dive"`
} // @name ConnReq
