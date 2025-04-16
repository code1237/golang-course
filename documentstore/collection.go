package documentstore

type Collection struct {
	*CollectionConfig
	Name      string
	documents map[string]Document
}

type CollectionConfig struct {
	PrimaryKey string
}

func (s *Collection) Put(doc Document) {
	for key, field := range doc.Fields {
		if key != s.PrimaryKey {
			continue
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
}

func (s *Collection) Get(key string) (*Document, bool) {
	if doc, ok := s.documents[key]; ok {
		return &doc, true
	}

	return nil, false
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
