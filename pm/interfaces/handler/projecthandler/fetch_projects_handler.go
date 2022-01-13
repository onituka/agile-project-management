package projecthandler

import (
	"net/http"
	"strconv"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectinput"
)

type fetchProjectsHandler struct {
	projectUsecase projectusecase.FetchProjectsUsecase
}

func NewFetchProjectsHandler(fetchProjectsUsecase projectusecase.FetchProjectsUsecase) *fetchProjectsHandler {
	return &fetchProjectsHandler{
		projectUsecase: fetchProjectsUsecase,
	}
}

func (h *fetchProjectsHandler) FetchProjects(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.ParseUint(r.URL.Query().Get("page"), 10, 32)
	if err != nil {
		handler.SetAppErrorToCtx(r, apperrors.InternalServerError)
		presenter.ErrorJSON(w, apperrors.InternalServerError)
		return
	}

	limit, err := strconv.ParseUint(r.URL.Query().Get("limit"), 10, 32)
	if err != nil {
		handler.SetAppErrorToCtx(r, apperrors.InternalServerError)
		presenter.ErrorJSON(w, apperrors.InternalServerError)
		return
	}
	//TODO: 今後JWTにてProductIDを認証する為、現時点ではProductIDを指定のものとする
	in := projectinput.FetchProjectsInput{
		ProductID: "4495c574-34c2-4fb3-9ca4-3a7c79c267a6",
		Page:      uint32(page),
		Limit:     uint32(limit),
	}

	out, err := h.projectUsecase.FetchProjects(r.Context(), &in)
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}
