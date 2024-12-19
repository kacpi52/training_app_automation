package typetraining

import (
	"fmt"
	params_data "myInternal/consumer/data"
	trainingType_data "myInternal/consumer/data/typeTraining"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	check_user_permission "myInternal/consumer/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseDeleteTypeTraning struct{
	Collection []trainingType_data.Delete `json:"collection"`
	Status     int                   `json:"status"`
	Error      string                `json:"error"`
}

func responseDeleteTypeTraning(c *gin.Context, col []trainingType_data.Delete, status int, err error){
	response := ResponseDeleteTypeTraning{
		Collection:         col,
		Status:             status,
	}
	
	if err != nil {
		response.Error = err.Error()
	}
	
	c.JSON(status, response)
}

func HandlerDeleteTypeTraining(c *gin.Context){
	
	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Param: c.Param("id"),
	}

	deleteTypeTraining, err := DeleteTypeTraining(params)
	if err != nil{
		responseDeleteTypeTraning(c, nil, http.StatusBadRequest, err)
		return
	}

	responseDeleteTypeTraning(c, deleteTypeTraining.Collection, deleteTypeTraining.Status, nil)
}

func DeleteTypeTraining(params params_data.Params)(ResponseDeleteTypeTraning, error){
	userData := params.Header
	typeTrainingId := params.Param
	var typeTrainingDelete []trainingType_data.Delete
	var usersData []user_data.User

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseDeleteTypeTraning{}, err
	}
	defer db.Close()

	_, users,  err := auth.CheckUser(userData)
	if err != nil{
		return ResponseDeleteTypeTraning{}, err
	}
	usersData = users

	permission, _ := check_user_permission.CheckPermissionsUser(params)
	if permission{
		return ResponseDeleteTypeTraning{}, fmt.Errorf("permission denied")
	}

	query := `DELETE FROM type_training WHERE "userId" = $1 AND "id" = $2 RETURNING "id", "userId", "name", "createdUp";`
	rows, err := db.Query(query, usersData[0].Id, typeTrainingId)
	if err != nil {
		return ResponseDeleteTypeTraning{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var typeTraning trainingType_data.Delete
		if err := rows.Scan(&typeTraning.Id, &typeTraning.UserId, &typeTraning.Name, &typeTraning.CreatedUp); err != nil {
			return ResponseDeleteTypeTraning{}, err
		}
		typeTrainingDelete = append(typeTrainingDelete, typeTraning)
	}

	return ResponseDeleteTypeTraning{
		Collection: typeTrainingDelete,
		Status: http.StatusOK,
		Error: "",
	}, nil
}