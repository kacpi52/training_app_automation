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

func TestDeleteTypeTraining(t *testing.T) {
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

	params = params_data.Params{
		Header: common_test.UserTest,
		Param: typeTrainingCreate.Collection[0].Id,
	}

	typeTrainingDelete, err := typeTraining_function.DeleteTypeTraining(params)
	if err != nil {
		t.Fatalf("error delete typeTraining function: %v", err)
	}

	if len(typeTrainingDelete.Collection) != 1{
		t.Fatalf("error collection from typeTraining function delete is not 1")
	}
}