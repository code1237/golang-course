package main

import (
	"fmt"
	documentstore "golang-course/document_store"
)

func main() {
	docString := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key": {
				Type:  documentstore.DocumentFieldTypeString,
				Value: "test",
			},
		},
	}

	docBool := documentstore.Document{
		Fields: map[string]documentstore.DocumentField{
			"key": {
				Type:  documentstore.DocumentFieldTypeBool,
				Value: true,
			},
		},
	}

	documentstore.Put(docString)
	documentstore.Put(docBool)

	fmt.Println("Length of document Store: ", documentstore.Length())

	if _, ok := documentstore.Get("1"); ok {
		fmt.Println("Document by key 1 was found")
	}

	if ok := documentstore.Delete("1"); ok {
		fmt.Println("Document by key 1 was deleted. Length of document Store: ", documentstore.Length())
	}

	documentsList := documentstore.List()
	fmt.Println(documentsList)
}
