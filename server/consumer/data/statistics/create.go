package statistics

type Create struct {
	Week        int               `json:"week"`
	StartWeight float64           `json:"startWeight"`
	EndWeight   float64           `json:"endWeight"`
	DownWeight  float64           `json:"downWeight"`
	SumKg       float64           `json:"sumKg"`
	AvgKg       float64           `json:"avgKg"`
	SumKcal     float64           `json:"sumKcal"`
	Training    []OneTrainingWeek `json:"training"`
	SumTime     string            `json:"sumTime"`
	CreatedUp   string            `json:"createdUp"`
	UpdateUp    string            `json:"updateUp"`
}

type OneTraining struct {
	PostId string `json:"postId"`
	Type   string `json:"type"`
	Time   string `json:"time"`
	Kcal   int64  `json:"kcal"`
}

type Statistics struct {
	Day                int64         `json:"day"`
	Weight             float64       `json:"weight"`
	Kcal               int64         `json:"kcal"`
	TrainingCollection []OneTraining `json:"trainingCollection"`
}

type CollectionTraining struct {
	Data []OneTrainingWeek `json:"data"`
}

type OneTrainingWeek struct {
	Data []TrainingsWeek `json:"data"`
}

type TrainingsWeek struct {
	Type     string `json:"type"`
	Currecnt int    `json:"currecnt"`
	SumKcal  int    `json:"sumKcal"`
	Time     string `json:"time"`
}