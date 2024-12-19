package test

import (
	"fmt"
	common_test "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	post_data "myInternal/consumer/data/post"
	post_function "myInternal/consumer/handler/post"
	training_function "myInternal/consumer/handler/training"
	helpers "myInternal/consumer/helper"
	env "myInternal/consumer/initializers"
	"testing"
)

func TestChangeTraining(t *testing.T){
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

	trainingChange := fmt.Sprintf(`
	{
		"collectionTraining": [
			{
				"id":"%s",
				"type":"gym",
				"time":"2:05:32",
				"kcal":986
			},
			{
				"id":"%s",
				"type":"bike",
				"time":"00:50:19",
				"kcal":543
			},
			{
				"id":"%s",
				"type":"bike",
				"time":"00:48:21",
				"kcal":491
			}
		]
	}
	`, createTraining.Collection[0].ID, createTraining.Collection[1].ID, createTraining.Collection[2].ID)

	var trainingChangeMap map[string]interface{}
	err = helpers.UnmarshalJSONToType(trainingChange, &trainingChangeMap)
	if err != nil {
		t.Fatalf("error unmarshalling trainingChange: %v", err)
	}
	
	jsonMap, err = helpers.BindJSONToMap(&trainingChangeMap)
	if err != nil {
		t.Fatalf("error binding JSON to map array: %v", err)
	}

	params = params_data.Params{
		Param: postCreate.Collection[0].Id,
		Json:  jsonMap,
	}

	_, err = training_function.ChangeTraining(params)
	if err != nil {
		t.Fatalf("error change training function: %v", err)
	}
}