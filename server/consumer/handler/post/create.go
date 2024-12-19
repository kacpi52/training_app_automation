package post

import (
	"fmt"
	params_data "myInternal/consumer/data"
	post_data "myInternal/consumer/data/post"
	training_data "myInternal/consumer/data/training"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	training_function "myInternal/consumer/handler/training"
	check_user_permission "myInternal/consumer/helper"
	helpers "myInternal/consumer/helper"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ResponseCreate struct {
	Collection []post_data.Post `json:"collection"`
	CollectionTraining []training_data.Create `json:"collectionTraining"`
	Status     int 				`json:"status"`
	Error      string 			`json:"error"`
}

func CreateHandler(c * gin.Context){

	var createPost post_data.CreatePost
	c.ShouldBindJSON(&createPost)

	jsonMap, err := helpers.BindJSONToMap(&createPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseCreate{
			Collection: nil,
			CollectionTraining: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}
		
	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Param: c.Param("projectId"),
		Json: jsonMap,
	}

	craete, err := Create(params)
	if err != nil{
		c.JSON(http.StatusBadRequest, ResponseCreate{
			Collection: nil,
			CollectionTraining: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	params = params_data.Params{
		Param: craete.Collection[0].Id,
		Json: jsonMap,
	}

	createTraining, err := training_function.CreateTraining(params)
	if err != nil{
		c.JSON(http.StatusBadRequest, ResponseCreate{
			Collection: nil,
			CollectionTraining: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseCreate{
		Collection: craete.Collection,
		CollectionTraining: createTraining.Collection,
		Status: craete.Status,
		Error: craete.Error,
	})
}

func Create(params params_data.Params) (ResponseCreate, error){
	userData := params.Header
	var usersData []user_data.User
	var postsData []post_data.Post

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseCreate{}, err
	}
	defer db.Close()

	_, users,  err := auth.CheckUser(userData)
	if err != nil{
		return ResponseCreate{}, err
	}
	usersData = users

	permission, _ := check_user_permission.CheckPermissionsUser(params)
	if permission{
		return ResponseCreate{}, fmt.Errorf("permission denied")
	}

	day := params.Json["day"]
	projectId := params.Param
	weight := params.Json["weight"]
	kcal := params.Json["kcal"]
	now := time.Now()
    formattedDate := now.Format("2006-01-02 15:04:05")


	query := `INSERT INTO post ("userId", "projectId", day, weight, kcal, "createdUp", "updateUp") VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING "id", "userId", "projectId", "day", "weight", "kcal", "createdUp", "updateUp";`

	rows, err := db.Query(query, usersData[0].Id, projectId, day, weight, kcal, formattedDate, formattedDate)
	if err != nil {
		return ResponseCreate{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var post post_data.Post
		if err := rows.Scan(&post.Id, &post.UserId, &post.ProjectId, &post.Day, &post.Weight, &post.Kcal, &post.CreatedUp, &post.UpdateUp); err != nil {
			return ResponseCreate{}, err
		}
		postsData = append(postsData, post)
	}

	return ResponseCreate{
		Collection: postsData,
		Status: http.StatusOK,
		Error: "",
	}, nil
}