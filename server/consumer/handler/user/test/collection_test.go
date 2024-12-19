package test

import (
	user_function "myInternal/consumer/handler/user"
	env "myInternal/consumer/initializers"
	"testing"
)

func TestCollectionUser(t *testing.T) {

	env.LoadEnv("./.env")
	collectionUser, err := user_function.CollectionUser()
	if err != nil {
		t.Fatalf("Error collection user function: %v", err)
	}

	if len(collectionUser.Collection) == 0 {
		t.Fatalf("Error collection user is empty, len = 0")
	}

}