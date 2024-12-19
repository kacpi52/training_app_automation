package project

type Create struct {
	Id          string  `json:"id"`
	UserId      string  `json:"userId"`
	IdLanguage  string  `json:"idLanguage"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	CreatedUp   *string `json:"createdUp"`
	UpdateUp    *string `json:"updateUp"`
}