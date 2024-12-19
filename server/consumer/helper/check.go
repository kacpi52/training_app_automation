package helper

import (
	params_data "myInternal/consumer/data"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
)

func CheckPermissionsUser(params params_data.Params) (bool, error) {
	userData := params.Header
	var usersData []user_data.User

	db, err := database.ConnectToDataBase()
	if err != nil {
		return false, err
	}
	defer db.Close()

	_, users, err := auth.CheckUser(userData)
	if err != nil {
		return false, err
	}

	usersData = users

	if(*usersData[0].Role == "guest"){
		return true, nil
	}

	return false, nil
}