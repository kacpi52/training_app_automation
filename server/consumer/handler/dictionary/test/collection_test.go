package test

import (
	common_test "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	dictionary_function "myInternal/consumer/handler/dictionary"
	env "myInternal/consumer/initializers"
	"testing"
)

func TestCollection(t *testing.T) {

	params := params_data.Params{
		AppLanguage: common_test.AppLanguagePL,
	}

	env.LoadEnv("./.env")
	dictionaryCollection, err := dictionary_function.CollectionDictionary(params)
	if err != nil {
		t.Fatalf("error collection dictionary function: %v", err)
	}

	if(len(dictionaryCollection.Collection) < 3){
		t.Fatalf("error collection dictionary is smaller then three, error: %v", err)
	}

}