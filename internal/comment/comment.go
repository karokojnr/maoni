package comment

import (
	"context"
	"fmt"
)

// Comment - Represents the comment structure
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store -
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
		return Comment{}, err
	}
	return cmt, nil
}
