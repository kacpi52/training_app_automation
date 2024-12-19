package consumer_database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)


func ConnectToDataBase() (*sql.DB, error){

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DBNAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("not found a database: %v", err)
	}
	

	// err = db.Ping()
	// if err != nil {
	// 	db.Close()
	// 	return nil, fmt.Errorf("not pinnging the database: %v", err)
	// }

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	//fmt.Println("Successfully connected to the database!")
	return db, nil
}