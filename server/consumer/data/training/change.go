package training

type Change struct {
	ID        string `json:"id"`
	PostId    string `json:"postId"`
	Type      string `json:"type"`
	Time      string `json:"time"`
	Kcal      int64  `json:"kcal"`
	CreatedUp string `json:"createdUp"`
	UpdateUp  string `json:"updateUp"`
}

