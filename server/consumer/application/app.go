package application

import (
	"context"
	"fmt"
	"net/http"
)

type App struct{
	router http.Handler
}

func New() *App{ 
	app := &App{
		router: loadRouters(),
	}
	fmt.Println("Successfully run App!")
	return app
}

func (a *App) Start(ctx context.Context) error{
	
	server := &http.Server{
		Addr: ":3001",
		Handler: a.router,
	}
	err := server.ListenAndServe()
	if err !=nil{
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}