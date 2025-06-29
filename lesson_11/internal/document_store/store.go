package document_store

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"sync"
)

type dumpDTO struct {
	Documents map[string]any   `json:"documents"`
	Config    CollectionConfig `json:"config"`
}

type Store struct {
	mx          sync.RWMutex
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

	s.mx.RLock()
	_, ok := s.collections[name]
	s.mx.RUnlock()

	if ok {
		return ErrCollectionAlreadyExists, nil
	}

	newCollection := &Collection{
		cfg:       *cfg,
		Name:      name,
		documents: make(map[string]Document),
	}

	s.mx.Lock()
	s.collections[name] = newCollection
	s.mx.Unlock()

	slog.Info("Collection created", "name", name)

	return nil, newCollection
}

func (s *Store) GetCollection(name string) (*Collection, error) {
	s.mx.RLock()
	collection, ok := s.collections[name]
	s.mx.RUnlock()

	if !ok {
		return nil, ErrCollectionNotFound
	}

	return collection, nil
}

func (s *Store) DeleteCollection(name string) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	if _, ok := s.collections[name]; ok {
		delete(s.collections, name)
		slog.Info("Collection deleted", "name", name)
		return nil
	}

	return ErrCollectionNotFound
}

func NewStoreFromDump(dump []byte) (*Store, error) {
	var storeCollections = make(map[string]*Collection)

	var tempCollections map[string]dumpDTO
	if err := json.Unmarshal(dump, &tempCollections); err != nil {
		return nil, fmt.Errorf("error unmarshalling dump: %w", err)
	}

	for name, dump := range tempCollections {
		documents := make(map[string]Document)

		for primaryKey, data := range dump.Documents {
			if document, err := MarshalDocument(data); err == nil {
				documents[primaryKey] = *document
			}
		}

		storeCollections[name] = &Collection{
			cfg:       dump.Config,
			Name:      name,
			documents: documents,
		}
	}

	return &Store{collections: storeCollections}, nil

}

func (s *Store) Dump() ([]byte, error) {
	var dumpMap = make(map[string]dumpDTO)

	s.mx.RLock()
	defer s.mx.RUnlock()

	for name, collection := range s.collections {
		var documentsMap = make(map[string]any)
		for _, document := range collection.List() {
			var documentFields = make(map[string]any)

			for fieldKey, fieldData := range document.Fields {
				documentFields[fieldKey] = fieldData.Value
			}

			documentsMap[document.Fields[collection.cfg.PrimaryKey].Value.(string)] = documentFields
		}

		dumpMap[name] = dumpDTO{Documents: documentsMap, Config: collection.cfg}
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
