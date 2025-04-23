package main

import (
	"fmt"
	"golang-course/documentstore"
)

const (
	PrimaryKey          string = "id"
	PrimaryKeyValue     string = "1"
	CollectionUsersName string = "users"
)

func main() {
	store := documentstore.NewStore()

	ok, _ := store.CreateCollection(CollectionUsersName, &documentstore.CollectionConfig{PrimaryKey: PrimaryKey})

	if ok {
		fmt.Printf("New collection created: %s\n", CollectionUsersName)
	}

	usersCollection, ok := store.GetCollection(CollectionUsersName)

	if ok {
		fmt.Printf("Collection was found in store by name %s\n", CollectionUsersName)
	}

	validDocument := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			PrimaryKey: {
				Type:  documentstore.DocumentFieldTypeString,
				Value: PrimaryKeyValue,
			},
			"name": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "Gopher",
			},
			"isVerified": {
				Type:  documentstore.DocumentFieldTypeBool,
				Value: false,
			},
		},
	}

	usersCollection.Put(validDocument)

	if _, ok := usersCollection.Get(PrimaryKeyValue); ok {
		fmt.Printf("Document was found in collection %s by primary key %s\n", usersCollection.Name, PrimaryKeyValue)
	}

	documentWithoutKey := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"name": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "Golang",
			},
			"isVerified": {
				Type:  documentstore.DocumentFieldTypeBool,
				Value: true,
			},
		},
	}

	usersCollection.Put(documentWithoutKey)

	fmt.Printf("Current length of collection %s: %d\n", usersCollection.Name, len(usersCollection.List()))

	if usersCollection.Delete(PrimaryKeyValue) {
		fmt.Printf("Document was found and deleted in collection %s by primary key %s\n", usersCollection.Name, PrimaryKeyValue)
		fmt.Printf("Current length of collection %s: %d\n", usersCollection.Name, len(usersCollection.List()))
	} else {
		fmt.Printf("Document was not found in collection %s by primary key %s\n", usersCollection.Name, PrimaryKeyValue)
	}

	if store.DeleteCollection(CollectionUsersName) {
		fmt.Printf("Deleted collection: %s\n", CollectionUsersName)
	}

}
