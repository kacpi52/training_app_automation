package test

import (
	common_test "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	post_data "myInternal/consumer/data/post"
	training_data "myInternal/consumer/data/training"
	post_function "myInternal/consumer/handler/post"
	post_function_test "myInternal/consumer/handler/post/test"
	project_function_test "myInternal/consumer/handler/project/test"
	training_function "myInternal/consumer/handler/training"
	helpers "myInternal/consumer/helper"
	env "myInternal/consumer/initializers"
	"testing"
)

func TestDeleteTraining(t *testing.T) {
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
	postCreate, err := post_function.Create(params)
	if err != nil {
		t.Fatalf("error create post function: %v", err)
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
	
	var trainingCollectionMap map[string]interface{}
	err = helpers.UnmarshalJSONToType(collectionTraining, &trainingCollectionMap)
	if err != nil {
		t.Fatalf("error unmarshalling trainingCollection: %v", err)
	}
	
	jsonMap, err = helpers.BindJSONToMap(&trainingCollectionMap)
	if err != nil {
		t.Fatalf("error binding JSON to map array: %v", err)
	}

	params = params_data.Params{
		Param: postCreate.Collection[0].Id,
		Json:  jsonMap,
	}

	createTraining, err := training_function.CreateTraining(params)
	if err != nil {
		t.Fatalf("error create training function: %v", err)
	}


	var removeCollection []string
	removeCollection = append(removeCollection, createTraining.Collection[0].ID)
	removeCollection = append(removeCollection, createTraining.Collection[1].ID)

	removeIds := training_data.RemoveIds{
		RemoveIds: removeCollection,
	}

	jsonMap, _ = helpers.BindJSONToMap(&removeIds)

	params = params_data.Params{
		Header: common_test.UserTest,
		Param: postCreate.Collection[0].Id,
		Json: jsonMap,
	}

	deleteTraining, err := training_function.DeleteTraining(params)
	if err != nil {
		t.Fatalf("error delete training function: %v", err)
	}

	if len(deleteTraining.Collection) != 2{
		t.Fatalf("error delete training function is not len == 2: %v", err)
	}
}

func TestDeleteTrainings(t *testing.T){

	projectId, err := project_function_test.CreateProject()
	if err != nil{
		t.Fatalf("error create project function: %v", err)
	}

	dataBody := `{
		"day":1,
		"weight":88,
		"kcal":2500
	}`

	postId, err := post_function_test.CreatePost(dataBody, projectId)
	if err != nil{
		t.Fatalf("error create post function: %v", err)
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

	err = CreateTraining(collectionTraining, postId)
	if err != nil{
		t.Fatalf("error create training function: %v", err)
	}

	params := params_data.Params{
		Header:      common_test.UserTest,
		Param: postId,
	}

	deleteTrainingF, err := training_function.DeleteTrainings(params)
	if err != nil{
		t.Fatalf("error delete trainings function: %v", err)
	}

	if len(deleteTrainingF.Collection)==0{
		t.Fatalf("error delete trainings collection len is 0")
	}
}