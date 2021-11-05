package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/config"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
)

func DBMiddlewareFunc(conn *sqlx.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var tx *sqlx.Tx
			var err error
			ctx := r.Context()

			if isTransactionMethod(r.Method) {
				tx, err = conn.Beginx()
				if err != nil {
					log.Printf("failed to begin transaction: %+v", err)
					presenter.ErrorJSON(w, apperrors.InternalServerError)
					return
				}

				ctx = context.WithValue(ctx, config.DBKey, tx)
			} else {
				ctx = context.WithValue(ctx, config.DBKey, conn)
			}

			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)

			if err = getAppErrorFromCtx(r.Context()); err != nil {
				if tx != nil {
					err = tx.Rollback()
					log.Printf("failed to handle transaction Rollback: %+v", err)
				}

				return
			}

			if tx != nil {
				if err = tx.Commit(); err != nil {
					log.Printf("failed to handle transaction commit: %+v", err)
					return
				}
			}
		})
	}
}

func isTransactionMethod(method string) bool {
	if method == http.MethodPost ||
		method == http.MethodPut ||
		method == http.MethodDelete {
		return true
	}

	return false
}
