package document_store

import (
	"errors"
	"log/slog"
)

var (
	ErrCollectionAlreadyExists   = errors.New("collection already exists")
	ErrCollectionNotFound        = errors.New("collection not found")
	ErrCollectionNameCantBeEmpty = errors.New("collection name can not be empty")
)

type Collection struct {
	cfg       CollectionConfig
	Name      string
	documents map[string]Document
}

type CollectionConfig struct {
	PrimaryKey string
}

func (s *Collection) Put(doc Document) {
	field, ok := doc.Fields[s.cfg.PrimaryKey]

	if !ok {
		return
	}

	if field.Type != DocumentFieldTypeString {
		return
	}

	valueAsString, ok := field.Value.(string)
	if !ok || len(valueAsString) == 0 {
		return
	}

	s.documents[valueAsString] = doc
	slog.Info("New document was added", "collection", s.Name, "id", valueAsString)
	return
}

func (s *Collection) Get(key string) (*Document, error) {
	if doc, ok := s.documents[key]; ok {
		return &doc, nil
	}

	return nil, ErrDocumentNotFound
}

func (s *Collection) Delete(key string) bool {
	if _, ok := s.documents[key]; ok {
		delete(s.documents, key)
		slog.Info("Document was deleted", "collection", s.Name, "id", key)
		return true
	}

	return false
}

func (s *Collection) List() []Document {
	documentsSlice := make([]Document, 0, len(s.documents))

	for _, doc := range s.documents {
		documentsSlice = append(documentsSlice, doc)
	}

	return documentsSlice
}
