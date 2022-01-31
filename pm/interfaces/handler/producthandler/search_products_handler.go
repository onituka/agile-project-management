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

type searchProductsHandler struct {
	productUsecase productusecase.SearchProductsUsecase
}

func NewSearchProductsHandler(searchProductsUsecase productusecase.SearchProductsUsecase) *searchProductsHandler {
	return &searchProductsHandler{
		productUsecase: searchProductsUsecase,
	}
}

func (h *searchProductsHandler) SearchProducts(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

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

	// TODO: 今後JWTにてGroupIDを認証する為、現時点ではGroupIDを指定のものとする
	in := &productinput.SearchProductsInput{
		GroupID:     "024d78d6-1d03-41ec-a478-0242ac180002",
		ProductName: name,
		Page:        uint32(page),
		Limit:       uint32(limit),
	}

	out, err := h.productUsecase.SearchProducts(r.Context(), in)
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}
