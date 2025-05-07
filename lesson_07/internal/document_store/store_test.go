package document_store

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const (
	DumpFileName = "dumpTest.txt"
	Dump         = `{"users":{"documents":{"1":{"age":34,"id":"1","name":"Go"}}}}`
)

func createTestStore() *Store {
	store := NewStore()

	testCollection := createTestCollection()
	store.collections[testCollection.Name] = testCollection

	return store
}

func TestStore_CreateCollection(t *testing.T) {
	store := createTestStore()
	err, coll := store.CreateCollection("test", &CollectionConfig{PrimaryKey: "id"})

	assert.Nil(t, err)
	assert.Equal(t, store.collections["test"], coll)
}

func TestStore_CreateCollectionFailed(t *testing.T) {
	tests := []struct {
		name           string
		collectionName string
		result         error
	}{
		{
			name:           "Collection name is empty",
			collectionName: "",
			result:         ErrCollectionNameCantBeEmpty,
		},
		{
			name:           "Collection name is already exists",
			collectionName: "users",
			result:         ErrCollectionAlreadyExists,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := createTestStore()
			err, coll := store.CreateCollection(tt.collectionName, &CollectionConfig{PrimaryKey: "id"})

			assert.Nil(t, coll)
			assert.ErrorIs(t, err, tt.result)
		})
	}
}

func TestStore_DeleteCollection(t *testing.T) {
	tests := []struct {
		name           string
		collectionName string
		result         error
	}{
		{
			name:           "Collection name exists",
			collectionName: "users",
			result:         nil,
		},
		{
			name:           "Collection name not exists",
			collectionName: "test",
			result:         ErrCollectionNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := createTestStore()
			err := store.DeleteCollection(tt.collectionName)

			assert.Equal(t, tt.result, err)
		})
	}
}

func TestStore_GetCollection(t *testing.T) {
	tests := []struct {
		name           string
		collectionName string
		result         error
	}{
		{
			name:           "Collection name exists",
			collectionName: "users",
			result:         nil,
		},
		{
			name:           "Collection name not exists",
			collectionName: "test",
			result:         ErrCollectionNotFound,
		},

		{
			name:           "Collection name is empty",
			collectionName: "test",
			result:         ErrCollectionNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := createTestStore()
			_, err := store.GetCollection(tt.collectionName)

			assert.Equal(t, tt.result, err)
		})
	}
}

func TestNewStoreFromDump(t *testing.T) {
	dump := []byte(`{"users":{"documents":{"1":{"age":32,"id":"1","name":"Go"},"2":{"age":25,"id":"2","name":"Lang"}}}}`)

	store, err := NewStoreFromDump(dump)

	assert.Nil(t, err)

	assert.Equal(t, store.collections["users"].Name, "users")
	assert.Equal(t, store.collections["users"].documents["1"].Fields["id"].Value, "1")
	assert.Equal(t, store.collections["users"].documents["2"].Fields["id"].Value, "2")
	assert.Equal(t, store.collections["users"].documents["1"].Fields["name"].Value, "Go")
	assert.Equal(t, store.collections["users"].documents["2"].Fields["name"].Value, "Lang")
	assert.EqualValues(t, store.collections["users"].documents["1"].Fields["age"].Value, 32)
	assert.EqualValues(t, store.collections["users"].documents["2"].Fields["age"].Value, 25)
}

func TestStore_Dump(t *testing.T) {
	store := createTestStore()
	dump, err := store.Dump()

	assert.Nil(t, err)
	assert.Equal(t, string(dump), Dump)
}

func TestStore_DumpToFile(t *testing.T) {
	defer os.Remove(DumpFileName)

	store := createTestStore()
	err := store.DumpToFile(DumpFileName)

	assert.Nil(t, err)

	result, err := os.ReadFile(DumpFileName)
	assert.Nil(t, err)

	assert.Equal(t, string(result), Dump)

}

func TestNewStoreFromFile(t *testing.T) {
	defer os.Remove(DumpFileName)

	os.WriteFile(DumpFileName, []byte(Dump), 0644)

	store, err := NewStoreFromFile(DumpFileName)
	assert.Nil(t, err)

	testStore := createTestStore()

	assert.Equal(t, *testStore, *store)
}
