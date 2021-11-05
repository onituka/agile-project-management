package handler

import (
	"context"
	"net/http"

	"github.com/onituka/agile-project-management/project-management/config"
)

func setAppErrorToCtx(r *http.Request, err error) {
	ctx := context.WithValue(r.Context(), config.AppErrorKey, err)
	*r = *r.WithContext(ctx)
}
