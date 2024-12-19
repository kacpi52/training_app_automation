package training

import (
	params_data "myInternal/consumer/data"
	training_data "myInternal/consumer/data/training"
	database "myInternal/consumer/database"
	helpers "myInternal/consumer/helper"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ResponseCreateTraining struct {
	Collection []training_data.Create `json:"collection"`
	Status     int                   `json:"status"`
	Error      string                `json:"error"`
}

func responseCreateStatus(c *gin.Context, col []training_data.Create, status int, err error) {
	response := ResponseCreateTraining{
		Collection:         col,
		Status:             status,
	}
	
	if err != nil {
		response.Error = err.Error()
	}
	
	c.JSON(status, response)
}

func HandlerCreateTraining(c *gin.Context){

	var createTraining training_data.CollectionTraining
	c.BindJSON(&createTraining)
	jsonMap, err := helpers.BindJSONToMap(&createTraining)
	if err != nil {
		responseCreateStatus(c, nil, http.StatusBadRequest, err)
		return
	}

	params := params_data.Params{
		Param: c.Param("postId"),
		Json: jsonMap,
	}

	createTrainingF, err := CreateTraining(params)
	if err != nil{
		responseCreateStatus(c, nil, http.StatusBadRequest, err)
		return
	}

	responseCreateStatus(c, createTrainingF.Collection, createTrainingF.Status, nil)
}

func CreateTraining(params params_data.Params)(ResponseCreateTraining, error){
	postId := params.Param
	var trainingData []training_data.Create

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseCreateTraining{}, err
	}
	defer db.Close()

	trainingCollection, _ := params.Json["collectionTraining"].([]interface{})
	now := time.Now()
    formattedDate := now.Format("2006-01-02 15:04:05")
	
	for _, value := range trainingCollection{

		trainingMap := value.(map[string]interface{})
		training := training_data.OneTraining{
            Type: trainingMap["type"].(string),
            Time: trainingMap["time"].(string),
            Kcal: int64(trainingMap["kcal"].(float64)),
        }

		query := `INSERT INTO training("postId", "type", "time", "kcal", "createdUp", "updateUp") VALUES ($1, $2, $3, $4, $5, $6) RETURNING "id", "postId", "type", "time", "kcal", "createdUp", "updateUp";`

		rows, err := db.Query(query, postId, training.Type, training.Time, training.Kcal, formattedDate, formattedDate)
		if err != nil {
			return ResponseCreateTraining{}, err
		}
		defer rows.Close()

		for rows.Next() {
			var training training_data.Create
			if err := rows.Scan(&training.ID, &training.PostId, &training.Type, &training.Time, &training.Kcal, &training.CreatedUp, &training.UpdateUp); err != nil {
				return ResponseCreateTraining{}, err
			}
			trainingData = append(trainingData, training)
		}
	}

	return ResponseCreateTraining{
		Collection: trainingData,
		Status: http.StatusOK,
		Error: "",
	}, nil
}