package users

import (
	"errors"
	"golang-course/lesson05/document_store"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

const (
	PrimaryKey          string = "id"
	PrimaryKeyValue     string = "1"
	CollectionUsersName string = "users"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Service struct {
	coll *document_store.Collection
}

func NewService(store *document_store.Store) *Service {
	if coll, err := store.GetCollection(CollectionUsersName); err == nil {
		service := &Service{
			coll: coll,
		}

		return service
	}

	return nil
}

func (s *Service) CreateUser(user User) (*User, error) {
	if _, err := s.coll.Get(user.ID); err == nil {
		return nil, ErrUserAlreadyExists
	}

	doc, err := document_store.MarshalDocument(user)

	if err != nil {
		return nil, err
	}

	s.coll.Put(*doc)

	return &user, nil
}

//func (s *Service) ListUsers() ([]User, error) {
//	// ...
//}

func (s *Service) GetUser(userID string) (*User, error) {
	doc, err := s.coll.Get(userID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	user := &User{}
	if err := document_store.UnmarshalDocument(doc, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) DeleteUser(userID string) error {
	if !s.coll.Delete(userID) {
		return ErrUserNotFound
	}

	return nil
}
