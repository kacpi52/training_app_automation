package statistics

import (
	params_data "myInternal/consumer/data"
	statistics_data "myInternal/consumer/data/statistics"
	helpers_statictics "myInternal/consumer/handler/statistics/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseCollectionStatistics struct {
	Collection []statistics_data.Create `json:"collection"`
	Status     int 				`json:"status"`
	Error      string 			`json:"error"`
}

func responseCollectionStatisticsStatus(c *gin.Context, col []statistics_data.Create, status int, err error) {
	response := ResponseCollectionStatistics{
		Collection:         col,
		Status:             status,
	}
	
	if err != nil {
		response.Error = err.Error()
	}
	
	c.JSON(status, response)
}

func HandlerCollectionStatistics(c * gin.Context){
	params := params_data.Params{
		Param: c.Param("projectId"),
	}

	createStatisticOption, err := CreateStatisticOption(params)
	if err !=nil{
		responseCollectionStatisticsStatus(c, nil, http.StatusBadRequest, err)
		return
	}

	createStatistics := CollectionStatistics(createStatisticOption.Statistics)
	responseCollectionStatisticsStatus(c, createStatistics, http.StatusOK, err)

}

func CollectionStatistics(data []statistics_data.Statistics) []statistics_data.Create {

	var oneCreate statistics_data.Create
	var collectionStatistics []statistics_data.Create
	j := 1

	sumKg := helpers_statictics.SumValue(data, "weight")
	sumKcal := helpers_statictics.SumValue(data, "kcal")
	sumTraining := helpers_statictics.SumTraining(data)
	sumTime := helpers_statictics.SumTime(data)

	for i, v := range data {

		if i%7 == 0 && len(data) > i+7 {

			if i == 0 {
				oneCreate.Week = 1
			} else {
				oneCreate.Week = (i / 7) + 1
			}

			oneCreate.StartWeight = v.Weight
			oneCreate.EndWeight = data[i+6].Weight
			oneCreate.DownWeight = helpers_statictics.SubtractionFloat(oneCreate.EndWeight, oneCreate.StartWeight)
			oneCreate.SumKg = sumKg[j-1]
			oneCreate.AvgKg = helpers_statictics.DivisionFloat(oneCreate.SumKg, 7)
			oneCreate.SumKcal = sumKcal[j-1]
			oneCreate.Training = append(oneCreate.Training, sumTraining.Data[j-1])
			oneCreate.SumTime = sumTime[j-1]
			collectionStatistics = append(collectionStatistics, oneCreate)
			oneCreate = statistics_data.Create{}
			oneCreate.Training = []statistics_data.OneTrainingWeek{}
			j++
		}
	}

	return collectionStatistics
}