package test

import (
	common_test "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	project_data "myInternal/consumer/data/project"
	project_function "myInternal/consumer/handler/project"
	helpers "myInternal/consumer/helper"
	env "myInternal/consumer/initializers"
	"testing"
)

func TestCollectionProject(t *testing.T) {
	dataBody := `{
		"title":"test title",
		"description":"desc test",
		"createdUp":"2024-05-12 10:30:00",
		"updateUp":"2024-05-12 10:30:00"
	}`

	var createProject project_data.Create
	err := helpers.UnmarshalJSONToType(dataBody, &createProject); 
	if err != nil {
		t.Fatalf("Error unmarshalling dataBody: %v", err)
	}
	jsonMap, _ := helpers.BindJSONToMap(&createProject)

	params := params_data.Params{
		Header: common_test.UserTest,
		Param: common_test.TestUUid,
		AppLanguage: common_test.AppLanguagePL,
		Json: jsonMap,
	}

	env.LoadEnv("./.env")
	_, err = project_function.CreateProject(params)
	if err != nil {
		t.Fatalf("error create function: %v", err)
	}

	params = params_data.Params{
		Header: common_test.UserTest,
		AppLanguage: common_test.AppLanguagePL,
		Param: "1",
	}

	projectCollection, err := project_function.CollectionProject(params)
	if err != nil {
		t.Fatalf("error collection function: %v", err)
	}

	if(len(projectCollection.Collection) == 0){
		t.Fatalf("error in len collection == 0")
	}
}

func TestCollectionOne(t *testing.T){
	dataBody := `{
		"title":"test title",
		"description":"desc test"
	}`

	var createProject project_data.Create
	err := helpers.UnmarshalJSONToType(dataBody, &createProject); 
	if err != nil {
		t.Fatalf("error unmarshalling dataBody: %v", err)
	}
	jsonMap, _ := helpers.BindJSONToMap(&createProject)

	params := params_data.Params{
		Header: common_test.UserTest,
		Param: common_test.TestUUid,
		AppLanguage: common_test.AppLanguagePL,
		Json: jsonMap,
	}

	env.LoadEnv("./.env")
	project, err := project_function.CreateProject(params)
	if err != nil {
		t.Fatalf("error create function: %v", err)
	}

	params = params_data.Params{
		Header: common_test.UserTest,
		Param: project.Collection[0].Id,
		AppLanguage: common_test.AppLanguagePL,
	}

	projectCollection, err := project_function.CollectionOneProject(params)
	if err != nil {
		t.Fatalf("error collection one function: %v", err)
	}

	if(len(projectCollection.Collection) == 0){
		t.Fatalf("error in len collection one == 0")
	}
}

func TestCollectionPublic(t *testing.T){
	_, err := CreateProject()
	if err != nil {
		t.Fatalf("error create function: %v", err)
	}

	collectionPublic, err := project_function.CollectionPublicProjects(common_test.UserId, common_test.AppLanguagePL, "1")
	if err != nil {
		t.Fatalf("error collection public projects function: %v", err)
	}

	if len(collectionPublic.Collection) == 0 {
		t.Fatalf("error in len collection public project one == 0")
	}
}

func TestCollectionAll(t *testing.T){

	params := params_data.Params{
		Header: common_test.UserTest,
		AppLanguage: common_test.AppLanguagePL,
	}

	_, err := CreateProject()
	if err != nil {
		t.Fatalf("error create function: %v", err)
	}

	collectionAll, err := project_function.CollectionAll(params)
	if err != nil {
		t.Fatalf("error collection all projects function: %v", err)
	}

	if len(collectionAll.Collection) == 0 {
		t.Fatalf("error in len collection all project one == 0")
	}
}