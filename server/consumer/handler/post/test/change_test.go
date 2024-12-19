package test

import (
	"fmt"
	common_test "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	post_data "myInternal/consumer/data/post"
	post_function "myInternal/consumer/handler/post"
	project_function_test "myInternal/consumer/handler/project/test"
	training_function "myInternal/consumer/handler/training"
	training_function_test "myInternal/consumer/handler/training/test"
	helpers "myInternal/consumer/helper"
	"testing"
)

func TestChange(t *testing.T) {

	var err error
	var params params_data.Params

	dataBody := `{
		"day":1,
		"weight":88,
		"kcal":2500
	}`

	var createPost post_data.Post
	err = helpers.UnmarshalJSONToType(dataBody, &createPost); 
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
	valueCreate, err := post_function.Create(params)
	if err != nil {
		t.Fatalf("error create function: %v", err)
	}


	var changePost post_data.Change

	dataChangeBody := `{
		"day":100,
		"weight":88.5,
		"kcal":3500,
		"collectionTrainingChange":[],
		"collectionTraining":[],
    	"removeIds":[]
	}`

	err = helpers.UnmarshalJSONToType(dataChangeBody, &changePost); 
	if err != nil {
		t.Fatalf("error unmarshalling dataChangeBody: %v", err)
	}

	jsonMap, _ = helpers.BindJSONToMap(&changePost)

	params = params_data.Params{
		Header: common_test.UserTest,
		Param: valueCreate.Collection[0].Id,
		Json: jsonMap,
	}

	changePostF, err := post_function.Change(params)
	if err != nil {
		t.Fatalf("error change function: %v", err)
	}

	if len(changePostF.Collection)==0{
		t.Fatalf("error change collection len is 0")
	}
}

func TestChangeAll(t *testing.T){

	//env.LoadEnv("./.env")
	createProjectId, err := project_function_test.CreateProject()
	if err != nil {
		t.Fatalf("error in create project function: %v", err)
	}

	dataBody := `{
		"day":1,
		"weight":88,
		"kcal":2500
	}`

	postId, err := CreatePost(dataBody, createProjectId)
	if err != nil {
		t.Fatalf("error in create post function: %v", err)
	}

	collectionTraining := `
	{
		"collectionTraining": [
			{
				"type":"gym",
				"time":"2:05:32",
				"kcal":986
			},
			{
				"type":"bike",
				"time":"00:50:19",
				"kcal":543
			},
			{
				"type":"bike",
				"time":"00:48:21",
				"kcal":491
			}
		]
	}
	`
	err = training_function_test.CreateTraining(collectionTraining, postId)
	if err != nil{
		t.Fatalf("error in Create Training function: %v", err)
	}

	params := params_data.Params{
		Header: common_test.UserTest,
		Param: postId,
	}

	collectionTrainingOne, err := training_function.CollectionOneTraining(params)
	if err != nil {
		t.Fatalf("error in collection one training function: %v", err)
	}


	id1 := collectionTrainingOne.Collection[0].ID
	id2 := collectionTrainingOne.Collection[1].ID
	removeId := collectionTrainingOne.Collection[2].ID

	bodyChangePost := fmt.Sprintf(`{
		"day": 100,
		"weight": 88.5,
		"kcal": 3500,
		"collectionTrainingChange": [
			{
				"id": "%s",
				"type": "bike",
				"time": "2:05:32",
				"kcal": 100
			},
			{
				"id": "%s",
				"type": "run",
				"time": "1:50:19",
				"kcal": 100
			}
		],
		"collectionTraining": [
			{
				"type": "bike",
				"time": "00:15:32",
				"kcal": 100
			},
			{
				"type": "run",
				"time": "1:20:19",
				"kcal": 100
			}
		],
		"removeIds": ["%s"]
	}`, id1, id2, removeId)

	
	var changePost post_data.ChangePost
	err = helpers.UnmarshalJSONToType(bodyChangePost, &changePost); 
	if err != nil {
		t.Fatalf("error unmarshalling change post: %v", err)
	}

	jsonMap, err := helpers.BindJSONToMap(&changePost)
	if err != nil {
		t.Fatalf("error BindJSONToMap change post: %v", err)
	}

	params = params_data.Params{
		Header: common_test.UserTest,
		Param: postId,
		Json: jsonMap,
	}

	changePostF, err := post_function.Change(params)
	if err != nil {
		t.Fatalf("error change function: %v", err)
	}

	if *changePostF.Collection[0].Day != 100{
		t.Fatalf("error change function day is not 100")
	}

	if *changePostF.Collection[0].Weight != 88.5{
		t.Fatalf("error change function weight is not 88.5")
	}

	if *changePostF.Collection[0].Kcal != 3500{
		t.Fatalf("error change function kcal is not 3500")
	}

	changeTraining, err := training_function.ChangeTraining(params)
	if err != nil {
		t.Fatalf("error change Training function: %v", err)
	}

	if len(changeTraining.Collection)==0{
		t.Fatalf("error change training collection len is 0 ")
	}

	changeTrainingAdd, err := training_function.CreateTraining(params)
	if err != nil {
		t.Fatalf("error create Training function: %v", err)
	}

	if len(changeTrainingAdd.Collection)==0{
		t.Fatalf("error create training collection len is 0 ")
	}

	changeTrainingRemove, err := training_function.DeleteTraining(params)
	if err != nil {
		t.Fatalf("error create Training function: %v", err)
	}

	if len(changeTrainingRemove.Collection)==0{
		t.Fatalf("error create training collection len is 0 ")
	}
}