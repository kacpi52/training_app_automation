package test

import (
	"fmt"
	params_data "myInternal/consumer/data"
	training_function "myInternal/consumer/handler/training"
	helpers "myInternal/consumer/helper"
)

func CreateTraining(trainingCollection string, postId string)error{
	var trainingCollectionMap map[string]interface{}
	err := helpers.UnmarshalJSONToType(trainingCollection, &trainingCollectionMap)
	if err != nil {
		return fmt.Errorf("error unmarshalling collectionTraining: %v", err)
	}
	
	jsonMap, err := helpers.BindJSONToMap(&trainingCollectionMap)
	if err != nil {
		return fmt.Errorf("error binding JSON to map array: %v", err)
	}
	
	params := params_data.Params{
		Param: postId,
		Json:  jsonMap,
	}

	_, err = training_function.CreateTraining(params)
	if err != nil {
		return fmt.Errorf("error create training function: %v", err)
	}

	return nil
}
