package application

import (
	auth_handler "myInternal/consumer/handler/auth"
	dictionary_handler "myInternal/consumer/handler/dictionary"
	file_handler "myInternal/consumer/handler/file"
	post_handler "myInternal/consumer/handler/post"
	project_handler "myInternal/consumer/handler/project"
	statistics_handelr "myInternal/consumer/handler/statistics"
	training_handler "myInternal/consumer/handler/training"
	typeTraining_handler "myInternal/consumer/handler/typeTraining"
	user_handler "myInternal/consumer/handler/user"
	"myInternal/consumer/middleware"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func loadRouters() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Static("/consumer/file", "./consumer/file")

	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "Start routers Gin")
	})

	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{os.Getenv("FRONT_URL"), "https://projektdieta.server.arturscibor.pl/", "http://localhost:3000/", "http://localhost:3000"}, 
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"}, 
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "UserData", "AppLanguage"}, 
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour, 
    }))

	// project routers
	projectGroup := router.Group("/api/project")
	{
		projectGroup.POST("/create", middleware.EnsureValidToken(), project_handler.HandlerCreateProject)
		projectGroup.DELETE("delete/:projectId", middleware.EnsureValidToken(), project_handler.HandlerDeleteProject)
		projectGroup.PATCH("/change/:projectId", middleware.EnsureValidToken(), project_handler.HandlerChangeProject)
		projectGroup.GET("/collection/:page", project_handler.HandlerCollectionProject)
		projectGroup.GET("/collectionOne/:projectId", project_handler.HandlerCollectionOneProject)
		projectGroup.GET("/collectionAll", project_handler.HandlerCollectionAll)
		projectGroup.POST("/collectionPublic", project_handler.HandlerCollectionPublicProject)
	}

	//post routers
	postGroup := router.Group("/api/post")
	{
		postGroup.POST("/create/:projectId", middleware.EnsureValidToken(), post_handler.CreateHandler)
		postGroup.POST("/collection/:page", post_handler.HandlerCollection)
		postGroup.GET("/one/:id", post_handler.HandlerCollectionOne)
		postGroup.POST("/collectionOnePublic", post_handler.HandlerCollectionOnePublic)
		postGroup.POST("/collectionPublic", post_handler.HandlerCollectionPublic)
		postGroup.PATCH("/change/:id", middleware.EnsureValidToken(), post_handler.HandlerChange)
		postGroup.DELETE("/delete/:id", middleware.EnsureValidToken(), post_handler.HandlerDelete)
	}

	//file routers
	fileGroup := router.Group("/api/file")
	{
		fileGroup.POST("/create", middleware.EnsureValidToken(), file_handler.HandlerCreateFile)
		fileGroup.DELETE("/delete/:deleteId", middleware.EnsureValidToken(), file_handler.HandlerFileDelete)
		fileGroup.GET("/collection/:projectId", file_handler.HandlerFileCollection)
		fileGroup.DELETE("/deleteAll", middleware.EnsureValidToken(), file_handler.HandlerFileAllDelete)
		fileGroup.POST("/collectionMultiple", file_handler.HandlerFileCollectionMultiple)
		fileGroup.GET("/downolad/zip/:projectId", middleware.EnsureValidToken(), file_handler.HandlerZipDownolad)
	}

	//dictionary routers
	dictionaryGroup := router.Group("/api/dictionary")
	{
		dictionaryGroup.GET("/collection", dictionary_handler.HandlerCollectionDictionary)
	}

	// training routers
	trainingGroup := router.Group("/api/training")
	{
		trainingGroup.DELETE("/delete/:postId", middleware.EnsureValidToken(), training_handler.HandlerDeleteTraining)
		trainingGroup.POST("/create/:postId", middleware.EnsureValidToken(), training_handler.HandlerCreateTraining)
	}

	// typeTraining routers
	typeTrainingGroup := router.Group("/api/typeTraining")
	{
		typeTrainingGroup.POST("/create", middleware.EnsureValidToken(), typeTraining_handler.HandlerCreateTypeTraining)
		typeTrainingGroup.GET("/collection", middleware.EnsureValidToken(), typeTraining_handler.HandlerCollectionTypeTraining)
		typeTrainingGroup.DELETE("/delete/:id", middleware.EnsureValidToken(), typeTraining_handler.HandlerDeleteTypeTraining)
	}

	// statistics routers
	statisticsGroup := router.Group("/api/statistics")
	{
		statisticsGroup.GET("/collection/:projectId", statistics_handelr.HandlerCollectionStatistics)
	}

	// auth jwt
	authHander := &auth_handler.Auth{}
	authGroup := router.Group("/api/auth", middleware.EnsureValidToken())
	{
		authGroup.POST("/authorization", authHander.Authorization)
	}

	//user routers
	userGroup := router.Group("/api/user")
	{
		userGroup.PATCH("/change", middleware.EnsureValidToken(), user_handler.HandlerChangeUser)
		userGroup.GET("/collection", user_handler.HandlerCollectionUser)
	}

	return router
}