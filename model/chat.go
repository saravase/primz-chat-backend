package model

type Chat struct {
	UserID          string    `json:"user_id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Email           string    `json:"email"`
	Role            string    `json:"role"`
	Info            Info      `json:"info"`
	PrivChennalList []Channel `json:"priv_chennal_list"`
	PubChannelList  []Channel `json:"pub_channel_list"`
}
type Info struct {
	OrgID          string   `json:"org_id"`
	DeptID         string   `json:"dept_id"`
	GroupID        string   `json:"group_id"`
	PrivChannelIds []string `json:"priv_channel_ids"`
	PubChannelIds  []string `json:"pub_channel_ids"`
	ActiveStatus   bool     `json:"active_status"`
	AccessRole     string   `json:"access_role"`
}
