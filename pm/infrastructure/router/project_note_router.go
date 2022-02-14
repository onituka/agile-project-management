package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb/persistence"
	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb/query"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler/projectnotehandler"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase"
)

func newProjectNoteRouter(router *mux.Router) {
	projectNoteRepository := persistence.NewProjectNoteRepository()
	projectNoteQueryService := query.NewProjectNoteQueryServiceImpl()
	productRepository := persistence.NewProductRepository()
	projectRepository := persistence.NewProjectRepository()

	createProjectNoteUsecase := projectnoteusecase.NewCreateProjectNoteUsecase(projectNoteRepository, productRepository, projectRepository)
	createProjectNoteHandler := projectnotehandler.NewCreateProjectNoteHandler(createProjectNoteUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/projects/{projectID:[a-z0-9-]{36}}/notes", createProjectNoteHandler.CreateProjectNote).Methods(http.MethodPost)

	updateProjectNoteUsecase := projectnoteusecase.NewUpdateProjectNoteUsecase(projectNoteRepository, productRepository, projectRepository)
	updateProjectNoteHandler := projectnotehandler.NewUpdateProjectNoteHandler(updateProjectNoteUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/projects/{projectID:[a-z0-9-]{36}}/notes/{projectNoteID:[a-z-0-9-]{36}}", updateProjectNoteHandler.UpdateProjectNote).Methods(http.MethodPut)

	fetchProjectNoteUsecase := projectnoteusecase.NewFetchProjectNoteByIDUsecase(projectNoteRepository, productRepository, projectRepository)
	fetchProjectNoteHandler := projectnotehandler.NewFetchProjectNoteByIDHandler(fetchProjectNoteUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/projects/{projectID:[a-z0-9-]{36}}/notes/{projectNoteID:[a-z-0-9-]{36}}", fetchProjectNoteHandler.FetchProjectNoteByID).Methods(http.MethodGet)

	fetchProjectNotesUsecase := projectnoteusecase.NewFetchProjectNotesUsecase(projectNoteQueryService, productRepository, projectRepository)
	fetchProjectNotesHandler := projectnotehandler.NewFetchProjectNotesHandler(fetchProjectNotesUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/projects/{projectID:[a-z0-9-]{36}}/notes", fetchProjectNotesHandler.FetchProjectNotes).Queries("page", "{page}", "limit", "{limit}").Methods(http.MethodGet)

	searchProjectNotesUsecase := projectnoteusecase.NewSearchProjectNotesUsecase(projectNoteQueryService, productRepository, projectRepository)
	searchProjectNotesHandler := projectnotehandler.NewSearchProjectNotesHandler(searchProjectNotesUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/projects/{projectID:[a-z0-9-]{36}}/notes/search", searchProjectNotesHandler.SearchProjectNotes).Queries("title", "{title}", "page", "{page}", "limit", "{limit}").Methods(http.MethodGet)

	deleteProjectNoteUsecase := projectnoteusecase.NewDeleteProjectNoteUsecase(projectNoteRepository, productRepository, projectRepository)
	deleteProjectNoteHandler := projectnotehandler.NewDeleteProjectNoteHandler(deleteProjectNoteUsecase)
	router.HandleFunc("/products/{productID:[a-z0-9-]{36}}/projects/{projectID:[a-z0-9-]{36}}/notes/{projectNoteID:[a-z-0-9-]{36}}/delete", deleteProjectNoteHandler.DeleteProjectNote).Methods(http.MethodDelete)
}
