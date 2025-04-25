package main

import (
	"fmt"
	documentstore "golang-course/lesson05/document_store"
	"golang-course/lesson05/users"
)

const (
	PrimaryKey          string = "id"
	PrimaryKeyValue     string = "1"
	CollectionUsersName string = "users"
)

func main() {
	store := documentstore.NewStore()

	err, _ := store.CreateCollection(CollectionUsersName, &documentstore.CollectionConfig{PrimaryKey: PrimaryKey})

	if err != nil {
		fmt.Printf("New collection created: %s\n", CollectionUsersName)
	}

	_, err = store.GetCollection(CollectionUsersName)

	if err == nil {
		fmt.Printf("Collection was found in store by name %s\n", CollectionUsersName)
	}

	usersService := users.NewService(store)

	_, err = usersService.CreateUser(users.User{ID: PrimaryKeyValue, Name: "Go"})

	if err != nil {
		return
	}

	if user, ok := usersService.GetUser(PrimaryKeyValue); ok == nil {
		fmt.Printf("Document was found in collection %s by primary key %s\n", users.CollectionUsersName, PrimaryKeyValue)
		fmt.Printf("%+v\n", *user)
	}
}
