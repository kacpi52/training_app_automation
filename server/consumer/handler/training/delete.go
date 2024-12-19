package training

import (
	"fmt"
	params_data "myInternal/consumer/data"
	training_data "myInternal/consumer/data/training"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	check_user_permission "myInternal/consumer/helper"
	helpers "myInternal/consumer/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ResponseDeleteTraining struct {
	Collection []training_data.Delete `json:"collection"`
	Status     int                   `json:"status"`
	Error      string                `json:"error"`
}

func responseDeleteStatus(c *gin.Context, col []training_data.Delete, status int, err error) {
	response := ResponseDeleteTraining{
		Collection:         col,
		Status:             status,
	}
	
	if err != nil {
		response.Error = err.Error()
	}
	
	c.JSON(status, response)
}

func HandlerDeleteTraining(c *gin.Context){
	var removeIds training_data.RemoveIds
	c.BindJSON(&removeIds)
	jsonMap, err := helpers.BindJSONToMap(&removeIds)
	if err != nil {
		responseDeleteStatus(c, nil, http.StatusBadRequest, err)
		return
	}

	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Param: c.Param("postId"),
		Json: jsonMap,
	}

	deleteTraining, err := DeleteTraining(params)
	if err != nil{
		responseDeleteStatus(c, nil, http.StatusBadRequest, err)
		return
	}

	responseDeleteStatus(c, deleteTraining.Collection, deleteTraining.Status, nil)
}

func DeleteTraining(params params_data.Params)(ResponseDeleteTraining, error){
	userData := params.Header
	postId := params.Param
	var removeIdsTraning []string
	var deleteTrainings []training_data.Delete

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseDeleteTraining{}, err
	}
	defer db.Close()

	_, _,  err = auth.CheckUser(userData)
	if err != nil{
		return ResponseDeleteTraining{}, err
	}

	permission, _ := check_user_permission.CheckPermissionsUser(params)
	if permission{
		return ResponseDeleteTraining{}, fmt.Errorf("permission denied")
	}

	removeIds, _ := params.Json["removeIds"].([]interface{})

	for _, value := range removeIds{
		id, _ := value.(string)
		removeIdsTraning = append(removeIdsTraning, fmt.Sprintf("'%s'", id))
	}

	if len(removeIdsTraning) != 0 {
		query := `DELETE FROM training WHERE "postId" = $1 AND "id" IN (` + strings.Join(removeIdsTraning, ", ") + `) RETURNING "id", "postId", "type", "time", "kcal", "createdUp", "updateUp";`
		rows, err := db.Query(query, postId)
		if err != nil {
			return ResponseDeleteTraining{}, err
		}
		defer rows.Close()
	
		for rows.Next(){
			var deleteTraining training_data.Delete
			if err := rows.Scan(&deleteTraining.ID, &deleteTraining.PostId, &deleteTraining.Type, &deleteTraining.Time, &deleteTraining.Kcal, &deleteTraining.CreatedUp, &deleteTraining.UpdateUp); err != nil {
				return ResponseDeleteTraining{}, err
			}
			deleteTrainings = append(deleteTrainings, deleteTraining)
	
		}
	}

	return ResponseDeleteTraining{
		Collection: deleteTrainings,
		Status: http.StatusOK,
		Error: "",
	}, nil
}

func DeleteTrainings(params params_data.Params)(ResponseDeleteTraining, error){
	userData := params.Header
	postId := params.Param
	var deleteTrainings []training_data.Delete

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseDeleteTraining{}, err
	}
	defer db.Close()

	_, _,  err = auth.CheckUser(userData)
	if err != nil{
		return ResponseDeleteTraining{}, err
	}

	query := `DELETE FROM training WHERE "postId" = $1 RETURNING "id", "postId", "type", "time", "kcal", "createdUp", "updateUp";`
	rows, err := db.Query(query, postId)
	if err != nil {
		return ResponseDeleteTraining{}, err
	}
	defer rows.Close()

	for rows.Next(){
		var deleteTraining training_data.Delete
		if err := rows.Scan(&deleteTraining.ID, &deleteTraining.PostId, &deleteTraining.Type, &deleteTraining.Time, &deleteTraining.Kcal, &deleteTraining.CreatedUp, &deleteTraining.UpdateUp); err != nil {
			return ResponseDeleteTraining{}, err
		}
		deleteTrainings = append(deleteTrainings, deleteTraining)

	}

	return ResponseDeleteTraining{
		Collection: deleteTrainings,
		Status: http.StatusOK,
		Error: "",
	}, nil
}