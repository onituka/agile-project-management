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

type searchProjectsHandler struct {
	projectUsecase projectusecase.SearchProjectsUsecase
}

func NewSearchProjectsHandler(searchProjectsUsecase projectusecase.SearchProjectsUsecase) *searchProjectsHandler {
	return &searchProjectsHandler{
		projectUsecase: searchProjectsUsecase,
	}
}

func (h *searchProjectsHandler) SearchProjects(w http.ResponseWriter, r *http.Request) {
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

	keyWord := r.URL.Query().Get("keyWord")

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

	in := projectinput.SearchProjectsInput{
		ProductID: productID,
		KeyWord:   keyWord,
		Page:      uint32(page),
		Limit:     uint32(limit),
	}

	out, err := h.projectUsecase.SearchProjects(r.Context(), &in)
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}
