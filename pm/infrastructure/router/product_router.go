package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb/persistence"
	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb/query"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler/producthandler"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase"
)

func newProductRouter(router *mux.Router) {
	productRepository := persistence.NewProductRepository()
	productQueryService := query.NewProductQueryServiceImpl()

	createProductUsecase := productusecase.NewCreateProductUsecase(productRepository)
	createProductHandler := producthandler.NewCreateProductHandler(createProductUsecase)
	router.HandleFunc("/products", createProductHandler.CreateProduct).Methods(http.MethodPost)

	updateProductUsecase := productusecase.NewUpdateProductUsecase(productRepository)
	updateProductHandler := producthandler.NewUpdateProductHandler(updateProductUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}", updateProductHandler.UpdateProduct).Methods(http.MethodPut)

	fetchProductByIDUsecase := productusecase.NewFetchProductByIDUsecase(productRepository)
	fetchProductByIDHandler := producthandler.NewFetchProductByIDHandler(fetchProductByIDUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}", fetchProductByIDHandler.FetchProductByID).Methods(http.MethodGet)

	fetchProductsUsecase := productusecase.NewFetchProductsUsecase(productQueryService)
	fetchProductsHandler := producthandler.NewFetchProductsHandler(fetchProductsUsecase)
	router.HandleFunc("/products", fetchProductsHandler.FetchProducts).Queries("page", "{page}", "limit", "{limit}").Methods(http.MethodGet)

	searchProductsUsecase := productusecase.NewSearchProductsUsecase(productQueryService)
	searchProductsHandler := producthandler.NewSearchProductsHandler(searchProductsUsecase)
	router.HandleFunc("/products/search", searchProductsHandler.SearchProducts).Queries("name", "{name}", "page", "{page}", "limit", "{limit}").Methods(http.MethodGet)
}
