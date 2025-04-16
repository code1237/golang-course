package main

import (
	"fmt"
	documentstore "golang-course/document_store"
)

func main() {
	const PrimaryKey string = "1"

	validDocument := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: PrimaryKey,
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

	documentstore.Put(validDocument)
	documentstore.Put(documentWithoutKey)

	fmt.Printf("Length of document Store: %d\n", documentstore.Length())

	if _, ok := documentstore.Get(PrimaryKey); ok {
		fmt.Printf("Document by key %s was found\n", PrimaryKey)
	}

	if ok := documentstore.Delete(PrimaryKey); ok {
		fmt.Printf("Document by key %s was deleted. Length of document Store: %d\n", PrimaryKey, documentstore.Length())
	}

	documentsList := documentstore.List()
	fmt.Println(documentsList)
}
