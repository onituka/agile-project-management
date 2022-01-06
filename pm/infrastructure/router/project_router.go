package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb/persistence"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler/projecthandler"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase"
)

func newProjectRouter(router *mux.Router) {
	projectRepository := persistence.NewProjectRepository()

	createProjectUsecase := projectusecase.NewCreateProjectUsecase(projectRepository)
	createProjectHandler := projecthandler.NewCreateProjectHandler(createProjectUsecase)
	router.HandleFunc("/projects", createProjectHandler.CreateProject).Methods(http.MethodPost)

	updateProjectUsecase := projectusecase.NewUpdateProjectUsecase(projectRepository)
	updateProjectHandler := projecthandler.NewUpdateProjectHandler(updateProjectUsecase)
	router.HandleFunc("/projects/{projectID}", updateProjectHandler.UpdateProject).Methods(http.MethodPut)

	fetchProjectByIDUsecase := projectusecase.NewFetchProjectByIDUsecase(projectRepository)
	fetchProjectByIDHandler := projecthandler.NewFetchProjectByIDHandler(fetchProjectByIDUsecase)
	router.HandleFunc("/projects/{projectID}", fetchProjectByIDHandler.FetchProjectByID).Methods(http.MethodGet)

	fetchProjectsUsecase := projectusecase.NewFetchProjectsUsecase(projectRepository)
	fetchProjectsHandler := projecthandler.NewFetchProjectsHandler(fetchProjectsUsecase)
	router.HandleFunc("/projects", fetchProjectsHandler.FetchProjects).Methods(http.MethodGet)

	trashedProjectUsecase := projectusecase.NewTrashedProjectUsecase(projectRepository)
	trashedProjectHandler := projecthandler.NewTrashedProjectHandler(trashedProjectUsecase)
	router.HandleFunc("/projects/{projectID}/trash-box", trashedProjectHandler.TrashedProject).Methods(http.MethodPut)
}
