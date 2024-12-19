package post

import (
	params_data "myInternal/consumer/data"
	post_data "myInternal/consumer/data/post"
	training_data "myInternal/consumer/data/training"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	training_function "myInternal/consumer/handler/training"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseCollectionOne struct{
	Collection []post_data.Collection `json:"collection"`
	CollectionTraining []training_data.Collection `json:"collectionTraining"`
	Status     int							`json:"status"`
	Error      string 						`json:"error"`
}

func responseCollectionOne(c *gin.Context, col []post_data.Collection, colTra []training_data.Collection, status int, err error){
	response := ResponseCollectionOne{
		Collection:         col,
		CollectionTraining: colTra,
		Status:             status,
	}
	
	if err != nil {
		response.Error = err.Error()
	}
	c.JSON(status, response)
}

func HandlerCollectionOne(c *gin.Context){

	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Param: c.Param("id"),
	}

	collectionOne, err := CollectionOne(params)
	if err != nil{
		responseCollectionOne(c, nil, nil, http.StatusBadRequest, err)
		return
	}

	params = params_data.Params{
		Param: collectionOne.Collection[0].Id,
	}

	collectionOneTraining, err := training_function.CollectionOneTraining(params)
	if err != nil{
		responseCollectionOne(c, nil, nil, http.StatusBadRequest, err)
		return
	}

	responseCollectionOne(c, collectionOne.Collection, collectionOneTraining.Collection, collectionOne.Status, nil)
}

func HandlerCollectionOnePublic(c *gin.Context){
	var searchPost post_data.SearchPost
	c.BindJSON(&searchPost)

	params := params_data.Params{
		Param: searchPost.Id,
	}

	collectionOnePublic, err := CollectionOnePublic(params)
	if err != nil{
		responseCollectionOne(c, nil, nil, http.StatusBadRequest, err)
		return
	}

	collectionOneTraining, err := training_function.CollectionOneTraining(params)
	if err != nil{
		responseCollectionOne(c, nil, nil, http.StatusBadRequest, err)
		return
	}

	responseCollectionOne(c, collectionOnePublic.Collection, collectionOneTraining.Collection, collectionOnePublic.Status, nil)
}

func CollectionOne(params params_data.Params)(ResponseCollectionOne, error){
	userData := params.Header
	
	var usersData []user_data.User
	var collectionOneData []post_data.Collection

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseCollectionOne{}, err
	}
	defer db.Close()

	_, users, err := auth.CheckUser(userData)
    if err != nil {
        return ResponseCollectionOne{}, err
    }
    usersData = users
	
	id := params.Param

	query := `SELECT * FROM post WHERE "id" = $1 AND "userId" = $2;`
	rows, err := db.Query(query, &id, &usersData[0].Id)
	if err != nil {
		return ResponseCollectionOne{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var collection post_data.Collection
		if err := rows.Scan(&collection.Id, &collection.UserId, &collection.ProjectId, &collection.Day, &collection.Weight, &collection.Kcal, &collection.CreatedUp, &collection.UpdateUp); err != nil {
			return ResponseCollectionOne{}, err
		}
		collectionOneData = append(collectionOneData, collection)
	}

	return ResponseCollectionOne{
		Collection: collectionOneData,
		CollectionTraining: nil,
		Status: http.StatusOK,
		Error: "",
	}, nil
}

func CollectionOnePublic(params params_data.Params)(ResponseCollectionOne, error){

	id := params.Param
	var collectionOneData []post_data.Collection

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseCollectionOne{}, err
	}
	defer db.Close()

	query := `SELECT * FROM post WHERE "id" = $1;`
	rows, err := db.Query(query, &id)
	if err != nil {
		return ResponseCollectionOne{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var collection post_data.Collection
		if err := rows.Scan(&collection.Id, &collection.UserId, &collection.ProjectId, &collection.Day, &collection.Weight, &collection.Kcal, &collection.CreatedUp, &collection.UpdateUp); err != nil {
			return ResponseCollectionOne{}, err
		}
		collectionOneData = append(collectionOneData, collection)
	}

	return ResponseCollectionOne{
		Collection: collectionOneData,
		CollectionTraining: nil,
		Status: http.StatusOK,
		Error: "",
	}, nil
}