package test

import (
	"encoding/json"
	"fmt"
	params_data "myInternal/consumer/data"
	common_post "myInternal/consumer/handler/post/test"
	common_project "myInternal/consumer/handler/project/test"
	statistics_functions "myInternal/consumer/handler/statistics"
	common_training "myInternal/consumer/handler/training/test"
	env "myInternal/consumer/initializers"
	"testing"
)

type Training struct {
	Type string `json:"type"`
	Time string `json:"time"`
	Kcal int    `json:"kcal"`
}

type TrainingCollection struct {
	TrainingCollection []Training `json:"collectionTraining"`
}

func createDumpStatistics() (string, error) {
	createProjectId, err := common_project.CreateProject()
	if err != nil {
		return "", fmt.Errorf("create dump statistics project: %v", err)
	}

	dataBodiesPost := []struct {
		Day    int
		Weight float64
		Kcal   int
	}{
		{Day: 1, Weight: 98.8, Kcal: 2496},
		{Day: 2, Weight: 98.5, Kcal: 2535},
		{Day: 3, Weight: 98.3, Kcal: 2563},
		{Day: 4, Weight: 97.1, Kcal: 2485},
		{Day: 5, Weight: 96.6, Kcal: 3936},
		{Day: 6, Weight: 96.3, Kcal: 2429},
		{Day: 7, Weight: 96.3, Kcal: 2496},
		{Day: 8, Weight: 96.3, Kcal: 2500},
	}

	trainingCollections := [][]Training{
		{
			{Type: "bike", Time: "01:56:56", Kcal: 1341},
		},
		{
			{Type: "bike", Time: "2:11:04", Kcal: 1424},
		},
		{},
		{
			{Type: "bike", Time: "01:41:54", Kcal: 1093},
		},
		{
			{Type: "bike", Time: "01:56:56", Kcal: 1341},
		},
		{},
		{
			{Type: "bike", Time: "1:56:56", Kcal: 1341},
		},
		{},
	}

	for i, data := range dataBodiesPost {
		dataBody := fmt.Sprintf(`{
			"day": %d,
			"weight": %.2f,
			"kcal": %d
		}`, data.Day, data.Weight, data.Kcal)

		createPostId, err := common_post.CreatePost(dataBody, createProjectId)
		if err != nil {
			return "", fmt.Errorf("create dump statistics post: %v", err)
		}

		trainingData := trainingCollections[i]
		trainingCollection := TrainingCollection{TrainingCollection: trainingData}
		trainingCollectionJson, err := json.Marshal(trainingCollection)
		if err != nil {
			return "", fmt.Errorf("marshal training collection: %v", err)
		}

		err = common_training.CreateTraining(string(trainingCollectionJson), createPostId)
		if err != nil {
			return "", fmt.Errorf("create dump statistics training: %v", err)
		}
	}

	return createProjectId, nil
}

func TestCreateStatistics(t *testing.T) {

	projectId, err := createDumpStatistics()
	if err != nil {
		t.Fatalf("Error create dump statistics function: %v", err)
	}

	params := params_data.Params{
		Param: projectId,
	}

	env.LoadEnv("./.env")
	createStatisticOption, err := statistics_functions.CreateStatisticOption(params)
	if err != nil {
		t.Fatalf("Error create statistic options function: %v", err)
	}

	if len(createStatisticOption.Statistics) == 0{
		t.Fatalf("Error len function statistic options is 0")
	}

	createStatistics := statistics_functions.CollectionStatistics(createStatisticOption.Statistics)
	if len(createStatistics) == 0{
		t.Fatalf("Error len function statistic is 0")
	}

	if createStatistics[0].StartWeight != 98.8{
		t.Fatalf("Error StartWeight is not 98.8")
	}

	if createStatistics[0].EndWeight != 96.3{
		t.Fatalf("Error EndWeight is not 96.3")
	}

	if createStatistics[0].DownWeight != -2.5{
		t.Fatalf("Error DownWeight is not -2.5")
	}

	if createStatistics[0].SumKg != 681.9{
		t.Fatalf("Error SumKg is not 681.9")
	}

	if createStatistics[0].AvgKg != 97.41{
		t.Fatalf("Error SumKg is not 97.41")
	}

	if createStatistics[0].SumKcal != 18940{
		t.Fatalf("Error SumKg is not 18940")
	}
	
	if createStatistics[0].Training[0].Data[0].Type != "bike"{
		t.Fatalf("Error Training Type is not bike")
	}

	if createStatistics[0].Training[0].Data[0].Currecnt != 5{
		t.Fatalf("Error Training Currecnt is not 5")
	}

	if createStatistics[0].Training[0].Data[0].SumKcal != 6540{
		t.Fatalf("Error Training SumKcal is not 6540")
	}

	if createStatistics[0].Training[0].Data[0].Time != "9:43:46"{
		t.Fatalf("Error Training Time is not 9:43:46")
	}
}