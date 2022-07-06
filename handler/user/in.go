package user

type UserUpdateReq struct {
	FirstName    string `json:"first_name" validate:"omitempty,min=2,max=100"`
	LastName     string `json:"last_name" validate:"omitempty,min=2,max=100"`
	Email        string `json:"email" binding:"omitempty,email"`
	Role         string `json:"role" validate:"omitempty,eq=ADMIN|eq=USER"`
	AvatarURL    string `json:"avatar_url" validate:"omitempty,url"`
	OrgID        string `json:"org_id" validate:"omitempty,len=8"`
	DeptID       string `json:"dept_id" validate:"omitempty,len=8"`
	GroupID      string `json:"group_id" validate:"omitempty,len=8"`
	ActiveStatus bool   `json:"active_status" validate:"omitempty"`
} // @name UserUpdateReq
