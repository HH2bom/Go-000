package dao

import (
	"context"
	"database/sql"
	"errors"

	pkgerrors "github.com/pkg/errors"
)

var ErrNotFound  = errors.New("data not found")

type Record struct{}

func SearchRecordByID(ctx context.Context,id interface{}) (*Record, error) {
	_, err := scan()
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, pkgerrors.WithStack(ErrNotFound)
		default:
			return nil, pkgerrors.WithStack(err)
		}
	}
	return &Record{}, nil
}

func scan() (interface{}, error) {
	return nil, sql.ErrNoRows
}
