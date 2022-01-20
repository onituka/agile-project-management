package projecthandler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectinput"
)

type restoreFromTrashProjectHandler struct {
	projectUsecase projectusecase.RestoreFromTrashProjectUsecase
}

func NewRestoreFromTrashProjectHandler(restoreFromTrashProjectUsecase projectusecase.RestoreFromTrashProjectUsecase) *restoreFromTrashProjectHandler {
	return &restoreFromTrashProjectHandler{
		projectUsecase: restoreFromTrashProjectUsecase,
	}
}

func (h *restoreFromTrashProjectHandler) RestoreFromTrashProject(w http.ResponseWriter, r *http.Request) {
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

	in := projectinput.RestoreFromTrashProjectIDInput{
		ID:        projectID,
		ProductID: productID,
	}

	out, err := h.projectUsecase.RestoreFromTrashProject(r.Context(), &in)
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}
