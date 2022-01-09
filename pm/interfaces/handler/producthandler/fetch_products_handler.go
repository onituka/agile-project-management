package producthandler

import (
	"net/http"
	"strconv"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productinput"
)

type fetchProductsHandler struct {
	productUsecase productusecase.FetchProductsUsecase
}

func NewFetchProductsHandler(fetchProductsUsecase productusecase.FetchProductsUsecase) *fetchProductsHandler {
	return &fetchProductsHandler{
		productUsecase: fetchProductsUsecase,
	}
}

func (h *fetchProductsHandler) FetchProducts(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		handler.SetAppErrorToCtx(r, apperrors.InvalidParameter)
		presenter.ErrorJSON(w, apperrors.InvalidParameter)
		return
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		handler.SetAppErrorToCtx(r, apperrors.InvalidParameter)
		presenter.ErrorJSON(w, apperrors.InvalidParameter)
		return
	}

	// TODO: 今後JWTにてGroupIDを認証する為、現時点ではGroupIDを指定のものとする
	in := &productinput.FetchProductsInput{
		GroupID: "024d78d6-1d03-11ec-a478-0242ac180002",
		Page:    page,
		Limit:   limit,
	}

	out, err := h.productUsecase.FetchProducts(r.Context(), in)
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}
