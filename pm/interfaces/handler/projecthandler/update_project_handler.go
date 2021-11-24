package projecthandler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase"
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
	projectID := mux.Vars(r)["projectID"]

	in := projectusecase.UpdateProjectInput{
		ID: projectID,
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
