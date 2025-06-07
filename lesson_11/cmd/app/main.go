package main

import (
	"fmt"
	document_store2 "golang-course/lesson_11/internal/document_store"
	"golang-course/lesson_11/internal/users"
	"strconv"
	"sync"
)

func main() {
	store := document_store2.NewStore()

	var wg sync.WaitGroup

	for i := 0; i < 1001; i++ {
		wg.Add(1)
		go func(store *document_store2.Store, i int) {
			defer wg.Done()
			MainFunc(store, i)
		}(store, i)
	}

	wg.Wait()

	usersService := users.NewService(store)

	usersList, err := usersService.ListUsers()

	if err != nil {
		fmt.Printf("Failed to list users: %s\n", err)
	}

	fmt.Printf("%+v\n", usersList)
}

func MainFunc(store *document_store2.Store, id int) {
	usersService := users.NewService(store)

	_, err := usersService.CreateUser(users.User{ID: strconv.Itoa(id), Name: strconv.Itoa(id), Age: id})

	if err != nil {
		fmt.Println(err)
	}

	user, err := usersService.GetUser(strconv.Itoa(id))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", user)

	if err = usersService.DeleteUser(strconv.Itoa(id)); err != nil {
		fmt.Println(err)
	}
}
