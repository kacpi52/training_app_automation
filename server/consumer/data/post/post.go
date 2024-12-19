package post

type Post struct {
	Id        string  `json:"id"`
	UserId    string  `json:"userId"`
	ProjectId string  `json:"projectId"`
	Day       int64   `json:"day"`
	Weight    float64 `json:"weight"`
	Kcal      int64   `json:"kcal"`
	CreatedUp *string `json:"createdUp"`
	UpdateUp  *string `json:"updateUp"`
}

type CreatePost struct {
	Day                int64         `json:"day"`
	Weight             float64       `json:"weight"`
	Kcal               int64         `json:"kcal"`
	CollectionTraining []OneTraining `json:"collectionTraining"`
}

type OneTraining struct {
	Type string `json:"type"`
	Time string `json:"time"`
	Kcal int64  `json:"kcal"`
}