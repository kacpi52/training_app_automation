package helpers

import (
	"fmt"
	"math"
	statistics_data "myInternal/consumer/data/statistics"
	"strings"
	"time"
)



func SumTime(data []statistics_data.Statistics) []string{

	var arrTime []string
	var currentTime string = "00:00:00"

	for i := 1; i<= len(data)-1; i++{
		for _, j := range data[i-1].TrainingCollection{
			result, _ := AddTime(currentTime, j.Time)
			currentTime = result
		}

		if i%7==0 && len(data)>0{
			arrTime = append(arrTime, currentTime)
			currentTime = "00:00:00"
		}
	}
	return arrTime
}

func AddTime(currentTime string, secondTime string) (string, error){

	t0, err := time.Parse(time.RFC3339, secondTime)
	if err != nil {
		return "", fmt.Errorf("time parsing error 1!: %v", err)
	}
	formattedTime := t0.Format("15:04:05")
	secondTime = formattedTime

	t1, err := time.ParseDuration(formatDuration(currentTime))
	if err != nil{
		return "", fmt.Errorf("time parsing error 1!: %v", err)
	}

	t2, err := time.ParseDuration(formatDuration(secondTime))
	if err != nil{
		return "", fmt.Errorf("time parsing error 2!: %v", err)
	}

	sum := t1 + t2

	hours := int(sum.Hours())
	minutes := int(sum.Minutes()) % 60
	seconds := int(sum.Seconds()) % 60

	return fmt.Sprintf("%d:%02d:%02d", hours, minutes, seconds), nil
}

func formatDuration(t string) string{
	
	t = strings.Replace(t, ":", "h", 1)
	t = strings.Replace(t, ":", "m", 2)
	t += "s"
	
	return t
}

func SumValue(data []statistics_data.Statistics, vType string) []float64 {
	var sum float64
	var rSum []float64

	for i := 1; i <= len(data); i++ {
		if vType == "weight" {
			sum += data[i-1].Weight
			if i != 0 && i%7 == 0 {
				roundedSum := math.Round(sum*100) / 100
				rSum = append(rSum, roundedSum)
				sum = 0
			}
		}

		if vType == "kcal" {
			sum += float64(data[i-1].Kcal)
			if i != 0 && i%7 == 0 {
				roundedSum := math.Round(sum*100) / 100
				rSum = append(rSum, roundedSum)
				sum = 0
			}
		}
	}

	return rSum
}

func SumTraining(data []statistics_data.Statistics) statistics_data.CollectionTraining {

    var collectionTraining statistics_data.CollectionTraining
	var currentTime string = "00:00:00"

    type TrainingSummary struct {
        Current int
        SumKcal int
		SumTime string
    }

    weeklySummary := make(map[string]TrainingSummary)
    arrTrainings := statistics_data.OneTrainingWeek{}
	
    for i := 1; i <= len(data); i++ {
        for _, t := range data[i-1].TrainingCollection {
            summary, exists := weeklySummary[t.Type]
            if !exists {
				currentTime = "00:00:00"
                summary = TrainingSummary{}
            }else{
				currentTime = summary.SumTime
			}
            summary.Current++
            summary.SumKcal += int(t.Kcal)
			resultTime, _ := AddTime(currentTime, t.Time)
			currentTime = resultTime
			summary.SumTime = resultTime
            weeklySummary[t.Type] = summary
        }

        if i%7 == 0 && len(data) > 0 {
            for typ, summary := range weeklySummary {
                arrTrainings.Data = append(arrTrainings.Data, statistics_data.TrainingsWeek{
                    Type:     typ,
                    Currecnt: summary.Current,
                    SumKcal:  summary.SumKcal,
					Time: summary.SumTime,
                })
            }

            if len(arrTrainings.Data) > 0 {
                collectionTraining.Data = append(collectionTraining.Data, arrTrainings)
                arrTrainings.Data = []statistics_data.TrainingsWeek{}
                weeklySummary = make(map[string]TrainingSummary)
            }
        }
    }

    return collectionTraining
}
