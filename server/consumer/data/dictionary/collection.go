package dictionary

type Collection struct {
	Id        string `json:"id"`
	CreatedUp string `json:"createdUp"`
	UpdateUp  string `json:"updateUp"`
}

type MultiCollection struct {
	Id           string `json:"id"`
	DictionaryId string `json:"dictionaryId"`
	Key          string `json:"key"`
	Translation  string `json:"translation"`
}

type ResponseCollection struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Translation string `json:"translation"`
	CreatedUp   string `json:"createdUp"`
	UpdateUp    string `json:"updateUp"`
}