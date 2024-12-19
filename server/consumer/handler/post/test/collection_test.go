package test

import (
	"fmt"
	common_test "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	post_data "myInternal/consumer/data/post"
	post_function "myInternal/consumer/handler/post"
	project_function_test "myInternal/consumer/handler/project/test"
	helpers "myInternal/consumer/helper"
	"testing"
)

func TestCollectionAll(t *testing.T) {
	var err error
	var params params_data.Params

	dataBody := `{
		"day":1,
		"weight":88,
		"kcal":2500,
		"description":"desc"
	}`

	var createPost post_data.Post
	err = helpers.UnmarshalJSONToType(dataBody, &createPost)
	if err != nil {
		t.Fatalf("error unmarshalling dataBody: %v", err)
	}

	jsonMap, _ := helpers.BindJSONToMap(&createPost)

	params = params_data.Params{
		Header: common_test.UserTest,
		Param: common_test.TestUUid,
		Json: jsonMap,
	}

	//env.LoadEnv("./.env")

	_, err = post_function.Create(params)
	if err != nil {
		t.Fatalf("error in create function: %v", err)
	}

	dataBody = fmt.Sprintf(`{
		"id":"%s"
	}`, common_test.TestUUid)


	var searchPost post_data.SearchPost
	err = helpers.UnmarshalJSONToType(dataBody, &searchPost)
	if err != nil {
		t.Fatalf("error unmarshalling searchPost: %v", err)
	}

	jsonMap, _ = helpers.BindJSONToMap(&searchPost)

	params = params_data.Params{
		Header: common_test.UserTest,
		Param: "1",
		Json: jsonMap,
	}

	valueCollection, err := post_function.Collection(params)
	if err != nil {
		t.Fatalf("error in collection function: %v", err)
	}

	if(len(valueCollection.Collection) == 0){
		t.Fatalf("error in len collection == 0")
	}
}

func TestCollectionOne(t *testing.T){
	var err error
	var params params_data.Params

	dataBody := `{
		"day":1,
		"weight":88,
		"kcal":2500
	}`

	var createPost post_data.Post
	err = helpers.UnmarshalJSONToType(dataBody, &createPost)
	if err != nil {
		t.Fatalf("error unmarshalling dataBody: %v", err)
	}

	jsonMap, _ := helpers.BindJSONToMap(&createPost)

	params = params_data.Params{
		Header: common_test.UserTest,
		Param: common_test.TestUUid,
		Json:   jsonMap,
	}

	//env.LoadEnv("./.env")
	valueCreate, err := post_function.Create(params)
	if err != nil {
			t.Fatalf("error in create function: %v", err)
	}

	params = params_data.Params{
		Header: common_test.UserTest,
		Param: valueCreate.Collection[0].Id,
	}

	_, err = post_function.CollectionOne(params)
	if err != nil {
		t.Fatalf("error in create function: %v", err)
	}
}

func TestCollectionOnePublic(t *testing.T){
	var err error
	var params params_data.Params
	//env.LoadEnv("./.env")

	dataBody := `{
		"day":1,
		"weight":88,
		"kcal":2500
	}`

	var createPost post_data.Post
	err = helpers.UnmarshalJSONToType(dataBody, &createPost)
	if err != nil {
		t.Fatalf("error unmarshalling dataBody: %v", err)
	}

	jsonMap, _ := helpers.BindJSONToMap(&createPost)

	params = params_data.Params{
		Header: common_test.UserTest,
		Param: common_test.TestUUid,
		Json:   jsonMap,
	}

	valueCreate, err := post_function.Create(params)
	if err != nil {
			t.Fatalf("error in create function: %v", err)
	}

	params = params_data.Params{
		Param: valueCreate.Collection[0].Id,
	}

	_, err = post_function.CollectionOnePublic(params)
	if err != nil {
		t.Fatalf("error in create function: %v", err)
	}
}

func TestCollectionPublic(t *testing.T){

	createProjectId, err := project_function_test.CreateProject()
	if err != nil {
		t.Fatalf("error in create project function: %v", err)
	}

	dataBody := `{
		"day":1,
		"weight":88,
		"kcal":2500
	}`

	_, err = CreatePost(dataBody, createProjectId)
	if err != nil {
		t.Fatalf("error in create post function: %v", err)
	}

	collectionPublicPost, err := post_function.CollectionPublic(common_test.UserId, createProjectId, common_test.AppLanguagePL, "1")
	if err != nil {
		t.Fatalf("error collection public post function: %v", err)
	}

	if len(collectionPublicPost.Collection) == 0 {
		t.Fatalf("error in len collection public post == 0")
	}
}