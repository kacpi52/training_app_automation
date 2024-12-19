package user

import (
	data_user "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseCollectionUser struct{
	Collection []data_user.User `json:"collection"`
	Status     int                   `json:"status"`
	Error      string                `json:"error"`
}

func responseCollectionUser(c *gin.Context, col []data_user.User, status int, err error){
	response := ResponseCollectionUser{
		Collection:         col,
		Status:             status,
	}
	
	if err != nil {
		response.Error = err.Error()
	}
	
	c.JSON(status, response)
}

func HandlerCollectionUser(c *gin.Context) {
	collectionUsers, err := CollectionUser()
	if err != nil{
		responseCollectionUser(c, nil, http.StatusBadRequest, err)
		return
	}

	responseCollectionUser(c, collectionUsers.Collection, collectionUsers.Status, nil)
}

func CollectionUser()(ResponseCollectionUser, error) {

	var collectionUser []data_user.User

	db, err := database.ConnectToDataBase()
    if err != nil {
        return ResponseCollectionUser{}, err
    }
	defer db.Close()

	query := `SELECT "id", "userName", "lastName", "nickName", "email" FROM users;`
	rows, err := db.Query(query)
	if err != nil {
		return ResponseCollectionUser{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var user data_user.User
		if err := rows.Scan(&user.Id, &user.UserName, &user.LastName, &user.NickName, &user.Email); err != nil {
			return ResponseCollectionUser{}, err
		}
		collectionUser = append(collectionUser, user)
	}

	return ResponseCollectionUser{
		Collection: collectionUser,
		Status: http.StatusOK,
		Error: "",
	}, nil

}