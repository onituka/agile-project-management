package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb/persistence"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler/producthandler"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase"
)

func newProductRouter(router *mux.Router) {
	productRepository := persistence.NewProductRepository()

	createProductUsecase := productusecase.NewCreateProductUsecase(productRepository)
	createProductHandler := producthandler.NewCreateProductHandler(createProductUsecase)
	router.HandleFunc("/products", createProductHandler.CreateProduct).Methods(http.MethodPost)

	updateProductUsecase := productusecase.NewUpdateProductUsecase(productRepository)
	updateProductHandler := producthandler.NewUpdateProductHandler(updateProductUsecase)
	router.HandleFunc("/products/{productID}", updateProductHandler.UpdateProduct).Methods(http.MethodPut)

	fetchProductByIDUsecase := productusecase.NewFetchProductByIDUsecase(productRepository)
	fetchProductByIDHandler := producthandler.NewFetchProductByIDHandler(fetchProductByIDUsecase)
	router.HandleFunc("/products/{productID}", fetchProductByIDHandler.FetchProductByID).Methods(http.MethodGet)

	fetchProductsUsecase := productusecase.NewFetchProductsUsecase(productRepository)
	fetchProductsHandler := producthandler.NewFetchProductsHandler(fetchProductsUsecase)
	router.HandleFunc("/products", fetchProductsHandler.FetchProducts).Methods(http.MethodGet)
}
