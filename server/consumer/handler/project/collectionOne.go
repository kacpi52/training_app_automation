package project

import (
	"database/sql"
	params_data "myInternal/consumer/data"
	project_data "myInternal/consumer/data/project"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseCollectionOneProject struct {
	Collection []project_data.Collection `json:"collection"`
	Status     int                       `json:"status"`
	Error      string                    `json:"error"`
}


func HandlerCollectionOneProject(c *gin.Context){
	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Query: c.Query("private"),
		AppLanguage: c.GetHeader("AppLanguage"),
		Param: c.Param("projectId"),
	}

	collectionOne, err := CollectionOneProject(params)
	if err != nil{
		c.JSON(http.StatusBadRequest, ResponseCollectionOneProject{
			Collection: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseCollectionOneProject{
		Collection: collectionOne.Collection,
		Status: collectionOne.Status,
		Error: collectionOne.Error,
	})
}

func CollectionOneProject(params params_data.Params)(ResponseCollectionOneProject, error){
	userData := params.Header
	queryParam := params.Query
	appLanguage := params.AppLanguage

	var usersData []user_data.User
	var collectionOneData []project_data.Collection
	var query string

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseCollectionOneProject{}, err
	}
	defer db.Close()

	if queryParam == "true" {
        _, users, err := auth.CheckUser(userData)
        if err != nil {
            return ResponseCollectionOneProject{}, err
        }
        usersData = users

        query = `WITH filtered_projects AS (
			SELECT * 
			FROM project 
			WHERE "userId" = $1
		)
		SELECT 
			p.id, 
			p."userId", 
			pml."idLanguage", 
			pml.title, 
			pml.description, 
			p."createdUp", 
			p."updateUp"
		FROM 
			filtered_projects p
		JOIN 
			project_multi_language pml ON p.id = pml."idProject"
		WHERE 
			p.id = $2
			AND pml."idLanguage" = $3;`
    } else {
        query = `SELECT 
		p.id,
		p."userId",
		pml."idLanguage",
		pml.title,
		pml.description,
		p."createdUp",
		p."updateUp"
		FROM 
			project p
		JOIN 
			project_multi_language pml ON p.id = pml."idProject"
		WHERE 
			p.id = $1
			AND pml."idLanguage" = $2;`
    }

	id := params.Param

	var rows *sql.Rows
    if queryParam == "true" {
        rows, err = db.Query(query, &usersData[0].Id, &id, appLanguage)
    } else {
        rows, err = db.Query(query, &id, appLanguage)
    }

	if err != nil {
		return ResponseCollectionOneProject{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var collection project_data.Collection
		if err := rows.Scan(&collection.Id, &collection.UserId, &collection.IdLanguage, &collection.Title, &collection.Description, &collection.CreatedUp, &collection.UpdateUp); err != nil {
			return ResponseCollectionOneProject{}, err
		}
		collectionOneData = append(collectionOneData, collection)
	}

	return ResponseCollectionOneProject{
		Collection: collectionOneData,
		Status: http.StatusOK,
		Error: "",
	}, nil
}