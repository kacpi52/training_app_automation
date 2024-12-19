package file

import (
	"archive/zip"
	"fmt"
	"io"
	params_data "myInternal/consumer/data"
	file_data "myInternal/consumer/data/file"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	check_user_permission "myInternal/consumer/helper"
	random "myInternal/consumer/helper"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

type ResponseFileCollection struct {
	Collection []file_data.Collection 	`json:"collection"`
	Status     int 						`json:"status"`
	Error      string 					`json:"error"`
}

type ResponseFileZip struct {
	Zip []byte	`json:"zip"`
	Status     int 						`json:"status"`
	Error      string 					`json:"error"`
}

func HandlerFileCollection(c *gin.Context){
	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Param: c.Param("projectId"),
	}

	fileCollection, err := FileCollection(params)
	if err != nil {
        c.JSON(http.StatusBadRequest, ResponseFileCollection{
            Collection: nil,
            Status:     http.StatusBadRequest,
            Error:      err.Error(),
        })
        return
    }

	c.JSON(http.StatusOK, ResponseFileCollection{
		Collection: fileCollection.Collection,
		Status: fileCollection.Status,
		Error: fileCollection.Error,
	})
}

func HandlerZipDownolad(c *gin.Context){
	params := params_data.Params{
		Header: c.GetHeader("UserData"),
        Param: c.Param("projectId"),
    }

    zipFilePath, err := CreateZip(params)
    if err != nil {
        c.JSON(http.StatusBadRequest, ResponseFileZip{
            Status: http.StatusBadRequest,
            Error:  err.Error(),
        })
        return
    }

    defer func() {
        if err := os.Remove(zipFilePath); err != nil {
            fmt.Printf("Failed to delete zip file: %v\n", err)
        }
    }()

    zipBytes, err := os.ReadFile(zipFilePath)
    if err != nil {
        c.JSON(http.StatusInternalServerError, ResponseFileZip{
            Status: http.StatusInternalServerError,
            Error:  fmt.Sprintf("Failed to read zip file: %v", err),
        })
        return
    }

    zipFileName := filepath.Base(zipFilePath)

    c.Header("Content-Description", "File Transfer")
    c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", zipFileName))
    c.Header("Content-Type", "application/zip")

	c.Data(http.StatusOK, "application/zip", zipBytes)
}

func FileCollection(params params_data.Params)(ResponseFileCollection, error){
	userData := params.Header
    var filesData []file_data.Collection

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseFileCollection{}, err
	}
	defer db.Close()

	_, _,  err = auth.CheckUser(userData)
	if err != nil{
		return ResponseFileCollection{}, err
	}

	projectId := params.Param

	query := `SELECT * FROM images WHERE "projectId" = $1`
	rows, err := db.Query(query, &projectId)
	if err != nil {
		return ResponseFileCollection{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var collection file_data.Collection
		if err := rows.Scan(&collection.Id, &collection.ProjectId, &collection.Name, &collection.Folder, &collection.FolderPath, &collection.Path, &collection.Url, &collection.CreatedUp, &collection.UpdateUp); err != nil {
			return ResponseFileCollection{}, err
		}
		filesData = append(filesData, collection)
	}

	return ResponseFileCollection{
		Collection: filesData,
		Status: http.StatusOK,
		Error: "",
	}, nil
}

func CreateZip(params params_data.Params)(string, error){

	projectId := params.Param
	var collectionPostId []string
	var collectionFolderPath []string

	db, err := database.ConnectToDataBase()
	if err != nil{
		return "", err
	}
	defer db.Close()

	permission, _ := check_user_permission.CheckPermissionsUser(params)
	if permission{
		return "" , fmt.Errorf("permission denied")
	}

	query := `SELECT id FROM post WHERE "projectId" = $1`
	rows, err := db.Query(query, &projectId)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next(){
		var postId string
		if err := rows.Scan(&postId); err != nil {
			return "", err
		}
		collectionPostId = append(collectionPostId, postId)
	}

	if len(collectionPostId) == 0 {
		return "", err
	}

	collectionPostIdStr := strings.Join(collectionPostId, "','")
	query = fmt.Sprintf(`SELECT "folderPath" FROM images WHERE "projectId" IN ('%s')`, collectionPostIdStr)
	rows, err = db.Query(query)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next(){
		var folderPath string
		if err := rows.Scan(&folderPath); err != nil {
			return "", err
		}
		collectionFolderPath = append(collectionFolderPath, folderPath)
	}

	collectionFolderPath = removeDuplicates(collectionFolderPath)

	randomStr, err := random.GenerateRandomString(8)
	folderName := fmt.Sprintf(`project_%s`, randomStr)
	folderPath := filepath.Join("consumer", "download", folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.MkdirAll(folderPath, 0755); err != nil {
			return "", err
		}
	}

	for _, srcFolder := range collectionFolderPath {
		dstFolder := filepath.Join(folderPath, filepath.Base(srcFolder))
		if err := copyDir(srcFolder, dstFolder); err != nil {
			return "", err
		}
	}

	zipPath := folderPath + ".zip"
	if err := zipDir(folderPath, zipPath); err != nil {
		return "", err
	}

	dir := folderPath
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		if err := os.RemoveAll(dir); err != nil {
			return "", err
		}
	}

	return zipPath, err
}

func removeDuplicates(collection []string) []string {
	uniqueMap := make(map[string]bool)
	uniqueList := []string{}

	for _, item := range collection {
		if _, exists := uniqueMap[item]; !exists {
			uniqueMap[item] = true
			uniqueList = append(uniqueList, item)
		}
	}
	return uniqueList
}

func copyDir(src string, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		return err
	})
}

func zipDir(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name, err = filepath.Rel(filepath.Dir(source), path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)
		return err
	})

	return err
}