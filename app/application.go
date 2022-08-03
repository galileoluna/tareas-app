package app

import (
	"github.com/galileoluna/tareas/database"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()
	database.Init()
	srv := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
