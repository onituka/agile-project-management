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

type createProjectHandler struct {
	projectUsecase projectusecase.CreateProjectUsecase
}

func NewCreateProjectHandler(createProjectUsecase projectusecase.CreateProjectUsecase) *createProjectHandler {
	return &createProjectHandler{
		projectUsecase: createProjectUsecase,
	}
}

func (h *createProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	rv := mux.Vars(r)
	if rv == nil {
		handler.SetAppErrorToCtx(r, apperrors.InternalServerError)
		presenter.ErrorJSON(w, apperrors.InternalServerError)
	}

	productID, ok := rv["productID"]
	if !ok {
		handler.SetAppErrorToCtx(r, apperrors.InternalServerError)
		presenter.ErrorJSON(w, apperrors.InternalServerError)
		return
	}

	in := projectinput.CreateProjectInput{
		ProductID: productID,
	}

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
