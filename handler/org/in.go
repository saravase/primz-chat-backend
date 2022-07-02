package org

type Department struct {
	ID     string  `json:"id"`
	Name   string  `json:"name" validate:"required,min=2,max=100"`
	Groups []Group `json:"groups" validate:"omitempty,dive"`
}

type Group struct {
	ID   string `json:"id"`
	Name string `json:"name" validate:"required,min=2,max=100"`
}

type OrgCreateReq struct {
	Name        string       `json:"name" validate:"required,min=2,max=100"`
	Email       string       `json:"email" validate:"required,email"`
	Website     string       `json:"website" validate:"omitempty,url"`
	AvatarURL   string       `json:"avatar_url" validate:"omitempty,url"`
	Departments []Department `json:"departments" validate:"omitempty,dive"`
} // @name OrgCreateReq

type OrgUpdateReq struct {
	Name      string `json:"name" validate:"omitempty,min=2,max=100"`
	Email     string `json:"email" validate:"omitempty,email"`
	Website   string `json:"website" validate:"omitempty,url"`
	AvatarURL string `json:"avatar_url" validate:"omitempty,url"`
} // @name OrgUpdateReq
