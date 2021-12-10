package projecthandler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
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
	rv := mux.Vars(r)
	if rv == nil {
		handler.SetAppErrorToCtx(r, apperrors.InternalServerError)
		presenter.ErrorJSON(w, apperrors.InternalServerError)
		return
	}
	projectID, ok := rv["projectID"]
	if !ok {
		handler.SetAppErrorToCtx(r, apperrors.InternalServerError)
		presenter.ErrorJSON(w, apperrors.InternalServerError)
		return
	}

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
