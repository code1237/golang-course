package document_store

import "errors"

var (
	ErrCollectionAlreadyExists = errors.New("collection already exists")
	ErrCollectionNotFound      = errors.New("collection not found")
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
