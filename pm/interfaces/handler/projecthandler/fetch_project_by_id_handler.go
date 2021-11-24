package projecthandler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase"
)

type fetchProjectByIDHandler struct {
	projectUsecase projectusecase.FetchProjectByIDUsecase
}

func NewFetchProjectByIDHandler(fetchProjectByIDUsecase projectusecase.FetchProjectByIDUsecase) *fetchProjectByIDHandler {
	return &fetchProjectByIDHandler{
		projectUsecase: fetchProjectByIDUsecase,
	}
}

func (h *fetchProjectByIDHandler) FetchProjectByID(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["projectID"]

	in := projectusecase.FetchProjectByIDInput{
		ID: projectID,
	}

	out, err := h.projectUsecase.FetchProjectByID(r.Context(), &in)
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}
