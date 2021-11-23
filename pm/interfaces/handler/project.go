package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase"
)

type projectHandler struct {
	projectUsecase projectusecase.ProjectUsecase
}

func NewProjectHandler(projectUsecase projectusecase.ProjectUsecase) *projectHandler {
	return &projectHandler{
		projectUsecase: projectUsecase,
	}
}

func (h *projectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var in projectusecase.CreateProjectInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, apperrors.InvalidParameter)
		return
	}

	out, err := h.projectUsecase.CreateProject(r.Context(), &in)
	if err != nil {
		SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusCreated, out)
}

func (h *projectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["projectID"]

	in := projectusecase.UpdateProjectInput{
		ID: projectID,
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, apperrors.InvalidParameter)
		return
	}

	out, err := h.projectUsecase.UpdateProject(r.Context(), &in)
	if err != nil {
		SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}

func (h *projectHandler) FetchProjectByID(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["projectID"]

	in := projectusecase.FetchProjectByIDInput{
		ID: projectID,
	}

	out, err := h.projectUsecase.FetchProjectByID(r.Context(), &in)
	if err != nil {
		SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}

func (h *projectHandler) FetchProjects(w http.ResponseWriter, r *http.Request) {
	out, err := h.projectUsecase.FetchProjects(r.Context())
	if err != nil {
		SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}
