package training

import (
	"fmt"
	params_data "myInternal/consumer/data"
	training_data "myInternal/consumer/data/training"
	database "myInternal/consumer/database"
	"net/http"
	"strings"
	"time"
)

type ResponseChangeTraining struct {
	Collection []training_data.Change `json:"collection"`
	Status     int               `json:"status"`
	Error      string            `json:"error"`
}

func ChangeTraining(params params_data.Params)(ResponseChangeTraining, error) {
	postId := params.Param
	var trainingData []training_data.Change

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseChangeTraining{}, err
	}
	defer db.Close()

	trainingCollection, _ := params.Json["collectionTrainingChange"].([]interface{})
	now := time.Now()
    formattedDate := now.Format("2006-01-02 15:04:05")

	for _, value := range trainingCollection{
		trainingMap := value.(map[string]interface{})
		training := training_data.OneTraining{
			Id: trainingMap["id"].(string),
            Type: trainingMap["type"].(string),
            Time: trainingMap["time"].(string),
            Kcal: int64(trainingMap["kcal"].(float64)),
        }

		var updateFields []string

		if training.Type != "" {
			updateFields = append(updateFields, fmt.Sprintf(`"type"='%s'`, training.Type)) 
		}

		if training.Time != "" {
			updateFields = append(updateFields, fmt.Sprintf(`"time"='%s'`, training.Time)) 
		}

		if training.Kcal != 0 {
			updateFields = append(updateFields, fmt.Sprintf(`"kcal"=%d`, training.Kcal)) 
		}

		updateFields = append(updateFields, fmt.Sprintf(`"updateUp"='%s'`, formattedDate))

		if len(updateFields) == 0 {
			if err != nil {
				return ResponseChangeTraining{}, fmt.Errorf("updateFields len is 0, %v", err)
			}
		}

		query := `UPDATE training SET` +  strings.Join(updateFields, ", ") + ` WHERE "id" = $1 AND "postId" = $2 RETURNING "id", "postId", "type", "time", "kcal", "createdUp", "updateUp";`
		rows, err := db.Query(query, training.Id, postId)
		if err != nil {
			return ResponseChangeTraining{}, err
		}
		defer rows.Close()

		for rows.Next() {
			var change training_data.Change
			if err := rows.Scan(&change.ID, &change.PostId, &change.Type, &change.Time, &change.Kcal, &change.CreatedUp, &change.UpdateUp); err != nil {
				return ResponseChangeTraining{}, err
			}
			trainingData = append(trainingData, change)
		}
	}

	return ResponseChangeTraining{
		Collection: trainingData,
		Status: http.StatusOK,
		Error: "",
	}, nil
}