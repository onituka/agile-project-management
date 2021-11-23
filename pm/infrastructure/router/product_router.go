package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/infrastructure/persistence"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler/producthandler"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase"
	"github.com/onituka/agile-project-management/project-management/usecase/timemanager"
)

func newProductRouter(router *mux.Router, realTime timemanager.TimeManager) {
	productRepository := persistence.NewProductRepository()

	createProductUsecase := productusecase.NewCreateProductUsecase(productRepository, realTime)
	createProductHandler := producthandler.NewCreateProductHandler(createProductUsecase)
	router.HandleFunc("/products", createProductHandler.CreateProduct).Methods(http.MethodPost)

	updateProductUsecase := productusecase.NewUpdateProductUsecase(productRepository, realTime)
	updateProductHandler := producthandler.NewUpdateProductHandler(updateProductUsecase)
	router.HandleFunc("/products/{productID}", updateProductHandler.UpdateProduct).Methods(http.MethodPut)
}
