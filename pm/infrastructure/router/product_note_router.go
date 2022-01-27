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
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/notes", createProductNoteHandler.CreateProductNote).Methods(http.MethodPost)

	updateProductNoteUsecase := productnoteusecase.NewUpdateProductNoteUsecase(productNoteRepository, productRepository)
	updateProductNoteHandler := productnotehandler.NewUpdateProductNoteHandler(updateProductNoteUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/notes/{productNoteID:[a-z0-9-]{36}}", updateProductNoteHandler.UpdateProductNote).Methods(http.MethodPut)

	fetchProductNoteByIDUsecase := productnoteusecase.NewFetchProductNoteByIDUsecase(productNoteRepository, productRepository)
	fetchProductNoteByIDHandler := productnotehandler.NewFetchProductNoteByIDHandler(fetchProductNoteByIDUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/notes/{productNoteID:[a-z0-9-]{36}}", fetchProductNoteByIDHandler.FetchProductNoteByID).Methods(http.MethodGet)

	fetchProductNotesUsecase := productnoteusecase.NewFetchProductNotesUsecase(productNoteQueryService, productRepository)
	fetchProductNotesHandler := productnotehandler.NewFetchProductNotesHandler(fetchProductNotesUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/notes", fetchProductNotesHandler.FetchProductNotes).Queries("page", "{page}", "limit", "{limit}").Methods(http.MethodGet)

	searchProductNotesUsecase := productnoteusecase.NewSearchProductNotesUsecase(productNoteQueryService, productRepository)
	searchProductNotesHandler := productnotehandler.NewSearchProductNotesHandler(searchProductNotesUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/notes/search", searchProductNotesHandler.SearchProductNotes).Queries("page", "{page}", "limit", "{limit}").Methods(http.MethodGet)

	deleteProductNoteUsecase := productnoteusecase.NewDeleteProductNoteUsecase(productNoteRepository, productRepository)
	deleteProductNoteHandler := productnotehandler.NewDeleteProductNoteHandler(deleteProductNoteUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/notes/{productNoteID:[a-z0-9-]{36}}", deleteProductNoteHandler.DeleteProductNote).Methods(http.MethodDelete)
}
