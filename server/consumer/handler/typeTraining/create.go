package typetraining

import (
	"fmt"
	params_data "myInternal/consumer/data"
	trainingType_data "myInternal/consumer/data/typeTraining"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	check_user_permission "myInternal/consumer/helper"
	helpers "myInternal/consumer/helper"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ResponseCreateTypeTraning struct{
	Collection []trainingType_data.Create `json:"collection"`
	Status     int                   `json:"status"`
	Error      string                `json:"error"`
}

func responseCreateTypeTraning(c *gin.Context, col []trainingType_data.Create, status int, err error){
	response := ResponseCreateTypeTraning{
		Collection:         col,
		Status:             status,
	}
	
	if err != nil {
		response.Error = err.Error()
	}
	
	c.JSON(status, response)
}

func HandlerCreateTypeTraining(c *gin.Context){
	var typeTraning trainingType_data.Create
	c.BindJSON(&typeTraning)

	jsonMap, err := helpers.BindJSONToMap(&typeTraning)
	if err != nil {
		responseCreateTypeTraning(c, nil, http.StatusBadRequest, err)
		return
	}

	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Json: jsonMap,
	}

	createTypeTraining, err := CreateTypeTraining(params)
	if err != nil{
		responseCreateTypeTraning(c, nil, http.StatusBadRequest, err)
		return
	}

	responseCreateTypeTraning(c, createTypeTraining.Collection, createTypeTraining.Status, nil)
}

func CreateTypeTraining(params params_data.Params)(ResponseCreateTypeTraning, error){

	userData := params.Header
	var typeTrainingCollection []trainingType_data.Create
	var usersData []user_data.User

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseCreateTypeTraning{}, err
	}
	defer db.Close()

	_, users,  err := auth.CheckUser(userData)
	if err != nil{
		return ResponseCreateTypeTraning{}, err
	}
	usersData = users

	permission, _ := check_user_permission.CheckPermissionsUser(params)
	if permission{
		return ResponseCreateTypeTraning{}, fmt.Errorf("permission denied")
	}

	name := params.Json["name"].(string) 
	if name == "" {
		return ResponseCreateTypeTraning{}, fmt.Errorf("json name is error: key 'name' is either missing or not a string")
	}
	now := time.Now()
    formattedDate := now.Format("2006-01-02 15:04:05")

	query := `INSERT INTO type_training ("name", "userId", "createdUp") VALUES ($1, $2, $3) RETURNING "id", "userId", "name", "createdUp";`
	rows, err := db.Query(query, name, usersData[0].Id, formattedDate)
	if err != nil {
		return ResponseCreateTypeTraning{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var typeTraning trainingType_data.Create
		if err := rows.Scan(&typeTraning.Id, &typeTraning.UserId, &typeTraning.Name, &typeTraning.CreatedUp); err != nil {
			return ResponseCreateTypeTraning{}, err
		}
		typeTrainingCollection = append(typeTrainingCollection, typeTraning)
	}

	return ResponseCreateTypeTraning{
		Collection: typeTrainingCollection,
		Status: http.StatusOK,
		Error: "",
	}, nil
}