package comment

import (
	"context"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
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
type CommentStore interface {
	PostComment(context.Context, Comment) (Comment, error)
	GetComment(context.Context, string) (Comment, error)
	UpdateComment(context.Context, string, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	Ping(context.Context) error
}

// Service - struct on which all the logic will be built on top of
type Service struct {
	Store CommentStore
}

// NewService - returns a pointer to a new Service
func NewService(store CommentStore) *Service {
	return &Service{
		store,
	}
}

// PostComment - adds a new comment to the database
func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	cmt, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}
	return cmt, nil
}

// GetComment - Gets a comment by id
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

// UpdateComment - updates a comment by ID with new comment info
func (s *Service) UpdateComment(
	ctx context.Context, id string, comment Comment,
) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, id, comment)
	if err != nil {
		log.Errorf("an error occurred updating the comment: %s", err.Error())
	}
	return cmt, nil
}

// DeleteComment - deletes a comment from the database by ID
func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.Store.DeleteComment(ctx, id)
}

// ReadyCheck - a function that tests we are functionally ready to serve requests
func (s *Service) ReadyCheck(ctx context.Context) error {
	log.Info("Checking readiness")
	return s.Store.Ping(ctx)
}
