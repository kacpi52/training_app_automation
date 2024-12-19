package statistics

type Collection struct {
	Week         int     `json:"week"`
	StartWeight  float64 `json:"startWeight"`
	EndWeight    float64 `json:"endWeight"`
	DownWeight   float64 `json:"downWeight"`
	SumKg        float64 `json:"sumKg"`
	AvgKg        float64 `json:"avgKg"`
	SumKcal      int     `json:"sumKcal"`
	TypeTraining string  `json:"typeTraining"`
	SumTime      string  `json:"sumTime"`
	CreatedUp    string  `json:"createdUp"`
	UpdateUp     string  `json:"updateUp"`
}