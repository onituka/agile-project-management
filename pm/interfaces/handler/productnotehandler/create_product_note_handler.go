package productnotehandler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteinput"
)

type createProductNoteHandler struct {
	productNoteUsecase productnoteusecase.CreateProductNoteUsecase
}

func NewCreateProductNoteHandler(createProductNoteUsecase productnoteusecase.CreateProductNoteUsecase) *createProductNoteHandler {
	return &createProductNoteHandler{
		productNoteUsecase: createProductNoteUsecase,
	}
}

func (h *createProductNoteHandler) CreateProductNote(w http.ResponseWriter, r *http.Request) {
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

	// TODO: 今後JWT claimsからGroupIDを取得する為、現時点ではGroupIDを指定のものとする
	in := productnoteinput.CreateProductNoteInput{
		ProductID: productID,
		GroupID:   "024d78d6-1d03-11ec-a478-0242ac180002",
	}

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, apperrors.InvalidParameter)
		return
	}

	out, err := h.productNoteUsecase.CreateProductNote(r.Context(), &in)
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusCreated, out)
}
