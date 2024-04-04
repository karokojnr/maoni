package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

// Comment - Represents the comment structure
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface defines all the methods
// that the service needs to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service - struct on which all the logic will be built on top of
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new Service
func NewService(store Store) *Service {
	return &Service{
		store,
	}
}

// GetComment - Gets a comment by id
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving a comment by its id")

	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
