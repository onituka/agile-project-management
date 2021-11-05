package persistence

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/config"
)

func execFromCtx(ctx context.Context) (sqlx.ExtContext, error) {
	val := ctx.Value(config.DBKey)
	if val == nil {
		return nil, apperrors.InternalServerError
	}

	conn, ok := val.(sqlx.ExtContext)
	if !ok {
		return nil, apperrors.InternalServerError
	}

	return conn, nil
}
