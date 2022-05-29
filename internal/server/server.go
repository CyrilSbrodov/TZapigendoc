package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"TZapigendoc/internal/handler"
	"github.com/gorilla/mux"
)

type App struct {
	server *http.Server
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {

	router := mux.NewRouter()
	service := handler.NewService()
	handler := handler.NewHandler(service)
	handler.Register(router)
	//router.HandleFunc("/", handler.Handler)
	a.server = &http.Server{
		Addr:         ":8282",
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("server is listen on port", a.server.Addr)

	<-done

	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := a.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
