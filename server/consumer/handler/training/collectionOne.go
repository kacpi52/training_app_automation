package training

import (
	params_data "myInternal/consumer/data"
	training_data "myInternal/consumer/data/training"
	database "myInternal/consumer/database"
	"net/http"
)

type ResponseCollectionOneTraining struct {
	Collection []training_data.Collection `json:"collection"`
	Status     int                   `json:"status"`
	Error      string                `json:"error"`
}

func CollectionOneTraining(params params_data.Params)(ResponseCollectionOneTraining, error){
	postId := params.Param
	var collectionOneData []training_data.Collection

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseCollectionOneTraining{}, err
	}
	defer db.Close()

	query := `SELECT * FROM training WHERE "postId" = $1`
	rows, err := db.Query(query, postId)
	if err != nil{
		return ResponseCollectionOneTraining{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var collection training_data.Collection
		if err := rows.Scan(&collection.ID, &collection.PostId, &collection.Type, &collection.Time, &collection.Kcal, &collection.CreatedUp, &collection.UpdateUp); err != nil {
			return ResponseCollectionOneTraining{}, err
		}
		collectionOneData = append(collectionOneData, collection)
	}

	return ResponseCollectionOneTraining{
		Collection: collectionOneData,
		Status: http.StatusOK,
		Error: "",
	}, nil
}