package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/config"
	"github.com/onituka/agile-project-management/project-management/infrastructure/middleware"
	"github.com/onituka/agile-project-management/project-management/infrastructure/persistence"
	"github.com/onituka/agile-project-management/project-management/infrastructure/persistence/rdb"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/usecase/clock"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecse"
)

func Run() error {
	router := mux.NewRouter()

	conn, err := rdb.NewDB()
	if err != nil {
		return err
	}
	defer conn.Close()

	realTime := clock.NewRealTime()

	router.Use(middleware.DBMiddlewareFunc(conn))

	projectRepository := persistence.NewProjectRepository()
	projectUsecase := projectusecse.NewProjectUsecase(projectRepository, realTime)
	projectHandler := handler.NewProjectHandler(projectUsecase)

	router.HandleFunc("/projects", projectHandler.CreateProject).Methods(http.MethodPost)
	router.HandleFunc("/projects/{projectID}", projectHandler.UpdateProject).Methods(http.MethodPut)
	router.HandleFunc("/projects/{projectID}", projectHandler.FetchProjectByID).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Env.Server.Port),
		Handler: router,
	}

	errorCh := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			errorCh <- err
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err := <-errorCh:
		return err
	case s := <-signalCh:
		log.Printf("SIGNAL %s received", s.String())
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			return err
		}
	}

	return nil
}
