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

type deleteProjectNoteHandler struct {
	projectNoteUsecase projectnoteusecase.DeleteProjectNoteUsecase
}

func NewDeleteProjectNoteHandler(deleteProjectNoteUsecase projectnoteusecase.DeleteProjectNoteUsecase) *deleteProjectNoteHandler {
	return &deleteProjectNoteHandler{projectNoteUsecase: deleteProjectNoteUsecase}
}

func (h *deleteProjectNoteHandler) DeleteProjectNote(w http.ResponseWriter, r *http.Request) {
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

	in := projectnoteinput.DeleteProjectNoteInput{
		ID:        projectNoteID,
		ProductID: productID,
		ProjectID: projectID,
	}

	if err := h.projectNoteUsecase.DeleteProjectNote(r.Context(), &in); err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusNoContent, nil)
}
