package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecse"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecse/input"
)

type projectHandler struct {
	projectUsecase projectusecse.ProjectUsecase
}

func NewProjectHandler(projectUsecase projectusecse.ProjectUsecase) *projectHandler {
	return &projectHandler{
		projectUsecase: projectUsecase,
	}
}

func (h *projectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var in input.CreateProject
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		setAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, apperrors.InvalidParameter)
		return
	}

	out, err := h.projectUsecase.CreateProject(r.Context(), &in)
	if err != nil {
		setAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusCreated, out)
}

func (h *projectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["projectID"]

	in := input.UpdateProject{
		ID: projectID,
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		setAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, apperrors.InvalidParameter)
		return
	}

	out, err := h.projectUsecase.UpdateProject(r.Context(), &in)
	if err != nil {
		setAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}
