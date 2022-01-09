package user

import "minepin-backend/model"

type CreateReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type CreateRep struct {
	Username string `json:"username"`
}

type ListReq struct {
	Username string `json:"username"`
	Offset   int `json:"offset"`
	Limit    int `json:"limit"`
}

type ListRep struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}