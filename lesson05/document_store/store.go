package document_store

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
		return nil
	}

	return ErrCollectionNotFound
}
