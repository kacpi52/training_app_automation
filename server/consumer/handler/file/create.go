package file

import (
	"fmt"
	"io"
	"mime/multipart"
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
	"time"

	"github.com/gin-gonic/gin"
)


type ResponseFileCreate struct {
	Collection []file_data.Create 	`json:"collection"`
	Status     int 					`json:"status"`
	Error      string 				`json:"error"`
}

func HandlerCreateFile(c *gin.Context) {
    formData := make(map[string][]*multipart.FileHeader)
    var nameData []string
    
    for i := 0; ; i++ {
        file, err := c.FormFile(fmt.Sprintf("file[%d]", i))
        if err != nil {
            if err == http.ErrMissingFile {
                break 
            }
            c.JSON(http.StatusBadRequest, ResponseFileCreate{
                Collection: nil,
                Status:     http.StatusBadRequest,
                Error:      err.Error(),
            })
            return
        }
        formData[fmt.Sprintf("file[%d]", i)] = append(formData[fmt.Sprintf("file[%d]", i)], file)
    }

    for j := 0; ; j++ {
        name := c.PostForm(fmt.Sprintf("name[%d]", j))
        if name == "" {
            break
        }
        nameData = append(nameData, name)
    }

    
    params := params_data.Params{
        Header: c.GetHeader("UserData"),
        FormData: formData,
        FormDataParams: map[string]interface{}{
            "projectId": c.PostForm("projectId"),
            "folder": c.PostForm("folder"),
			"names": nameData,
        },
    }

    
    fileCreate, err := CreateFile(params)
    if err != nil {
        c.JSON(http.StatusBadRequest, ResponseFileCreate{
            Collection: nil,
            Status:     http.StatusBadRequest,
            Error:      err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, ResponseFileCreate{
		Collection: fileCreate.Collection,
		Status: fileCreate.Status,
		Error: fileCreate.Error,
	})
}

func CreateFile(params params_data.Params)(ResponseFileCreate, error){
	userData := params.Header
    var filesData []file_data.Create
    filesFormData := params.FormData

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseFileCreate{}, err
	}
    defer db.Close()

	_, _,  err = auth.CheckUser(userData)
	if err != nil{
		return ResponseFileCreate{}, err
	}

    permission, _ := check_user_permission.CheckPermissionsUser(params)
	if permission{
		return ResponseFileCreate{}, fmt.Errorf("permission denied")
	}

    index := 0

	for _, files := range filesFormData {
        for _, file := range files {
            src, err := file.Open()
            if err != nil {
                return ResponseFileCreate{}, err
            }
            defer src.Close()

			fileExtension := filepath.Ext(file.Filename)
			fileNameWithoutExt := file.Filename[:len(file.Filename)-len(fileExtension)]
			randomStr, err := random.GenerateRandomString(8)
            if err != nil {
                return ResponseFileCreate{}, err
            }
            
			folder := removePolishCharsAndCleanWhiteSpace(params.FormDataParams["folder"].(string)) 

			folderPath := filepath.Join("consumer", "file", folder)
			if _, err := os.Stat(folderPath); os.IsNotExist(err) {
				if err := os.MkdirAll(folderPath, 0755); err != nil {
					return ResponseFileCreate{}, err
				}
			}

			var fileName string
			if index < len(params.FormDataParams["names"].([]string)) {
				name := params.FormDataParams["names"].([]string)[index]
				fileName = fmt.Sprintf("%s_%s_%s%s", fileNameWithoutExt, name, randomStr, fileExtension)
			} else {
				fileName = fmt.Sprintf("%s_%s%s", fileNameWithoutExt, randomStr, fileExtension)
			}

            baseURL := os.Getenv("BASEURL")
            dstPath := filepath.Join(folderPath, fileName)
            dst, err := os.Create(dstPath)
            if err != nil {
                return ResponseFileCreate{}, err
            }
            defer dst.Close()
            fullURL := fmt.Sprintf("%s/%s", baseURL, dstPath)

            if _, err := io.Copy(dst, src); err != nil {
                return ResponseFileCreate{}, err
            }


            query := `INSERT INTO images ("projectId", name, folder, "folderPath", path, url, "createdUp", "updateUp") VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING "id", "projectId", name, folder, "folderPath", path, url, "createdUp", "updateUp";`

            id := params.FormDataParams["projectId"].(string)
            name := params.FormDataParams["names"].([]string)[index]
            now := time.Now()
            formattedDate := now.Format("2006-01-02 15:04:05")
            folderPath = fmt.Sprintf("%s\\", filepath.Join("consumer", "file", folder))


            rows, err := db.Query(query, id,name, folder, folderPath, dstPath, fullURL, formattedDate, formattedDate)
            if err != nil {
                return ResponseFileCreate{}, err
            }
            defer rows.Close()

            for rows.Next() {
                var file file_data.Create
                if err := rows.Scan(&file.Id, &file.ProjectId, &file.Name, &file.Folder, &file.FolderPath,  &file.Path, &file.Url, &file.CreatedUp, &file.UpdateUp); err != nil {
                    return ResponseFileCreate{}, err
                }
                filesData = append(filesData, file)
            }
            index++;
        }
        
    }

    return ResponseFileCreate{
        Collection: filesData,
        Status: http.StatusOK,
        Error: "",
    }, nil
}


func removePolishCharsAndCleanWhiteSpace(value string) string {
	polishChars := map[rune]rune{
		'ą': 'a', 'ć': 'c', 'ę': 'e', 'ł': 'l', 'ń': 'n', 'ó': 'o', 'ś': 's', 'ź': 'z', 'ż': 'z',
		'Ą': 'A', 'Ć': 'C', 'Ę': 'E', 'Ł': 'L', 'Ń': 'N', 'Ó': 'O', 'Ś': 'S', 'Ź': 'Z', 'Ż': 'Z',
	}

	var builder strings.Builder
	for _, char := range value {
		if repl, found := polishChars[char]; found {
			builder.WriteRune(repl)
		} else {
			builder.WriteRune(char)
		}
	}

	cleaned := strings.TrimSpace(builder.String())
	cleaned = strings.ReplaceAll(cleaned, " ", "")

	return cleaned
}