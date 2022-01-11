package user

type CreateReq struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type CreateRep struct {
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type ListReq struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

//type ListRep struct {
//	TotalCount uint64            `json:"totalCount"`
//	UserList   []*model.UserInfo `json:"userList"`
//}
