package project

import (
	"fmt"
	params_data "myInternal/consumer/data"
	post_data "myInternal/consumer/data/post"
	project_data "myInternal/consumer/data/project"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	check_user_permission "myInternal/consumer/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseDeleteProject struct {
	Collection []project_data.Delete 	`json:"collection"`
	CollectionRemoveId []string			`json:"collectionRemoveId"`
	Status     int 						`json:"status"`
	Error      string 					`json:"error"`
}


func HandlerDeleteProject(c *gin.Context) {
	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Param: c.Param("projectId"),
	}

	projectDelete, err := DeleteProject(params)
	if err != nil{
		c.JSON(http.StatusBadRequest, ResponseDeleteProject{
			Collection: nil,
			CollectionRemoveId: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseDeleteProject{
		Collection: projectDelete.Collection,
		CollectionRemoveId: projectDelete.CollectionRemoveId,
		Status: projectDelete.Status,
		Error: projectDelete.Error,
	})
}

func DeleteProject(params params_data.Params)(ResponseDeleteProject, error){
	userData := params.Header
	var usersData []user_data.User
	var deletesData []project_data.Delete
	var removeFiles []string

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseDeleteProject{}, err
	}
	defer db.Close()

	_, users, err := auth.CheckUser(userData)
	if err != nil{
		return ResponseDeleteProject{}, err
	}

	permission, _ := check_user_permission.CheckPermissionsUser(params)
	if permission{
		return ResponseDeleteProject{}, fmt.Errorf("permission denied")
	}

	usersData = users
	projectId := params.Param

	query := `DELETE FROM post WHERE "projectId" = $1 AND "userId" = $2 RETURNING "id", "userId", "projectId", "day", "weight", "kcal", "createdUp", "updateUp";`
	rows, err := db.Query(query, &projectId, usersData[0].Id)
	if err != nil {
		return ResponseDeleteProject{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var post post_data.Delete
		if err := rows.Scan(&post.Id, &post.UserId, &post.ProjectId, &post.Day, &post.Weight, &post.Kcal, &post.CreatedUp, &post.UpdateUp); err != nil {
			return ResponseDeleteProject{}, err
		}
		removeFiles = append(removeFiles, post.Id)
	}

	query = `DELETE FROM project_multi_language WHERE "idProject" = $1;`
	rows, err = db.Query(query, &projectId)
	if err != nil {
		return ResponseDeleteProject{}, err
	}
	defer rows.Close()

	query = `DELETE FROM project WHERE "id" = $1 AND "userId" = $2 RETURNING "id", "userId", "createdUp", "updateUp";`
	rows, err = db.Query(query, &projectId, usersData[0].Id)
	if err != nil {
		return ResponseDeleteProject{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var project project_data.Delete
		if err := rows.Scan(&project.Id, &project.UserId, &project.CreatedUp, &project.UpdateUp); err != nil {
			return ResponseDeleteProject{}, err
		}
		removeFiles = append(removeFiles, projectId)
		deletesData = append(deletesData, project)
	}

	return ResponseDeleteProject{
		Collection: deletesData,
		CollectionRemoveId: removeFiles,
		Status: http.StatusOK,
		Error: "",
	}, nil
}