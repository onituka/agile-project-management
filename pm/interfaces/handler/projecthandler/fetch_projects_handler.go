package projecthandler

import (
	"net/http"

	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase"
)

type fetchProjectsHandler struct {
	projectUsecase projectusecase.FetchProjectsUsecase
}

func NewFetchProjectsHandler(fetchProjectsUsecase projectusecase.FetchProjectsUsecase) *fetchProjectsHandler {
	return &fetchProjectsHandler{
		projectUsecase: fetchProjectsUsecase,
	}
}

func (h *fetchProjectsHandler) FetchProjects(w http.ResponseWriter, r *http.Request) {
	out, err := h.projectUsecase.FetchProjects(r.Context())
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}
