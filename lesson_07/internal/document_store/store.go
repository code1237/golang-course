package document_store

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

const (
	DocumentsDumpKey = "documents"
)

type Store struct {
	collections map[string]*Collection
}

func NewStore() *Store {
	store := &Store{
		collections: make(map[string]*Collection),
	}

	return store
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (error, *Collection) {
	if len(name) == 0 {
		return ErrCollectionNameCantBeEmpty, nil
	}

	if _, ok := s.collections[name]; ok {
		return ErrCollectionAlreadyExists, nil
	}

	newCollection := &Collection{
		*cfg,
		name,
		make(map[string]Document),
	}

	s.collections[name] = newCollection
	slog.Info("Collection created", "name", name)

	return nil, newCollection
}

func (s *Store) GetCollection(name string) (*Collection, error) {
	collection, ok := s.collections[name]

	if !ok {
		return nil, ErrCollectionNotFound
	}

	return collection, nil
}

func (s *Store) DeleteCollection(name string) error {
	if _, ok := s.collections[name]; ok {
		delete(s.collections, name)
		slog.Info("Collection deleted", "name", name)
		return nil
	}

	return ErrCollectionNotFound
}

func NewStoreFromDump(dump []byte) (*Store, error) {
	var storeCollections = make(map[string]*Collection)

	var tempCollections map[string]map[string]any
	if err := json.Unmarshal(dump, &tempCollections); err != nil {

		return nil, fmt.Errorf("error unmarshalling dump: %w", err)
	}

	for name, collection := range tempCollections {
		documents := make(map[string]Document)

		if docs, ok := collection[DocumentsDumpKey]; ok {
			for primaryKey, data := range docs.(map[string]any) {
				if document, err := MarshalDocument(data); err == nil {
					documents[primaryKey] = *document
				}
			}
		}

		storeCollections[name] = &Collection{
			CollectionConfig{PrimaryKey: "id"},
			name,
			documents,
		}

	}

	return &Store{storeCollections}, nil

}

func (s *Store) Dump() ([]byte, error) {
	var dumpMap = make(map[string]map[string]any)

	for name, collection := range s.collections {
		var documentsMap = make(map[string]any)
		for _, document := range collection.List() {
			var documentFields = make(map[string]any)

			for fieldKey, fieldData := range document.Fields {
				documentFields[fieldKey] = fieldData.Value
			}

			documentsMap[document.Fields[collection.cfg.PrimaryKey].Value.(string)] = documentFields
		}

		dumpMap[name] = make(map[string]any)
		dumpMap[name][DocumentsDumpKey] = documentsMap
	}

	return json.Marshal(dumpMap)
}

func (s *Store) DumpToFile(filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)

	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println("Error closing file", err)
		}
	}()

	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	data, err := s.Dump()

	if err != nil {
		return err
	}

	if _, err = file.WriteString(string(data)); err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}

func NewStoreFromFile(filename string) (*Store, error) {
	dump, err := os.ReadFile(filename)

	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return NewStoreFromDump(dump)
}
