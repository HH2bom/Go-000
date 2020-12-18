package v1

import (
	"context"
)

type Book struct{}

type CreateBookRequest struct{}

type GetBookRequest struct{}

type ListBookRequest struct{}

type DeleteBookRequest struct{}

type UpdateBookRequest struct{}

type LibraryService interface {
	CreateBook(ctx context.Context, vo CreateBookRequest) (*Book, error)
	GetBook(ctx context.Context, vo GetBookRequest) (*Book, error)
	ListBook(ctx context.Context, vo ListBookRequest) ([]Book, error)
	DeleteBook(ctx context.Context, vo DeleteBookRequest) error
	UpdateBook(ctx context.Context, vo UpdateBookRequest) error
}
