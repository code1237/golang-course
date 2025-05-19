package document_store_test

import (
	"github.com/stretchr/testify/assert"
	"golang-course/lesson_07/internal/document_store"
	"golang-course/lesson_07/internal/users"
	"testing"
)

func TestMarshalDocument(t *testing.T) {
	user := users.User{
		ID:   "1",
		Name: "Gopher",
		Age:  34,
	}

	doc, err := document_store.MarshalDocument(user)

	assert.Nil(t, err)
	assert.Equal(t, doc.Fields["name"].Value, user.Name)
	assert.Equal(t, doc.Fields["id"].Value, user.ID)
	assert.Equal(t, doc.Fields["age"].Value, float64(user.Age))
}

func TestUnmarshalDocument(t *testing.T) {
	user := &users.User{}

	testDoc := &document_store.Document{
		Fields: map[string]document_store.DocumentField{
			"id": {
				Value: "1",
				Type:  document_store.DocumentFieldTypeString,
			},
			"name": {
				Value: "Gopher",
				Type:  document_store.DocumentFieldTypeString,
			},
			"age": {
				Value: 34,
				Type:  document_store.DocumentFieldTypeNumber,
			},
		},
	}

	err := document_store.UnmarshalDocument(testDoc, &user)

	assert.Nil(t, err)

	assert.Equal(t, testDoc.Fields["id"].Value, user.ID)
	assert.Equal(t, testDoc.Fields["name"].Value, user.Name)
	assert.Equal(t, testDoc.Fields["age"].Value, user.Age)
}
