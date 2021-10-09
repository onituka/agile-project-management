package handler

import (
	"encoding/json"
	"net/http"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase"
	"github.com/onituka/agile-project-management/project-management/usecase/input"
)

type projectHandler struct {
	projectUsecase usecase.ProjectUsecase
}

func NewProjectHandler(projectUsecase usecase.ProjectUsecase) *projectHandler {
	return &projectHandler{
		projectUsecase: projectUsecase,
	}
}

func (h *projectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var in input.Project
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		presenter.ErrorJSON(w, apperrors.InvalidParameter)
		return
	}

	out, err := h.projectUsecase.CreateProject(&in)
	if err != nil {
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusCreated, out)
}
