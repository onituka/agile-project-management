package projectnotehandler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteinput"
)

type fetchProjectNoteByIDHandler struct {
	projectNoteUsecase projectnoteusecase.FetchProjectNoteByIDUsecase
}

func NewFetchProjectNoteByIDHandler(fetchProjectNoteByIDUsecase projectnoteusecase.FetchProjectNoteByIDUsecase) *fetchProjectNoteByIDHandler {
	return &fetchProjectNoteByIDHandler{projectNoteUsecase: fetchProjectNoteByIDUsecase}
}

func (h *fetchProjectNoteByIDHandler) FetchProjectNoteByID(w http.ResponseWriter, r *http.Request) {
	rv := mux.Vars(r)
	if rv == nil {
		handler.SetAppErrorToCtx(r, apperrors.InternalServerError)
		presenter.ErrorJSON(w, apperrors.InternalServerError)
		return
	}

	productID, ok := rv["productID"]
	if !ok {
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

	projectNoteID, ok := rv["projectNoteID"]
	if !ok {
		handler.SetAppErrorToCtx(r, apperrors.InternalServerError)
		presenter.ErrorJSON(w, apperrors.InternalServerError)
		return
	}

	in := projectnoteinput.FetchProjectNoteByIDInput{
		ID:        projectNoteID,
		ProductID: productID,
		ProjectID: projectID,
	}

	out, err := h.projectNoteUsecase.FetchProjectNoteByID(r.Context(), &in)
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}
