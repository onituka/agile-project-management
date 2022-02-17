package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb/persistence"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler/productnotecommenthandler"
	"github.com/onituka/agile-project-management/project-management/usecase/productnotecommentusecase"
)

func newProductNoteCommentRouter(router *mux.Router) {
	productNoteCommentRepository := persistence.NewProductNoteCommentRepository()
	productNoteRepository := persistence.NewProductNoteRepository()
	productRepository := persistence.NewProductRepository()

	createProductNoteCommentUsecase := productnotecommentusecase.NewCreateProductNoteCommentUsecase(productNoteCommentRepository, productNoteRepository, productRepository)
	createProductNoteCommentHandler := productnotecommenthandler.NewCreateProductNoteCommentHandler(createProductNoteCommentUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/productnotes/{productNoteID:[a-z0-9-]{36}}/comments", createProductNoteCommentHandler.CreateProductNoteComment).Methods(http.MethodPost)
}
