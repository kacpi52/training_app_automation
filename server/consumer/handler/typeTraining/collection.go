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

type ResponseCollectionTypeTraning struct{
	Collection []trainingType_data.Collection `json:"collection"`
	Status     int                   `json:"status"`
	Error      string                `json:"error"`
}

func responseCollectionTypeTraning(c *gin.Context, col []trainingType_data.Collection, status int, err error){
	response := ResponseCollectionTypeTraning{
		Collection:         col,
		Status:             status,
	}
	
	if err != nil {
		response.Error = err.Error()
	}
	
	c.JSON(status, response)
}

func HandlerCollectionTypeTraining(c *gin.Context){

	params := params_data.Params{
		Header: c.GetHeader("UserData"),
	}

	collectionTypeTraining, err := CollectionTypeTraining(params)
	if err != nil{
		responseCollectionTypeTraning(c, nil, http.StatusBadRequest, err)
		return
	}

	responseCollectionTypeTraning(c, collectionTypeTraining.Collection, collectionTypeTraining.Status, nil)
}

func CollectionTypeTraining(params params_data.Params)(ResponseCollectionTypeTraning, error){
	userData := params.Header
	var usersData []user_data.User
	var collectionTypeTraining []trainingType_data.Collection


	db, err := database.ConnectToDataBase()
    if err != nil {
        return ResponseCollectionTypeTraning{}, err
    }
	defer db.Close()

	_, users, err := auth.CheckUser(userData)
        if err != nil {
            return ResponseCollectionTypeTraning{}, err
        }
    usersData = users
	
	permission, _ := check_user_permission.CheckPermissionsUser(params)
	if permission{
		return ResponseCollectionTypeTraning{}, fmt.Errorf("permission denied")
	}

	query := `SELECT * FROM type_training WHERE "userId" = $1;`
	rows, err := db.Query(query, &usersData[0].Id)
	if err != nil {
		return ResponseCollectionTypeTraning{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var typeTraining trainingType_data.Collection
		if err := rows.Scan(&typeTraining.Id, &typeTraining.UserId, &typeTraining.Name, &typeTraining.CreatedUp); err != nil {
			return ResponseCollectionTypeTraning{}, err
		}
		collectionTypeTraining = append(collectionTypeTraining, typeTraining)
	}

	return ResponseCollectionTypeTraning{
		Collection: collectionTypeTraining,
		Status: http.StatusOK,
		Error: "",
	}, nil
}

