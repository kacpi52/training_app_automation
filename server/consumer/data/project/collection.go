package project

type Collection struct {
	Id          string  `json:"id"`
	UserId      string  `json:"userId"`
	IdLanguage  string  `json:"idLanguage"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	CreatedUp   *string `json:"createdUp"`
	UpdateUp    *string `json:"updateUp"`
}

type SearchProject struct {
	Id         string `json:"id"`
	IdLanguage string `json:"idLanguage"`
	Page       string `json:"page"`
}