package test

import (
	"fmt"
	common_test "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	post_data "myInternal/consumer/data/post"
	post_function "myInternal/consumer/handler/post"
	helpers "myInternal/consumer/helper"
	env "myInternal/consumer/initializers"
)

func CreatePost(body string, id string)(string, error){

	var createPost post_data.Post
	err := helpers.UnmarshalJSONToType(body, &createPost); 
	if err != nil {
		return "", fmt.Errorf("error unmarshalling dataBody: %v", err)
	}

	jsonMap, _ := helpers.BindJSONToMap(&createPost)

	params := params_data.Params{
		Header: common_test.UserTest,
		Param: id,
		Json: jsonMap,
	}

	env.LoadEnv("./.env")
	createPostF, err := post_function.Create(params)
	if err != nil {
		return "", fmt.Errorf("error create function: %v", err)
	}

	return createPostF.Collection[0].Id, nil
}
