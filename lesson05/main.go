package main

import (
	"fmt"
	documentstore "golang-course/lesson05/document_store"
	"golang-course/lesson05/users"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in", r)
		}
	}()

	store := documentstore.NewStore()

	usersService := users.NewService(store)

	testUsers := []users.User{
		{ID: users.PrimaryKeyValue, Name: "Go"},
		{ID: users.PrimaryKeyValueTwo, Name: "Go2"},
	}

	for _, user := range testUsers {
		if _, err := usersService.CreateUser(user); err != nil {
			fmt.Printf("Failed to create user: %s\n", user.Name)
			return
		}
	}

	usersList, err := usersService.ListUsers()

	if err != nil {
		fmt.Printf("Failed to list users: %s\n", err)
	}

	fmt.Printf("%+v\n", usersList)

	for _, user := range testUsers {
		if _, err = usersService.GetUser(user.ID); err != nil {
			fmt.Println("Failed to get user:", err)
			return
		}
	}

	for _, user := range testUsers {
		if err = usersService.DeleteUser(user.ID); err != nil {
			fmt.Printf("Failed to delete user by id: %s\n", user.ID)
			return
		}
	}

	usersList, err = usersService.ListUsers()

	if err != nil {
		fmt.Printf("Failed to list users: %s\n", err)
	}

	fmt.Printf("%+v\n", usersList)
}
