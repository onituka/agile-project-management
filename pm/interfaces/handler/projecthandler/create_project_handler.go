package projecthandler

import (
	"encoding/json"
	"net/http"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectinput"
)

type createProjectHandler struct {
	projectUsecase projectusecase.CreateProjectUsecase
}

func NewCreateProjectHandler(createProjectUsecase projectusecase.CreateProjectUsecase) *createProjectHandler {
	return &createProjectHandler{
		projectUsecase: createProjectUsecase,
	}
}

func (h *createProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var in projectinput.CreateProjectInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, apperrors.InvalidParameter)
		return
	}

	out, err := h.projectUsecase.CreateProject(r.Context(), &in)
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusCreated, out)
}
