package projecthandler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectinput"
)

type updateProjectHandler struct {
	projectUsecase projectusecase.UpdateProjectUsecase
}

func NewUpdateProjectHandler(updateProjectUsecase projectusecase.UpdateProjectUsecase) *updateProjectHandler {
	return &updateProjectHandler{
		projectUsecase: updateProjectUsecase,
	}
}

func (h *updateProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
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

	in := projectinput.UpdateProjectInput{
		ID:        projectID,
		ProductID: productID,
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, apperrors.InvalidParameter)
		return
	}

	out, err := h.projectUsecase.UpdateProject(r.Context(), &in)
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}
