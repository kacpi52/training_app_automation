package test

import (
	common_test "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	user_data "myInternal/consumer/data/user"
	user_function "myInternal/consumer/handler/user"
	helpers "myInternal/consumer/helper"
	env "myInternal/consumer/initializers"
	"testing"
)

func TestChangeUser(t *testing.T) {
	var err error
	var params params_data.Params

	dataUserBody := `{
		"userName": "Artur",
    	"lastName": "Scibor",
        "nickName": "artek.scibor",
        "email": "artek.scibor@gmail.com"
	}`

	var changeUser user_data.User
	err = helpers.UnmarshalJSONToType(dataUserBody, &changeUser); 
	if err != nil {
		t.Fatalf("Error unmarshalling dataUserBody: %v", err)
	}

	jsonMap, _ := helpers.BindJSONToMap(&changeUser)

	params = params_data.Params{
		Header: common_test.UserTest2,
		Json: jsonMap,
	}

	env.LoadEnv("./.env")
	_, err = user_function.ChangeUser(params)
	if err != nil {
		t.Fatalf("Error change user function: %v", err)
	}

}