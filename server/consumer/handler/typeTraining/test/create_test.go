package test

import (
	common_test "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	typeTraining_data "myInternal/consumer/data/typeTraining"
	typeTraining_function "myInternal/consumer/handler/typeTraining"
	helpers "myInternal/consumer/helper"
	env "myInternal/consumer/initializers"
	"testing"
)

func TestCreateTypeTraining(t *testing.T) {
	dataBody := `{
		"name":"gym"
	}`

	var createTypeTraining typeTraining_data.Create
	err := helpers.UnmarshalJSONToType(dataBody, &createTypeTraining); 
	if err != nil {
		t.Fatalf("error unmarshalling dataBody: %v", err)
	}

	jsonMap, _ := helpers.BindJSONToMap(&createTypeTraining)

	params := params_data.Params{
		Header: common_test.UserTest,
		Json: jsonMap,
	}

	env.LoadEnv("./.env")
	typeTrainingCreate, err := typeTraining_function.CreateTypeTraining(params)
	if err != nil {
		t.Fatalf("error create typeTraining function: %v", err)
	}

	if len(typeTrainingCreate.Collection) != 1{
		t.Fatalf("error collection from typeTraining function is not 1: %v", err)
	}
}