package project

type Delete struct {
	Id        string  `json:"id"`
	UserId    string  `json:"userId"`
	CreatedUp *string `json:"createdUp"`
	UpdateUp  *string `json:"updateUp"`
}