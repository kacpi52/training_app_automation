package test

import (
	"fmt"
	common_test "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	project_data "myInternal/consumer/data/project"
	project_function "myInternal/consumer/handler/project"
	helpers "myInternal/consumer/helper"
	env "myInternal/consumer/initializers"
)

func CreateProject() (string, error) {
	dataBody := `{
		"title":"test title",
		"description":"desc test"
	}`

	var createProject project_data.Create
	err := helpers.UnmarshalJSONToType(dataBody, &createProject)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling dataBody: %v", err)
	}
	jsonMap, _ := helpers.BindJSONToMap(&createProject)

	params := params_data.Params{
		Header:      common_test.UserTest,
		AppLanguage: common_test.AppLanguagePL,
		Json:        jsonMap,
	}

	env.LoadEnv("./.env")
	createProjectF, err := project_function.CreateProject(params)
	if err != nil {
		return "", fmt.Errorf("error create project function: %v", err)
	}

	return createProjectF.Collection[0].Id, nil
}