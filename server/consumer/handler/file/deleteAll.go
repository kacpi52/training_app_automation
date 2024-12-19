package file

import (
	"fmt"
	params_data "myInternal/consumer/data"
	file_data "myInternal/consumer/data/file"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	check_user_permission "myInternal/consumer/helper"
	helpers "myInternal/consumer/helper"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type ResponseFileAllDelete struct {
	Collection []file_data.Delete `json:"collection"`
	Status     int                `json:"status"`
	Error      string             `json:"error"`
}

func HandlerFileAllDelete(c *gin.Context) {

	var removeIds file_data.RemoveId
	c.BindJSON(&removeIds)
	jsonMap, err := helpers.BindJSONToMap(&removeIds)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseFileAllDelete{
			Collection: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Json: jsonMap,
	}

	fileDelete, err := DeleteFileAll(params)
	if err != nil {
        c.JSON(http.StatusBadRequest, ResponseFileAllDelete{
            Collection: nil,
            Status:     http.StatusBadRequest,
            Error:      err.Error(),
        })
        return
    }

	c.JSON(http.StatusOK, ResponseFileAllDelete{
		Collection: fileDelete.Collection,
		Status: fileDelete.Status,
		Error: fileDelete.Error,
	})
}

func DeleteFileAll(params params_data.Params)(ResponseFileAllDelete, error){
	var deletesData []file_data.Delete
	var folderRemove []string
	userData := params.Header

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseFileAllDelete{}, err
	}
	defer db.Close()

	_, _,  err = auth.CheckUser(userData)
	if err != nil{
		return ResponseFileAllDelete{}, err
	}

	permission, _ := check_user_permission.CheckPermissionsUser(params)
	if permission{
		return ResponseFileAllDelete{}, fmt.Errorf("permission denied")
	}
	
	ids, _ := params.Json["ids"].([]interface{})

	for _, value := range ids{
		id, _ := value.(string)
		folderRemove = append(folderRemove, fmt.Sprintf("'%s'", id))
	}

	query := `DELETE FROM images WHERE "projectId" IN (` + strings.Join(folderRemove, ", ") + `) RETURNING "id", "projectId", "name", "folder", "folderPath", "path", "url", "createdUp", "updateUp";`
	rows, err := db.Query(query)
	if err != nil {
		return ResponseFileAllDelete{}, err
	}
	defer rows.Close()
	folderRemove = []string{}

	for rows.Next(){
		var file file_data.Delete
		if err := rows.Scan(&file.Id, &file.ProjectId, &file.Name, &file.Folder, &file.FolderPath,  &file.Path, &file.Url, &file.CreatedUp, &file.UpdateUp); err != nil {
			return ResponseFileAllDelete{}, err
		}
		deletesData = append(deletesData, file)
		folderRemove = append(folderRemove, file.FolderPath)

	}

	err = removeFolders(folderRemove)
    if err != nil {
		return ResponseFileAllDelete{}, err
    }

	return ResponseFileAllDelete{
		Collection: deletesData,
		Status: http.StatusOK,
		Error: "",
	}, nil

}

func removeFolders(paths []string) error {
    for _, path := range paths {
        err := os.RemoveAll(path)
        if err != nil {
            return err
        }
    }
    return nil
}