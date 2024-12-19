package file

import (
	"fmt"
	params_data "myInternal/consumer/data"
	file_data "myInternal/consumer/data/file"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	check_user_permission "myInternal/consumer/helper"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ResponseFileDelete struct {
	Collection []file_data.Delete 	`json:"collection"`
	Status     int 					`json:"status"`
	Error      string 				`json:"error"`
}

func HandlerFileDelete(c *gin.Context) {

	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Param: c.Param("deleteId"),
	}

	fileDelete, err := DeleteFile(params)
	if err != nil {
        c.JSON(http.StatusBadRequest, ResponseFileDelete{
            Collection: nil,
            Status:     http.StatusBadRequest,
            Error:      err.Error(),
        })
        return
    }

	c.JSON(http.StatusOK, ResponseFileDelete{
		Collection: fileDelete.Collection,
		Status: fileDelete.Status,
		Error: fileDelete.Error,
	})

}

func DeleteFile(params params_data.Params)(ResponseFileDelete, error){
	userData := params.Header
    var filesData []file_data.Delete

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseFileDelete{}, err
	}
	defer db.Close()

	_, _,  err = auth.CheckUser(userData)
	if err != nil{
		return ResponseFileDelete{}, err
	}

	permission, _ := check_user_permission.CheckPermissionsUser(params)
	if permission{
		return ResponseFileDelete{}, fmt.Errorf("permission denied")
	}

	id := params.Param

	query := `DELETE FROM images WHERE "id" = $1 RETURNING "id", "projectId", name, folder, "folderPath", path, url, "createdUp", "updateUp";`
	rows, err := db.Query(query, &id)
	if err != nil {
		return ResponseFileDelete{}, err
	}
	defer rows.Close()

	var file file_data.Delete
	for rows.Next() {
		if err := rows.Scan(&file.Id, &file.ProjectId, &file.Name, &file.Folder, &file.FolderPath,  &file.Path, &file.Url, &file.CreatedUp, &file.UpdateUp); err != nil{
			return ResponseFileDelete{}, err
		}

		err := os.Remove(file.Path)
		if err != nil{
			return ResponseFileDelete{}, err
		}

		filesData = append(filesData, file)
	}

	dir := file.FolderPath
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
        files, err := os.ReadDir(dir)
        if err != nil {
            return ResponseFileDelete{}, err
        }

		if len(files) == 0 {
			if err := os.Remove(dir); err != nil {
                return ResponseFileDelete{}, err
            }
        }

    }


	return ResponseFileDelete{
		Collection: filesData,
		Status: http.StatusOK,
		Error: "",
	}, nil
}