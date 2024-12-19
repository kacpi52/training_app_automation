package user

type User struct {
	Id       string  `json:"id"`
	UserName *string `json:"userName"`
	LastName *string `json:"lastName"`
	NickName *string `json:"nickName"`
	Email    *string `json:"email"`
	Role     *string `json:"role"`
	Sub      *string `json:"sub"`
}

type UserData struct {
	Nickname string `json:"nickName"`
	Name     string `json:"name"`
	Sub      string `json:"sub"`
}