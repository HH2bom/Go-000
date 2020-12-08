package service

import (
	"context"
	"fmt"

	"Go-000/week02/01/dao"
	pkgerrors "github.com/pkg/errors"
)

func DoSomething(ctx context.Context, args ...interface{}) (interface{}, error) {
	data, err := dao.SearchRecordByID(ctx, nil)
	if err != nil {
		// 返回异常附带 stack
		if pkgerrors.Is(err, dao.ErrNotFound) {
			return nil, pkgerrors.WithStack(&Err{notFoundCode, fmt.Sprintf("args is %v", args), err})
		}

		return nil, pkgerrors.WithStack(&Err{xxxxxxxxCode, "something wrong", err})
	}

	return data, nil
}
