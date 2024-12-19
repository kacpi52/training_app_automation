package statistics

import (
	"fmt"
	params_data "myInternal/consumer/data"
	post_data "myInternal/consumer/data/post"
	statistics_data "myInternal/consumer/data/statistics"
	database "myInternal/consumer/database"
	"net/http"
	"strings"
)

type ResponseCreateStatistics struct {
	Statistics   []statistics_data.Statistics `json:"statistics"`
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func CreateStatisticOption(params params_data.Params)(ResponseCreateStatistics, error){

	projectId := params.Param
	var postsData []post_data.Collection
	var postIds []string
	var trainingData []statistics_data.OneTraining
	var trainingCollection []statistics_data.OneTraining
	var statisticsCollection []statistics_data.Statistics

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseCreateStatistics{}, err
	}
	defer db.Close()

	query := `SELECT * FROM post WHERE "projectId" = $1 ORDER BY day ASC;`
	rows, err := db.Query(query, projectId)
	if err != nil {
		return ResponseCreateStatistics{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var post post_data.Collection
		if err := rows.Scan(&post.Id, &post.UserId, &post.ProjectId, &post.Day, &post.Weight, &post.Kcal, &post.CreatedUp, &post.UpdateUp); err != nil {
			return ResponseCreateStatistics{}, err
		}
		postIds = append(postIds, fmt.Sprintf("'%s'", post.Id))
		postsData = append(postsData, post)
	}

	query = `SELECT "postId", type, time, kcal FROM training WHERE "postId" IN (` + strings.Join(postIds, ", ") + `)`
	rows, err = db.Query(query)
	if err != nil {
		return ResponseCreateStatistics{}, err
	}

	for rows.Next() {
		var training statistics_data.OneTraining
		if err := rows.Scan(&training.PostId, &training.Type, &training.Time, &training.Kcal); err != nil {
			return ResponseCreateStatistics{}, err
		}
		trainingData = append(trainingData, training)
	}

	for _, post := range postsData{

		for _, training := range trainingData{
			if post.Id == training.PostId{
				trainingCollection = append(trainingCollection, training)
			}
		}

		statistics := statistics_data.Statistics{
			Day: post.Day,
			Weight: post.Weight,
			Kcal: post.Kcal,
			TrainingCollection: trainingCollection,
		}

		statisticsCollection = append(statisticsCollection, statistics)
		trainingCollection = []statistics_data.OneTraining{}
	}

	return ResponseCreateStatistics{
		Statistics: statisticsCollection,
		Status: http.StatusOK,
		Error: "",
	}, nil
}
