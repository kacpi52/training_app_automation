package main

import (
	"context"
	"fmt"
	"log"
	application "myInternal/consumer/application"
	database "myInternal/consumer/database"
	// initializers "myInternal/consumer/initializers"

	_ "github.com/lib/pq"
)

func main(){

	// err := initializers.LoadEnv(".env")
	// if err != nil{
	// 	log.Fatal(err)
	// }

	db ,err := database.ConnectToDataBase()
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	err=database.RunMigration(db)
	if err != nil{
		log.Fatal(err)
	}
	
	app := application.New()
	err = app.Start(context.TODO())
	if err !=nil{
		fmt.Println("failed to start sever app:", err)
	}

}