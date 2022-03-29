package comment

import (
	"context"
	"errors"
	"fmt"
)

// Errors  - defining errors here is better than allowing the error to directly pass through on a api call,
// which could expose implementation detials.

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented = errors.New("not implemented")
)

// Comment - a representation of the comment
// structure for our service
type Comment struct {
	ID string
	Slug string
	Body string
	Author string
}

// Store - this interface defines all the methods
// that the service needs to operate.
// 1. The benefit of this structure is that the storage interface can 
// reach out to postgress, kasandera, etc
// 2. the store can be mocked for unit testing
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service - is the struct on which 
// all the logic will be built on top of
type Service struct{
	Store Store
	// Serivice/Store - Store with a type of Store
	// when a service is instantiated, and we pass in the repository layer, it must match the interface Store
}

// NewService - returns a pointer to a 
// new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	// func explained below (1) (2) (3) 
	// (1st) -> create a method that takes in a pointer reciever
	// (2nd) -> it will getComment with a first argument of context and second as id
	// (3rd) -> it will return a comment or an error


	fmt.Println("retrieving a comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err) // send to a log activation system such as datadog.
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

func (s *Service) updateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}


func (s *Service) deleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) createComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}

