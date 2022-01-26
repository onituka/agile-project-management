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

type updateProjectNoteHandler struct {
	projectNoteUsecase projectnoteusecase.UpdateProjectNoteUsecase
}

func NewUpdateProjectNoteHandler(updateProjectNoteUsecase projectnoteusecase.UpdateProjectNoteUsecase) *updateProjectNoteHandler {
	return &updateProjectNoteHandler{projectNoteUsecase: updateProjectNoteUsecase}
}

func (h *updateProjectNoteHandler) UpdateProjectNote(w http.ResponseWriter, r *http.Request) {
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

	projectID, ok := rv["projectID"]
	if !ok {
		handler.SetAppErrorToCtx(r, apperrors.InternalServerError)
		presenter.ErrorJSON(w, apperrors.InternalServerError)
		return
	}

	projectNoteID, ok := rv["projectNoteID"]
	if !ok {
		handler.SetAppErrorToCtx(r, apperrors.InternalServerError)
		presenter.ErrorJSON(w, apperrors.InternalServerError)
		return
	}

	// TODO: 今後JWT claimsからUserIDを取得する為、現時点ではUserIDを指定のものとする
	in := projectnoteinput.UpdateProjectNoteInput{
		ID:        projectNoteID,
		ProductID: productID,
		ProjectID: projectID,
		UserID:    "777d78d6-1d03-11ec-a478-0242ac184402",
	}

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	out, err := h.projectNoteUsecase.UpdateProjectNote(r.Context(), &in)
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}
