package users

import (
	"errors"
	"fmt"
	document_store2 "golang-course/lesson_11/internal/document_store"
)

const (
	PrimaryKey          string = "id"
	PrimaryKeyValue     string = "1"
	PrimaryKeyValueTwo  string = "2"
	CollectionUsersName string = "users"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Service struct {
	coll *document_store2.Collection
}

func NewService(store *document_store2.Store) *Service {
	usersCollection, err := store.GetCollection(CollectionUsersName)

	if err != nil {
		err, usersCollection = store.CreateCollection(CollectionUsersName, &document_store2.CollectionConfig{PrimaryKey: PrimaryKey})
		if err != nil {
			fmt.Printf("%v:%s during creation of users service", err, CollectionUsersName)
		}
	}

	service := &Service{
		coll: usersCollection,
	}

	return service
}

func (s *Service) CreateUser(user User) (*User, error) {
	if _, err := s.coll.Get(user.ID); err == nil {
		return nil, ErrUserAlreadyExists
	}

	doc, err := document_store2.MarshalDocument(user)

	if err != nil {
		return nil, err
	}

	s.coll.Put(*doc)

	return &user, nil
}

func (s *Service) ListUsers() ([]User, error) {
	docs := s.coll.List()

	if len(docs) == 0 {
		return nil, nil
	}

	users := make([]User, 0, len(docs))

	for _, doc := range docs {
		tempUser := &User{}
		err := document_store2.UnmarshalDocument(&doc, tempUser)

		if err != nil {
			return nil, err
		}

		users = append(users, *tempUser)
	}

	return users, nil
}

func (s *Service) GetUser(userID string) (*User, error) {
	doc, err := s.coll.Get(userID)
	if err != nil {
		return nil, fmt.Errorf("%w by id%s", ErrUserNotFound, userID)
	}

	user := &User{}
	if err := document_store2.UnmarshalDocument(doc, user); err != nil {
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
