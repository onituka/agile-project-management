package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb/persistence"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler/projectnotehandler"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase"
)

func newProjectNoteRouter(router *mux.Router) {
	projectNoteRepository := persistence.NewProjectNoteRepository()
	productRepository := persistence.NewProductRepository()
	projectRepository := persistence.NewProjectRepository()

	createProjectNoteUsecase := projectnoteusecase.NewCreateProjectNoteUsecase(projectNoteRepository, productRepository, projectRepository)
	createProjectNoteHandler := projectnotehandler.NewCreateProjectNoteHandler(createProjectNoteUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/projects/{projectID:[a-z0-9-]{36}}/notes", createProjectNoteHandler.CreateProjectNote).Methods(http.MethodPost)

	updateProjectNoteRepository := projectnoteusecase.NewUpdateProjectNoteUsecase(projectNoteRepository, productRepository, projectRepository)
	updateProjectNoteHandler := projectnotehandler.NewUpdateProjectNoteHandler(updateProjectNoteRepository)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/projects/{projectID:[a-z0-9-]{36}}/notes/{projectNoteID:[a-z-0-9-]{36}}", updateProjectNoteHandler.UpdateProjectNote).Methods(http.MethodPut)

	fetchProjectNoteRepository := projectnoteusecase.NewFetchProjectNoteByIDUsecase(projectNoteRepository, productRepository, projectRepository)
	fetchProjectNoteHandler := projectnotehandler.NewFetchProjectNoteByIDHandler(fetchProjectNoteRepository)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/projects/{projectID:[a-z0-9-]{36}}/notes/{projectNoteID:[a-z-0-9-]{36}}", fetchProjectNoteHandler.FetchProjectNoteByID).Methods(http.MethodGet)
}
