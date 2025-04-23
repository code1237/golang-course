package documentstore

type Store struct {
	collections map[string]*Collection
}

func NewStore() *Store {
	store := &Store{
		collections: make(map[string]*Collection),
	}

	return store
}

func (s *Store) CreateCollection(name string, cfg *CollectionConfig) (bool, *Collection) {
	if len(name) == 0 {
		panic("Collection name must not be empty")
	}

	if _, ok := s.collections[name]; ok {
		return false, nil
	}

	newCollection := &Collection{
		*cfg,
		name,
		make(map[string]Document),
	}

	s.collections[name] = newCollection

	return true, newCollection
}

func (s *Store) GetCollection(name string) (*Collection, bool) {
	collection, ok := s.collections[name]
	return collection, ok
}

func (s *Store) DeleteCollection(name string) bool {
	if _, ok := s.collections[name]; ok {
		delete(s.collections, name)
		return true
	}

	return false
}
