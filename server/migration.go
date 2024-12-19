package main

import (
	"log"
	database "myInternal/consumer/database"
)

func migrate() {
	db, err := database.ConnectToDataBase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = database.RunMigration(db)
	if err != nil {
		log.Fatal(err)
	}
}

