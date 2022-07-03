package user

type UserUpdateReq struct {
	FirstName string `json:"first_name" validate:"omitempty,min=2,max=100"`
	LastName  string `json:"last_name" validate:"omitempty,min=2,max=100"`
	Email     string `json:"email" binding:"omitempty,email"`
	Role      string `json:"role" validate:"omitempty,eq=ADMIN|eq=USER"`
} // @name UserUpdateReq
