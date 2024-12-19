package consumer_database

import (
	env "myInternal/consumer/initializers"
	"testing"
)

func TestConnectToDataBase(t *testing.T) {

	env.LoadEnv("./../../.env")
	db, err := ConnectToDataBase()
	if err != nil {
		t.Errorf("not connect do database: %s", err.Error())
	}
	defer db.Close()


	
}