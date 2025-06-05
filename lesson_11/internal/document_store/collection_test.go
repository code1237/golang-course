package document_store

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createTestCollection() *Collection {
	return &Collection{
		cfg:  CollectionConfig{PrimaryKey: "id"},
		Name: "users",
		documents: map[string]Document{
			"1": {
				Fields: map[string]DocumentField{
					"id": {
						Type:  DocumentFieldTypeString,
						Value: "1",
					},
					"name": {
						Type:  DocumentFieldTypeString,
						Value: "Go",
					},
					"age": {
						Type:  DocumentFieldTypeNumber,
						Value: float64(34),
					},
				},
			},
		},
	}
}

func TestCollection_Delete(t *testing.T) {
	testCollection := createTestCollection()

	ok := testCollection.Delete("1")
	assert.Equal(t, true, ok)
}

func TestCollection_DeleteNonExistsDocument(t *testing.T) {
	testCollection := createTestCollection()

	ok := testCollection.Delete("2")
	assert.Equal(t, false, ok)
}

func TestCollection_Get(t *testing.T) {
	testCollection := createTestCollection()
	fmt.Println(testCollection)

	doc, err := testCollection.Get("1")
	assert.Nil(t, err)
	assert.Equal(t, testCollection.documents["1"], *doc)
}

func TestCollection_GetNonExistsDocument(t *testing.T) {
	testCollection := createTestCollection()

	doc, err := testCollection.Get("2")
	assert.ErrorIs(t, err, ErrDocumentNotFound)
	assert.Nil(t, doc)
}

func TestCollection_List(t *testing.T) {
	testCollection := createTestCollection()
	documents := testCollection.List()
	assert.Equal(t, 1, len(documents))

	assert.Equal(t, "1", documents[0].Fields["id"].Value)
}

func TestCollection_Put(t *testing.T) {
	testCollection := createTestCollection()

	testDoc := Document{
		Fields: map[string]DocumentField{
			"id": {
				Type:  DocumentFieldTypeString,
				Value: "2",
			},
			"name": {
				Type:  DocumentFieldTypeString,
				Value: "Go2",
			},
			"age": {
				Type:  DocumentFieldTypeNumber,
				Value: 34,
			},
		},
	}

	testCollection.Put(testDoc)
	assert.Equal(t, testDoc, testCollection.documents["2"])
}

func TestCollection_PutWrongDocument(t *testing.T) {
	type fields struct {
		Id   interface{}
		Type DocumentFieldType
	}

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Wrong Field Type",
			fields: fields{
				Id:   "2",
				Type: DocumentFieldTypeNumber,
			},
		},
		{
			name: "Primary key is empty string",
			fields: fields{
				Id:   "",
				Type: DocumentFieldTypeNumber,
			},
		},
		{
			name: "Type casting failed",
			fields: fields{
				Id:   2,
				Type: DocumentFieldTypeString,
			},
		},
		{
			name: "Without Primary key",
			fields: fields{
				Id:   nil,
				Type: DocumentFieldTypeString,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testCollection := createTestCollection()

			testDoc := Document{
				Fields: map[string]DocumentField{
					"id": {
						Type:  tt.fields.Type,
						Value: tt.fields.Id,
					},
					"name": {
						Type:  DocumentFieldTypeString,
						Value: "Go2",
					},
					"age": {
						Type:  DocumentFieldTypeNumber,
						Value: 34,
					},
				},
			}

			if tt.fields.Id == nil {
				delete(testDoc.Fields, "id")
			}

			testCollection.Put(testDoc)

			key, ok := tt.fields.Id.(string)

			if ok {
				_, ok = testCollection.documents[key]
				assert.False(t, ok)
			}
		})
	}
}
