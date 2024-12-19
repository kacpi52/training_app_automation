package project

import (
	"database/sql"
	params_data "myInternal/consumer/data"
	project_data "myInternal/consumer/data/project"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	helpers "myInternal/consumer/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResponseCollectionProject struct{
	Collection []project_data.Collection 	`json:"collection"`
	Status     int 								`json:"status"` 
	Pagination *helpers.PaginationCollectionPost `json:"pagination"`
	Error      string 							`json:"error"`
}

func responseCollectionProject(c *gin.Context, col []project_data.Collection, pagination *helpers.PaginationCollectionPost, status int, err error){
	response := ResponseCollectionProject{
		Pagination: pagination,
		Collection:         col,
		Status:             status,
	}
	
	if err != nil {
		response.Error = err.Error()
	}
	
	c.JSON(status, response)
}

func HandlerCollectionProject(c *gin.Context) {
	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		AppLanguage: c.GetHeader("AppLanguage"),
		Param: c.Param("page"),
	}

	collection, err := CollectionProject(params)
	if err != nil{
		responseCollectionProject(c, nil, nil, http.StatusBadRequest, err)
		return
	}

	responseCollectionProject(c, collection.Collection, collection.Pagination, collection.Status, nil)
}

func HandlerCollectionPublicProject(c *gin.Context){
	var searchProject project_data.SearchProject
	c.BindJSON(&searchProject)

	collection, err := CollectionPublicProjects(searchProject.Id, searchProject.IdLanguage, searchProject.Page)
	if err != nil{
		responseCollectionProject(c, nil, nil, http.StatusBadRequest, err)
		return
	}
	responseCollectionProject(c, collection.Collection, collection.Pagination, collection.Status, nil)
}

func HandlerCollectionAll(c *gin.Context){
	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		AppLanguage: c.GetHeader("AppLanguage"),
	}

	collection, err := CollectionProject(params)
	if err != nil{
		responseCollectionProject(c, nil, nil, http.StatusBadRequest, err)
		return
	}

	responseCollectionProject(c, collection.Collection, nil, collection.Status, nil)
}

func CollectionProject(params params_data.Params)(ResponseCollectionProject, error) {
	userData := params.Header
	appLanguage := params.AppLanguage

    var usersData []user_data.User
    var collectionsData []project_data.Collection
    var query string

    perPage := 16
    page := 1

    db, err := database.ConnectToDataBase()
    if err != nil {
        return ResponseCollectionProject{}, err
    }
	defer db.Close()

	_, users, err := auth.CheckUser(userData)
	if err != nil {
		return ResponseCollectionProject{}, err
	}
	usersData = users

	query = `WITH filtered_projects AS (
		SELECT * 
		FROM project 
		WHERE "userId" = $1 
		ORDER BY "createdUp" DESC
		LIMIT $2 OFFSET $3
		)
		SELECT 
			p.id, 
			p."userId", 
			pml."idLanguage", 
			pml.title, 
			pml.description, 
			p."createdUp", 
			p."updateUp"
		FROM filtered_projects p
		JOIN project_multi_language pml ON p.id = pml."idProject"
		WHERE pml."idLanguage" = $4
		ORDER BY p."createdUp" DESC;
	`

	pageStr := params.Param
    if pageStr != "" {
        page, _ = strconv.Atoi(pageStr)
    }

	pagination := helpers.GetPaginationData(db, "project", usersData[0].Id,  page, perPage, "")

	rows, err := db.Query(query, &usersData[0].Id, perPage, pagination.Offset, appLanguage)
	if err != nil {
		return ResponseCollectionProject{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var collection project_data.Collection
		if err := rows.Scan(&collection.Id, &collection.UserId, &collection.IdLanguage, &collection.Title, &collection.Description, &collection.CreatedUp, &collection.UpdateUp); err != nil {
			return ResponseCollectionProject{}, err
		}
		collectionsData = append(collectionsData, collection)
	}

	return ResponseCollectionProject{
		Collection: collectionsData,
		Status: http.StatusOK,
		Pagination: &pagination,
		Error: "",
	}, nil
}

func CollectionPublicProjects(userId string, appLanguage string, offsetPage string)(ResponseCollectionProject, error){

	var collectionsData []project_data.Collection
	perPage := 16
    page := 1


	db, err := database.ConnectToDataBase()
    if err != nil {
        return ResponseCollectionProject{}, err
    }
	defer db.Close()

	query := `WITH filtered_projects AS (
		SELECT * 
		FROM project 
		WHERE "userId" = $1 
		ORDER BY "createdUp" DESC
		LIMIT $2 OFFSET $3
		)
		SELECT 
			p.id, 
			p."userId", 
			pml."idLanguage", 
			pml.title, 
			pml.description, 
			p."createdUp", 
			p."updateUp"
		FROM filtered_projects p
		JOIN project_multi_language pml ON p.id = pml."idProject"
		WHERE pml."idLanguage" = $4
		ORDER BY p."createdUp" DESC;
	`

	pageStr := offsetPage
    if pageStr != "" {
        page, _ = strconv.Atoi(pageStr)
    }

	pagination := helpers.GetPaginationData(db, "project", userId,  page, perPage,"")

	var rows *sql.Rows
	rows, err = db.Query(query, userId, perPage, pagination.Offset, appLanguage)
	if err != nil{
		return ResponseCollectionProject{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var collection project_data.Collection
		if err := rows.Scan(&collection.Id, &collection.UserId, &collection.IdLanguage, &collection.Title, &collection.Description, &collection.CreatedUp, &collection.UpdateUp); err != nil {
			return ResponseCollectionProject{}, err
		}
		collectionsData = append(collectionsData, collection)
	}

	return ResponseCollectionProject{
		Collection: collectionsData,
		Status: http.StatusOK,
		Pagination: &pagination,
		Error: "",
	}, nil
}

func CollectionAll(params params_data.Params)(ResponseCollectionProject, error){

	userData := params.Header
	appLanguage := params.AppLanguage
	var collectionsData []project_data.Collection
	var usersData []user_data.User

	db, err := database.ConnectToDataBase()
    if err != nil {
        return ResponseCollectionProject{}, err
    }
	defer db.Close()

	_, users, err := auth.CheckUser(userData)
	if err != nil {
		return ResponseCollectionProject{}, err
	}
	usersData = users

	query := `WITH filtered_projects AS (
		SELECT * 
		FROM project 
		WHERE "userId" = $1 
		ORDER BY "createdUp" DESC
		)
		SELECT 
			p.id, 
			p."userId", 
			pml."idLanguage", 
			pml.title, 
			pml.description, 
			p."createdUp", 
			p."updateUp"
		FROM filtered_projects p
		JOIN project_multi_language pml ON p.id = pml."idProject"
		WHERE pml."idLanguage" = $2
		ORDER BY p."createdUp" DESC;
	`

	rows, err := db.Query(query, &usersData[0].Id, appLanguage)
	if err != nil{
		return ResponseCollectionProject{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var collection project_data.Collection
		if err := rows.Scan(&collection.Id, &collection.UserId, &collection.IdLanguage, &collection.Title, &collection.Description, &collection.CreatedUp, &collection.UpdateUp); err != nil {
			return ResponseCollectionProject{}, err
		}
		collectionsData = append(collectionsData, collection)
	}

	return ResponseCollectionProject{
		Collection: collectionsData,
		Status: http.StatusOK,
		Error: "",
	}, nil
}

