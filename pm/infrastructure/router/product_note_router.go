package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb/persistence"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler/productnotehandler"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase"
)

func newProductNoteRouter(router *mux.Router) {
	productNoteRepository := persistence.NewProductNoteRepository()
	productRepository := persistence.NewProductRepository()

	createProductNoteUsecase := productnoteusecase.NewCreateProductNoteUsecase(productNoteRepository, productRepository)
	createProductNoteHandler := productnotehandler.NewCreateProductNoteHandler(createProductNoteUsecase)
	router.HandleFunc("/products/{productID}/notes", createProductNoteHandler.CreateProductNote).Methods(http.MethodPost)
}
