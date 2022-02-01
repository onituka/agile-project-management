package productnotecommenthandler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/productnotecommentusecase"
	"github.com/onituka/agile-project-management/project-management/usecase/productnotecommentusecase/productnotecommentinput"
)

type createProductNoteCommentHandler struct {
	productNoteCommentUsecase productnotecommentusecase.CreateProductNoteCommentUsecase
}

func NewCreateProductNoteCommentHandler(createProductNoteUsecase productnotecommentusecase.CreateProductNoteCommentUsecase) *createProductNoteCommentHandler {
	return &createProductNoteCommentHandler{
		productNoteCommentUsecase: createProductNoteUsecase,
	}
}

func (h *createProductNoteCommentHandler) CreateProductNoteComment(w http.ResponseWriter, r *http.Request) {
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

	productNoteID, ok := rv["productNoteID"]
	if !ok {
		handler.SetAppErrorToCtx(r, apperrors.InternalServerError)
		presenter.ErrorJSON(w, apperrors.InternalServerError)
		return
	}

	// TODO: 今後JWT claimsからGroupID,UserIDを取得する為、現時点ではGroupID,UserIDを指定のものとする
	in := productnotecommentinput.CreateProductNoteCommentInput{
		ProductID:     productID,
		ProductNoteID: productNoteID,
		GroupID:       "024d78d6-1d03-41ec-a478-0242ac180002",
		UserID:        "024d78d6-1d03-41ec-a478-0242ac184402",
	}

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, apperrors.InvalidParameter)
		return
	}

	out, err := h.productNoteCommentUsecase.CreateProductNoteComment(r.Context(), &in)
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusCreated, out)
}
