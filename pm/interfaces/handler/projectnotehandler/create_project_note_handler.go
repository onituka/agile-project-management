package projectnotehandler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteinput"
)

type createProjectNoteHandler struct {
	projectNoteUsecase projectnoteusecase.CreateProjectNoteUsecase
}

func NewCreateProjectNoteHandler(createProjectNoteUsecase projectnoteusecase.CreateProjectNoteUsecase) *createProjectNoteHandler {
	return &createProjectNoteHandler{projectNoteUsecase: createProjectNoteUsecase}
}

func (h *createProjectNoteHandler) CreateProjectNote(w http.ResponseWriter, r *http.Request) {
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

	projectID, ok := rv["projectID"]
	if !ok {
		handler.SetAppErrorToCtx(r, apperrors.InternalServerError)
		presenter.ErrorJSON(w, apperrors.InternalServerError)
		return
	}

	// TODO: 今後JWT claimsからGroupID,UserIDを取得する為、現時点ではGroupID,UserIDを指定のものとする
	in := projectnoteinput.CreateProjectNoteInput{
		ProductID: productID,
		ProjectID: projectID,
		GroupID:   "024d78d6-1d03-41ec-a478-0242ac180002",
		UserID:    "777d78d6-1d03-11ec-a478-0242ac184402",
	}

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
	}

	out, err := h.projectNoteUsecase.CreateProjectNote(r.Context(), &in)
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}
