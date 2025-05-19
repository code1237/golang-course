package main

import (
	"fmt"
	"golang-course/lesson_07/internal/document_store"
	"golang-course/lesson_07/internal/users"
	"log/slog"
	"os"
)

func main() {
	l := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(l)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in", r)
		}
	}()

	store, err := document_store.NewStoreFromFile("dump.txt")

	if err != nil {
		fmt.Println("Error in NewStoreFromDump", err)
		return
	}

	usersService := users.NewService(store)

	usersList, err := usersService.ListUsers()

	if err != nil {
		fmt.Printf("Failed to list users: %s\n", err)
	}

	fmt.Printf("%+v\n", usersList)

	if err, _ = store.CreateCollection("orders", &document_store.CollectionConfig{PrimaryKey: "id"}); err != nil {
		fmt.Printf("Failed to create collection: %s\n", err)
	}

	if err = store.DumpToFile("dumpToFile.txt"); err != nil {
		fmt.Printf("Failed to create dump file: %s\n", err)
	}

	_, err = usersService.CreateUser(users.User{ID: "3", Name: "Log", Age: 77})

	if err != nil {
		fmt.Printf("Failed to create user: %s\n", err)
	}

	if err = usersService.DeleteUser("3"); err != nil {
		fmt.Printf("Failed to delete user: %s\n", err)
	}

	if err = store.DeleteCollection("orders"); err != nil {
		fmt.Printf("Failed to delete collection: %s\n", err)
	}
}
