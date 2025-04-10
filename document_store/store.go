package documentstore

import (
	"strconv"
)

type DocumentFieldType string

const (
	DocumentFieldTypeString DocumentFieldType = "string"
	DocumentFieldTypeNumber DocumentFieldType = "number"
	DocumentFieldTypeBool   DocumentFieldType = "bool"
	DocumentFieldTypeArray  DocumentFieldType = "array"
	DocumentFieldTypeObject DocumentFieldType = "object"
)

type DocumentField struct {
	Type  DocumentFieldType
	Value interface{}
}

type Document struct {
	Fields map[string]DocumentField
}

var documents = map[string]Document{}

func Put(doc Document) {
	for _, field := range doc.Fields {
		if field.Type == DocumentFieldTypeString {
			documents[strconv.Itoa(len(documents)+1)] = doc
			break
		}
	}
}

func Get(key string) (*Document, bool) {
	if doc, ok := documents[key]; ok {
		return &doc, true
	}

	return nil, false
}

func Delete(key string) bool {
	if _, ok := documents[key]; ok {
		delete(documents, key)
		return true
	}

	return false
}

func List() []Document {
	var documentsSlice []Document

	for _, document := range documents {
		documentsSlice = append(documentsSlice, document)
	}

	return documentsSlice
}

func Length() int {
	return len(documents)
}
