package file

import (
	"fmt"
	params_data "myInternal/consumer/data"
	file_data "myInternal/consumer/data/file"
	database "myInternal/consumer/database"
	helpers "myInternal/consumer/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ResponseFileCollectionMultiple struct {
	Collection []file_data.Collection 	`json:"collection"`
	Status     int 						`json:"status"`
	Error      string 					`json:"error"`
}

func responseFileCollectionMultiple(c *gin.Context, col []file_data.Collection, status int, err error){
	response := ResponseFileCollectionMultiple{
		Collection:         col,
		Status:             status,
	}
	
	if err != nil {
		response.Error = err.Error()
	}
	
	c.JSON(status, response)
}

func HandlerFileCollectionMultiple(c *gin.Context){

	var ids file_data.CollectionIds
	c.BindJSON(&ids)

	jsonMap, err := helpers.BindJSONToMap(&ids)
	if err != nil{
		responseFileCollectionMultiple(c, nil, http.StatusBadRequest, err)
		return
	}
	
	params := params_data.Params{
		Json: jsonMap,
	}

	fileCollection, err := FileCollectionMultiple(params)
	if err != nil{
		responseFileCollectionMultiple(c, nil, http.StatusBadRequest, err)
		return
	}

	responseFileCollectionMultiple(c, fileCollection.Collection, fileCollection.Status, nil)
}

func FileCollectionMultiple(params params_data.Params)(ResponseFileCollectionMultiple, error){
	var filesData []file_data.Collection
	var collectionId []string

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseFileCollectionMultiple{}, err
	}
	defer db.Close()

	ids, _ := params.Json["ids"].([]interface{})

	for _, value := range ids{
		id, _ := value.(string)
		collectionId = append(collectionId, fmt.Sprintf("'%s'", id))
	}

	query := `SELECT * FROM images WHERE "projectId" IN  (` + strings.Join(collectionId, ", ") + `)`
	rows, err := db.Query(query)
	if err != nil {
		return ResponseFileCollectionMultiple{}, err
	}
	defer rows.Close()

	for rows.Next(){
		var file file_data.Collection
		if err := rows.Scan(&file.Id, &file.ProjectId, &file.Name, &file.Folder, &file.FolderPath,  &file.Path, &file.Url, &file.CreatedUp, &file.UpdateUp); err != nil {
			return ResponseFileCollectionMultiple{}, err
		}
		filesData = append(filesData, file)
	}

	return ResponseFileCollectionMultiple{
		Collection: filesData,
		Status: http.StatusOK,
		Error: "",
	}, nil
}

