package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb/persistence"
	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb/query"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler/productnotehandler"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase"
)

func newProductNoteRouter(router *mux.Router) {
	productNoteRepository := persistence.NewProductNoteRepository()
	productRepository := persistence.NewProductRepository()
	productNoteQueryService := query.NewProductNoteQueryServiceImpl()

	createProductNoteUsecase := productnoteusecase.NewCreateProductNoteUsecase(productNoteRepository, productRepository)
	createProductNoteHandler := productnotehandler.NewCreateProductNoteHandler(createProductNoteUsecase)
	router.HandleFunc("/products/{productID}/notes", createProductNoteHandler.CreateProductNote).Methods(http.MethodPost)

	updateProductNoteUsecase := productnoteusecase.NewUpdateProductNoteUsecase(productNoteRepository, productRepository)
	updateProductNoteHandler := productnotehandler.NewUpdateProductNoteHandler(updateProductNoteUsecase)
	router.HandleFunc("/products/{productID}/notes/{productNoteID}", updateProductNoteHandler.UpdateProductNote).Methods(http.MethodPut)

	fetchProductNotesUsecase := productnoteusecase.NewFetchProductNotesUsecase(productNoteQueryService, productRepository)
	fetchProductNotesHandler := productnotehandler.NewFetchProductNotesHandler(fetchProductNotesUsecase)
	router.HandleFunc("/products/{productID}/notes", fetchProductNotesHandler.FetchProductNotes).Queries("page", "{page}", "limit", "{limit}").Methods(http.MethodGet)
}
