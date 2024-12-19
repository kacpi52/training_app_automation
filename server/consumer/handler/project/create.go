package project

import (
	"fmt"
	params_data "myInternal/consumer/data"
	project_data "myInternal/consumer/data/project"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	check_user_permission "myInternal/consumer/helper"
	helpers "myInternal/consumer/helper"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ResponseCreateProject struct {
	Collection []project_data.Create 	`json:"collection"`
	Status     int 						`json:"status"`
	Error      string 					`json:"error"`
}


func HandlerCreateProject(c *gin.Context){
	var createProject project_data.Create
	c.BindJSON(&createProject)

	jsonMap, err := helpers.BindJSONToMap(&createProject)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseCreateProject{
			Collection: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Param: c.Param("projectId"),
		AppLanguage: c.GetHeader("AppLanguage"),
		Json: jsonMap,
	}

	project, err := CreateProject(params)
	if err != nil{
		c.JSON(http.StatusBadRequest, ResponseCreateProject{
			Collection: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseCreateProject{
		Collection: project.Collection,
		Status: project.Status,
		Error: project.Error,
	})
}


func CreateProject(params params_data.Params)(ResponseCreateProject, error) {
	userData := params.Header
	appLanguage := params.AppLanguage
	var usersData []user_data.User
	var projectsData []project_data.Create

	if appLanguage == ""{
		return ResponseCreateProject{}, fmt.Errorf("appLanguage is nil or empty: %v", appLanguage)
	}

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseCreateProject{}, err
	}
	defer db.Close()

	_, users,  err := auth.CheckUser(userData)
	if err != nil{
		return ResponseCreateProject{}, err
	}
	usersData = users

	permission, _ := check_user_permission.CheckPermissionsUser(params)
	if permission{
		return ResponseCreateProject{}, fmt.Errorf("permission denied")
	}

	title := params.Json["title"]
	description := params.Json["description"]
	now := time.Now()
    formattedDate := now.Format("2006-01-02 15:04:05")

	query := `INSERT INTO project ("userId", "createdUp", "updateUp") VALUES ($1, $2, $3) RETURNING "id", "userId", "createdUp", "updateUp";`
	rows, err := db.Query(query, usersData[0].Id, formattedDate, formattedDate)
    if err != nil {
        return ResponseCreateProject{}, err
    }
	defer rows.Close()

	var project project_data.Create
	for rows.Next() {
		if err := rows.Scan(&project.Id, &project.UserId, &project.CreatedUp, &project.UpdateUp); err != nil {
			return ResponseCreateProject{}, err
		}
	}

	query = `INSERT INTO project_multi_language ("idProject", "idLanguage", "title", "description") VALUES($1, $2, $3, $4);`
	rows, err = db.Query(query, &project.Id, appLanguage, title, description)
    if err != nil {
        return ResponseCreateProject{}, err
    }
	defer rows.Close()

	query = `SELECT p.id, p."userId", pml."idLanguage", pml.title, pml.description, p."createdUp", p."updateUp" FROM project p JOIN 
    project_multi_language pml ON p.id = pml."idProject" WHERE  p.id = $1 AND pml."idLanguage" = $2;`
	rows, err = db.Query(query, &project.Id, appLanguage)
    if err != nil {
        return ResponseCreateProject{}, err
    }
	defer rows.Close()

	for rows.Next() {
		var project project_data.Create
		if err := rows.Scan(&project.Id, &project.UserId, &project.IdLanguage, &project.Title, &project.Description, &project.CreatedUp, &project.UpdateUp); err != nil {
			return ResponseCreateProject{}, err
		}
		projectsData = append(projectsData, project)
	}

	return ResponseCreateProject{
		Collection: projectsData,
		Status: http.StatusOK,
		Error: "",
	}, nil

}