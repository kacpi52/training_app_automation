package post

import (
	"fmt"
	params_data "myInternal/consumer/data"
	collection_data "myInternal/consumer/data/post"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	helpers "myInternal/consumer/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResponseCollection struct{
	Collection []collection_data.Collection 	`json:"collection"`
	Status     int 								`json:"status"` 
	Pagination *helpers.PaginationCollectionPost `json:"pagination"`
	Error      string 							`json:"error"`
}

func responseCollectionPost(c *gin.Context, col []collection_data.Collection, pagination *helpers.PaginationCollectionPost, status int, err error){
	response := ResponseCollection{
		Pagination: pagination,
		Collection:         col,
		Status:             status,
	}
	
	if err != nil {
		response.Error = err.Error()
	}
	
	c.JSON(status, response)
}

func HandlerCollection(c *gin.Context){

	var searchPost collection_data.SearchPost
	c.BindJSON(&searchPost)

	jsonMap, _ := helpers.BindJSONToMap(&searchPost)

	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Param: c.Param("page"),
		Json: jsonMap,
	}

	collection, err := Collection(params)
	if err != nil{
		responseCollectionPost(c, nil, nil, http.StatusBadRequest, err)
		return
	}

	responseCollectionPost(c, collection.Collection, collection.Pagination, collection.Status, nil)
}

func HandlerCollectionPublic(c *gin.Context){
	var searchPost collection_data.SearchPost
	c.BindJSON(&searchPost)

	collection, err := CollectionPublic(searchPost.UserId, searchPost.ProjectId, searchPost.IdLanguage, searchPost.Page)
	if err != nil{
		responseCollectionPost(c, nil, nil, http.StatusBadRequest, err)
		return
	}
	responseCollectionPost(c, collection.Collection, collection.Pagination, collection.Status, nil)
}

func Collection(params params_data.Params)(ResponseCollection, error){
	userData := params.Header

    var usersData []user_data.User
    var collectionsData []collection_data.Collection

	projectId := params.Json["id"]
    perPage := 16
    page := 1

    db, err := database.ConnectToDataBase()
    if err != nil {
        return ResponseCollection{}, err
    }
	defer db.Close()

	_, users, err := auth.CheckUser(userData)
	if err != nil {
		return ResponseCollection{}, err
	}
	usersData = users

	pageStr := params.Param
    if pageStr != "" {
        page, _ = strconv.Atoi(pageStr)
    }

	pagination := helpers.GetPaginationData(db, "post", usersData[0].Id, page, perPage, fmt.Sprintf(`"projectId" = '%s'`, projectId))

	query := `
		SELECT * FROM (
		SELECT * FROM post 
		WHERE "userId" = $1 
		AND"projectId" = $2
		ORDER BY "day" ASC
		LIMIT $3 OFFSET $4
	) subquery
	ORDER BY "day" DESC;`
	
    rows, err := db.Query(query, &usersData[0].Id, projectId, perPage, pagination.Offset)
	if err != nil {
		return ResponseCollection{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var collection collection_data.Collection
		if err := rows.Scan(&collection.Id, &collection.UserId, &collection.ProjectId, &collection.Day, &collection.Weight, &collection.Kcal, &collection.CreatedUp, &collection.UpdateUp); err != nil {
			return ResponseCollection{}, err
		}
		collectionsData = append(collectionsData, collection)
	}

	return ResponseCollection{
		Collection: collectionsData,
		Status: http.StatusOK,
		Pagination: &pagination,
		Error: "",
	}, nil
}

func CollectionPublic(userId string, projectId string, appLanguage string, offsetPage string)(ResponseCollection, error){
	var collectionsData []collection_data.Collection

	perPage := 16
    page := 1

	db, err := database.ConnectToDataBase()
    if err != nil {
        return ResponseCollection{}, err
    }
	defer db.Close()

	pageStr := offsetPage
    if pageStr != "" {
        page, _ = strconv.Atoi(pageStr)
    }

	pagination := helpers.GetPaginationData(db, "post", userId,  page, perPage, fmt.Sprintf(`"projectId" = '%s'`, projectId))
	query := `
		SELECT * FROM (
		SELECT * FROM post 
		WHERE "projectId" = $1 
		ORDER BY "day" ASC
		LIMIT $2 OFFSET $3
	) subquery
	ORDER BY "day" DESC;`
    rows, err := db.Query(query, projectId, perPage, pagination.Offset)
	if err != nil {
		return ResponseCollection{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var collection collection_data.Collection
		if err := rows.Scan(&collection.Id, &collection.UserId, &collection.ProjectId, &collection.Day, &collection.Weight, &collection.Kcal, &collection.CreatedUp, &collection.UpdateUp); err != nil {
			return ResponseCollection{}, err
		}
		collectionsData = append(collectionsData, collection)
	}

	return ResponseCollection{
		Collection: collectionsData,
		Status: http.StatusOK,
		Pagination: &pagination,
		Error: "",
	}, nil
}