package test

import (
	statistics_data "myInternal/consumer/data/statistics"
	helpers_statictics "myInternal/consumer/handler/statistics/helpers"
	"testing"
)

func TestSumTime(t *testing.T) {

	statistics := []statistics_data.Statistics{
		{Day: 1, Weight: 55.55, Kcal: 2300, TrainingCollection: []statistics_data.OneTraining{{Type: "gym", Time: "2019-10-12T1:10:10Z", Kcal: 1200}}},
		{Day: 2, Weight: 35.35, Kcal: 2500, TrainingCollection: []statistics_data.OneTraining{{Type: "bike", Time: "2019-10-12T00:59:40Z", Kcal: 2700}}},
		{Day: 3, Weight: 25.35, Kcal: 2500, TrainingCollection: []statistics_data.OneTraining{{Type: "bike", Time: "2019-10-12T00:55:40Z", Kcal: 1700}}},
		{Day: 4, Weight: 15.35, Kcal: 2500, TrainingCollection: []statistics_data.OneTraining{{Type: "bike", Time: "2019-10-12T00:45:40Z", Kcal: 1700}}},
		{Day: 5, Weight: 95.35, Kcal: 2500, TrainingCollection: []statistics_data.OneTraining{{Type: "bike", Time: "2019-10-12T00:35:40Z", Kcal: 400}}},
		{Day: 6, Weight: 95.35, Kcal: 2500, TrainingCollection: []statistics_data.OneTraining{{Type: "bike", Time: "2019-10-12T00:25:40Z", Kcal: 300}}},
		{Day: 7, Weight: 95.35, Kcal: 2500, TrainingCollection: []statistics_data.OneTraining{{Type: "bike", Time: "2019-10-12T00:15:40Z", Kcal: 400}}},
		{},
	}

	sumCollection := helpers_statictics.SumTime(statistics)
	if sumCollection[0] != "5:08:10"{
		t.Errorf("error time is not 5:08:10")
	}
}

func TestAddTime(t *testing.T){
	var currentTime string = "00:00:00"
	var result string

	var times []string
	times = append(times, "2019-10-12T1:10:10Z")
	times = append(times, "2019-10-12T00:59:40Z")

	for _, time := range times{
		result, _ = helpers_statictics.AddTime(currentTime, time)
		currentTime = result
	}

	if result != "2:09:50"{
		t.Errorf("error return from function AddTime is not 2:09:50")
	}
}

func TestSumValue(t *testing.T){
	statistics := []statistics_data.Statistics{
		{Day: 1, Weight: 55.55, Kcal: 2300, TrainingCollection: []statistics_data.OneTraining{}},
		{Day: 2, Weight: 35.35, Kcal: 2500, TrainingCollection: []statistics_data.OneTraining{}},
		{Day: 3, Weight: 25.35, Kcal: 2500, TrainingCollection: []statistics_data.OneTraining{}},
		{Day: 4, Weight: 15.35, Kcal: 2500, TrainingCollection: []statistics_data.OneTraining{}},
		{Day: 5, Weight: 95.35, Kcal: 2500, TrainingCollection: []statistics_data.OneTraining{}},
		{Day: 6, Weight: 95.35, Kcal: 2500, TrainingCollection: []statistics_data.OneTraining{}},
		{Day: 7, Weight: 95.35, Kcal: 2500, TrainingCollection: []statistics_data.OneTraining{}},
		{},
	}

	sumCollection := helpers_statictics.SumValue(statistics, "weight")
	if sumCollection[0] != 417.65{
		t.Errorf("error sumVaue weight is not 417.65")
	}

	sumCollection = helpers_statictics.SumValue(statistics, "kcal")
	if sumCollection[0] != 17300{
		t.Errorf("error sumVaue kcal is not 17300")
	}
}

func TestTraining(t *testing.T){
	statistics := []statistics_data.Statistics{
		{Day: 1, Weight: 55.55, Kcal: 2300, TrainingCollection: []statistics_data.OneTraining{{Type: "gym", Time: "2019-10-12T1:10:10Z", Kcal: 1200}}},
		{Day: 2, Weight: 35.35, Kcal: 2500, TrainingCollection: []statistics_data.OneTraining{{Type: "bike", Time: "2019-10-12T00:59:40Z", Kcal: 2700}}},
		{Day: 3, Weight: 25.35, Kcal: 2500, TrainingCollection: []statistics_data.OneTraining{{Type: "bike", Time: "2019-10-12T00:55:40Z", Kcal: 1700}}},
		{Day: 4, Weight: 15.35, Kcal: 2500, TrainingCollection: []statistics_data.OneTraining{{Type: "bike", Time: "2019-10-12T00:45:40Z", Kcal: 1700}}},
		{Day: 5, Weight: 95.35, Kcal: 2500, TrainingCollection: []statistics_data.OneTraining{{Type: "bike", Time: "2019-10-12T00:35:40Z", Kcal: 400}}},
		{Day: 6, Weight: 95.35, Kcal: 2500, TrainingCollection: []statistics_data.OneTraining{{Type: "bike", Time: "2019-10-12T00:25:40Z", Kcal: 300}}},
		{Day: 7, Weight: 95.35, Kcal: 2500, TrainingCollection: []statistics_data.OneTraining{{Type: "bike", Time: "2019-10-12T00:15:40Z", Kcal: 400}}},
		{},
	}

	sumCollection := helpers_statictics.SumTraining(statistics)
	if sumCollection.Data[0].Data[0].Time != "1:10:10"{
		t.Errorf("error SumTraining Time is not 1:10:10")
	}
	if sumCollection.Data[0].Data[0].Type != "gym"{
		t.Errorf("error SumTraining Type is not gym")
	}
	if sumCollection.Data[0].Data[0].SumKcal != 1200{
		t.Errorf("error SumTraining SumKcal is not 1200")
	}
	if sumCollection.Data[0].Data[0].Currecnt != 1{
		t.Errorf("error SumTraining Currecnt is not 1")
	}

	if sumCollection.Data[0].Data[1].Time != "3:58:00"{
		t.Errorf("error SumTraining Time is not 3:58:00")
	}
	if sumCollection.Data[0].Data[1].Type != "bike"{
		t.Errorf("error SumTraining Type is not bike")
	}
	if sumCollection.Data[0].Data[1].SumKcal != 7200{
		t.Errorf("error SumTraining SumKcal is not 7200")
	}
	if sumCollection.Data[0].Data[1].Currecnt != 6{
		t.Errorf("error SumTraining Currecnt is not 6")
	}
}