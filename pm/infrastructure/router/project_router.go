package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb/persistence"
	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb/query"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler/projecthandler"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase"
)

func newProjectRouter(router *mux.Router) {
	projectRepository := persistence.NewProjectRepository()
	projectQuery := query.NewProjectQuery()
	productRepository := persistence.NewProductRepository()

	createProjectUsecase := projectusecase.NewCreateProjectUsecase(projectRepository, productRepository)
	createProjectHandler := projecthandler.NewCreateProjectHandler(createProjectUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/projects", createProjectHandler.CreateProject).Methods(http.MethodPost)

	updateProjectUsecase := projectusecase.NewUpdateProjectUsecase(projectRepository, productRepository)
	updateProjectHandler := projecthandler.NewUpdateProjectHandler(updateProjectUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/projects/{projectID:[a-z0-9-]{36}}", updateProjectHandler.UpdateProject).Methods(http.MethodPut)

	fetchProjectByIDUsecase := projectusecase.NewFetchProjectByIDUsecase(projectRepository, productRepository)
	fetchProjectByIDHandler := projecthandler.NewFetchProjectByIDHandler(fetchProjectByIDUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/projects/{projectID:[a-z0-9-]{36}}", fetchProjectByIDHandler.FetchProjectByID).Methods(http.MethodGet)

	fetchProjectsUsecase := projectusecase.NewFetchProjectsUsecase(projectQuery, productRepository)
	fetchProjectsHandler := projecthandler.NewFetchProjectsHandler(fetchProjectsUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/projects", fetchProjectsHandler.FetchProjects).Queries("page", "{page}", "limit", "{limit}").Methods(http.MethodGet)

	searchProjectsUsecase := projectusecase.NewSearchProjectsUsecase(projectQuery, productRepository)
	searchProjectsHandler := projecthandler.NewSearchProjectsHandler(searchProjectsUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/projects/search", searchProjectsHandler.SearchProjects).Queries("keyword", "{keyword}", "page", "{page}", "limit", "{limit}").Methods(http.MethodGet)

	trashedProjectUsecase := projectusecase.NewTrashedProjectUsecase(projectRepository, productRepository)
	trashedProjectHandler := projecthandler.NewTrashedProjectHandler(trashedProjectUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/projects/{projectID:[a-z0-9-]{36}}/trash-box", trashedProjectHandler.TrashedProject).Methods(http.MethodPut)

	restoreFromTrashProjectUsecase := projectusecase.NewRestoreFromTrashProjectUsecase(projectRepository, productRepository)
	restoreFromTrashProjectHandler := projecthandler.NewRestoreFromTrashProjectHandler(restoreFromTrashProjectUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/projects/{projectID:[a-z0-9-]{36}}/restore/trash-box", restoreFromTrashProjectHandler.RestoreFromTrashProject).Methods(http.MethodPut)

	fetchTrashedProjectUsecase := projectusecase.NewFetchTrashedProjectsUsecase(projectQuery, productRepository)
	fetchTrashedProjectHandler := projecthandler.NewFetchTrashedProjectsHandler(fetchTrashedProjectUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/projects/trash-box", fetchTrashedProjectHandler.FetchTrashedProjects).Queries("page", "{page}", "limit", "{limit}").Methods(http.MethodGet)
}
