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
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ResponseChnageProject struct {
	Collection []project_data.Change 	`json:"collection"`
	Status     int 						`json:"status"`
	Error      string 					`json:"error"`
}


func HandlerChangeProject(c *gin.Context) {
	var changePost project_data.Change
	c.BindJSON(&changePost)
	jsonMap, err := helpers.BindJSONToMap(&changePost)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseChnageProject{
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

	change, err := ChangeProject(params)
	if err != nil{
		c.JSON(http.StatusBadRequest, ResponseChnageProject{
			Collection: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseChnageProject{
		Collection: change.Collection,
		Status: change.Status,
		Error: change.Error,
	})
}

func ChangeProject(params params_data.Params)(ResponseChnageProject, error){
	userData := params.Header
	appLanguage := params.AppLanguage
	
	var usersData []user_data.User
	var changesData []project_data.Change


	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseChnageProject{}, err
	}
	defer db.Close()

	_, users,  err := auth.CheckUser(userData)
	if err != nil{
		return ResponseChnageProject{}, err
	}
	usersData = users

	permission, _ := check_user_permission.CheckPermissionsUser(params)
	if permission{
		return ResponseChnageProject{}, fmt.Errorf("permission denied")
	}

	id := params.Param

	title, titleOk := params.Json["title"].(string)
	description, descriptionOk := params.Json["description"].(string)
	now := time.Now()
    formattedDate := now.Format("2006-01-02 15:04:05")
	updateUp := formattedDate
	
	var updateFields []string
	if titleOk {
		updateFields = append(updateFields, fmt.Sprintf(`"title"='%s'`, title)) 
	}
	if descriptionOk {
		updateFields = append(updateFields, fmt.Sprintf(`"description"='%s'`, description))
	}
	
	if len(updateFields) == 0 {
		if err != nil {
			return ResponseChnageProject{}, err
		}
	}

	updateUpData := fmt.Sprintf(`"updateUp"='%s'`, updateUp)
	query := `UPDATE project SET` + updateUpData + ` WHERE "id" = $1 AND "userId" = $2`
	rows, err := db.Query(query, &id, &usersData[0].Id)
	if err != nil {
		return ResponseChnageProject{}, err
	}
	defer rows.Close()


	query = `UPDATE project_multi_language SET ` + strings.Join(updateFields, ", ") + ` WHERE "idProject" = $1 AND "idLanguage" = $2;`
	result, err := db.Exec(query, id, appLanguage)
	if err != nil {
		return ResponseChnageProject{}, err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return ResponseChnageProject{}, err
	}
	
	if rowsAffected == 0 {
		query = `INSERT INTO project_multi_language ("idProject", "idLanguage", "title", "description") VALUES($1, $2, $3, $4);`
		rows, err = db.Query(query, id, appLanguage, title, description)
		if err != nil {
			return ResponseChnageProject{}, err
		}
	} 
	defer rows.Close()

	query = `SELECT p.id, p."userId", pml."idLanguage", pml.title, pml.description, p."createdUp", p."updateUp" FROM project p JOIN 
    project_multi_language pml ON p.id = pml."idProject" WHERE  p.id = $1 AND pml."idLanguage" = $2;`
	rows, err = db.Query(query, id, appLanguage)
    if err != nil {
        return ResponseChnageProject{}, err
    }
	defer rows.Close()

	for rows.Next() {
		var change project_data.Change
		if err := rows.Scan(&change.Id, &change.UserId, &change.IdLanguage, &change.Title, &change.Description, &change.CreatedUp, &change.UpdateUp); err != nil {
			return ResponseChnageProject{}, err
		}
		changesData = append(changesData, change)
	}

	return ResponseChnageProject{
		Collection: changesData,
		Status: http.StatusOK,
		Error: "",
	}, nil
}
