package v1

import (
	"context"

	"app/internal/app-server/api/library/v1"
	"app/internal/pkg/appcontext"
)

type LibraryService struct {
	appCtx appcontext.Context
}

func NewLibraryService(appCtx appcontext.Context) *LibraryService {
	return &LibraryService{appCtx: appCtx}
}

func (*LibraryService) CreateBook(ctx context.Context, vo v1.CreateBookRequest) (*v1.Book, error) {
	panic("implement me")
}

func (*LibraryService) GetBook(ctx context.Context, vo v1.GetBookRequest) (*v1.Book, error) {
	panic("implement me")
}

func (*LibraryService) ListBook(ctx context.Context, vo v1.ListBookRequest) ([]v1.Book, error) {
	panic("implement me")
}

func (*LibraryService) DeleteBook(ctx context.Context, vo v1.DeleteBookRequest) error {
	panic("implement me")
}

func (*LibraryService) UpdateBook(ctx context.Context, vo v1.UpdateBookRequest) error {
	panic("implement me")
}
