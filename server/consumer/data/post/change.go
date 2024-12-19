package post

type Change struct {
	Id        string   `json:"id"`
	UserId    string   `json:"userId"`
	ProjectId string   `json:"projectId"`
	Day       *int64   `json:"day"`
	Weight    *float64 `json:"weight"`
	Kcal      *int64   `json:"kcal"`
	CreatedUp string   `json:"createdUp"`
	UpdateUp  *string  `json:"updateUp"`
}

type ChangePost struct {
	Day                      int64               `json:"day"`
	Weight                   float64             `json:"weight"`
	Kcal                     int64               `json:"kcal"`
	CollectionTraining       []OneTraining       `json:"collectionTraining"`
	CollectionTrainingChange []OneTrainingChange `json:"collectionTrainingChange"`
	RemoveIds                []string            `json:"removeIds"`
}

type OneTrainingChange struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Time string `json:"time"`
	Kcal int64  `json:"kcal"`
}
