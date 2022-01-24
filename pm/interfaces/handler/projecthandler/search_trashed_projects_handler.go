package projecthandler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectinput"
)

type searchTrashedProjectsHandler struct {
	projectUsecase projectusecase.SearchTrashedProjectsUsecase
}

func NewSearchTrashedProjectsHandler(searchTrashedProjectsUsecase projectusecase.SearchTrashedProjectsUsecase) *searchTrashedProjectsHandler {
	return &searchTrashedProjectsHandler{
		projectUsecase: searchTrashedProjectsUsecase,
	}
}

func (h *searchTrashedProjectsHandler) SearchTrashedProjects(w http.ResponseWriter, r *http.Request) {
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

	keyword := r.URL.Query().Get("keyword")

	page, err := strconv.ParseUint(r.URL.Query().Get("page"), 10, 32)
	if err != nil {
		handler.SetAppErrorToCtx(r, apperrors.InvalidParameter)
		presenter.ErrorJSON(w, apperrors.InvalidParameter)
		return
	}

	limit, err := strconv.ParseUint(r.URL.Query().Get("limit"), 10, 32)
	if err != nil {
		handler.SetAppErrorToCtx(r, apperrors.InvalidParameter)
		presenter.ErrorJSON(w, apperrors.InvalidParameter)
		return
	}

	in := projectinput.SearchTrashedProjectsInput{
		ProductID: productID,
		KeyWord:   keyword,
		Page:      uint32(page),
		Limit:     uint32(limit),
	}

	out, err := h.projectUsecase.SearchTrashedProjects(r.Context(), &in)
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}
