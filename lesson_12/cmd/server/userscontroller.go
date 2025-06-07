package main

import (
	"encoding/json"
	"golang-course/lesson_12/internal/users"
)

func getUser(args string) ([]byte, error) {
	usersService := users.NewService(store)

	user, err := usersService.GetUser(args)

	if err != nil {
		return nil, err
	}

	return json.Marshal(user)
}

func addUser(args string) ([]byte, error) {
	usersService := users.NewService(store)

	newUser := users.User{}

	if err := json.Unmarshal([]byte(args), &newUser); err != nil {
		return nil, err
	}

	user, err := usersService.CreateUser(newUser)

	if err != nil {
		return nil, err
	}

	return json.Marshal(user)
}

func deleteUser(args string) ([]byte, error) {
	usersService := users.NewService(store)

	if err := usersService.DeleteUser(args); err != nil {
		return nil, err
	}

	return []byte("User was deleted"), nil
}

func getUsers() ([]byte, error) {
	usersService := users.NewService(store)

	usersList, err := usersService.ListUsers()

	if err != nil {
		return nil, err
	}

	return json.Marshal(usersList)
}
