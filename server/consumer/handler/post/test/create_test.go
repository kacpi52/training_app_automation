package test

import (
	common_test "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	post_data "myInternal/consumer/data/post"
	post_function "myInternal/consumer/handler/post"
	helpers "myInternal/consumer/helper"
	env "myInternal/consumer/initializers"
	"testing"
)

func TestCreatePost(t *testing.T) {

	dataBody := `{
		"day":1,
		"weight":88,
		"kcal":2500
	}`

	var createPost post_data.Post
	err := helpers.UnmarshalJSONToType(dataBody, &createPost); 
	if err != nil {
		t.Fatalf("error unmarshalling dataBody: %v", err)
	}

	jsonMap, _ := helpers.BindJSONToMap(&createPost)

	params := params_data.Params{
		Header: common_test.UserTest,
		Param: common_test.TestUUid,
		Json: jsonMap,
	}

	env.LoadEnv("./.env")
	_, err = post_function.Create(params)
	if err != nil {
		t.Fatalf("error create function: %v", err)
	}
}