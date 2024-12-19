package user

import (
	"fmt"
	"net/http"
	"strings"

	params_data "myInternal/consumer/data"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	auth "myInternal/consumer/handler/auth"
	check_user_permission "myInternal/consumer/helper"
	helpers "myInternal/consumer/helper"

	"github.com/gin-gonic/gin"
)


type ResponseChangeUser struct{
	Collection []user_data.User `json:"collection"`
	Status     int              `json:"status"`
	Error      string           `json:"error"`
}

func HandlerChangeUser(c *gin.Context){
	
	var changeUser user_data.User
	c.BindJSON(&changeUser)
	jsonMap, err := helpers.BindJSONToMap(&changeUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseChangeUser{
			Collection: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}
		
	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Param: c.Param("id"),
		Json: jsonMap,
	}

	change, err := ChangeUser(params)
	if err != nil{
		c.JSON(http.StatusBadRequest, ResponseChangeUser{
			Collection: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, ResponseChangeUser{
		Collection: change.Collection,
		Status: change.Status,
		Error: change.Error,
	})

	
}

func ChangeUser(params params_data.Params)(ResponseChangeUser, error){

	var user user_data.User
	var users []user_data.User
	userData := params.Header


	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseChangeUser{}, err
	}
	defer db.Close()

	_, users, err = auth.CheckUser(userData)
	if err != nil{
		return ResponseChangeUser{}, err
	}

	permission, _ := check_user_permission.CheckPermissionsUser(params)
	if permission{
		return ResponseChangeUser{}, fmt.Errorf("permission denied")
	}

	userName, userNameOk := params.Json["userName"].(string) 
	lastName, lastNameOk := params.Json["lastName"].(string) 
	nickName, nickNameOk := params.Json["nickName"].(string) 
	email, emailOk := params.Json["email"].(string) 


	var updateFields []string
	if userNameOk {
        updateFields = append(updateFields, fmt.Sprintf(`"userName"='%s'`, userName))
    }
    if lastNameOk {
        updateFields = append(updateFields, fmt.Sprintf(`"lastName"='%s'`, lastName))
    }
    if nickNameOk {
        updateFields = append(updateFields, fmt.Sprintf(`"nickName"='%s'`, nickName))
    }
    if emailOk {
        updateFields = append(updateFields, fmt.Sprintf(`"email"='%s'`, email))
    }

	if len(updateFields) == 0 {
		return ResponseChangeUser{}, err
	}

	query := `UPDATE users SET ` + strings.Join(updateFields, ", ") + ` WHERE id = $1 RETURNING "id", "userName", "lastName", "nickName", "email", "role";`
	row := db.QueryRow(query, users[0].Id)
	err = row.Scan(&user.Id, &user.UserName, &user.LastName, &user.NickName, &user.Email, &user.Role)
	users = []user_data.User{}
	users = append(users, user)
	defer db.Close()

	if err != nil {
		return ResponseChangeUser{}, err
	}

	return ResponseChangeUser{
		Collection: users,
		Status: http.StatusOK,
		Error: "",
	}, nil
}
