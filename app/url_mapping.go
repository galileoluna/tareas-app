package app

import (
	"github.com/galileoluna/tareas/api"
	"net/http"
)

func mapUrls() {
	tareasRouter := router.PathPrefix("/tareas").Subrouter()
	tareasRouter.Path("").Methods(http.MethodPost).HandlerFunc(api.TareaAPI.InsertTarea)
	/*
		tareasRouter.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(controllers.EventoController.DeleteEvento)
		tareasRouter.Path("/{id}").Methods(http.MethodPut).HandlerFunc(controllers.EventoController.UpdateEvento)
		tareasRouter.Path("/data/{id}").Methods(http.MethodGet).HandlerFunc(controllers.EventoController.GetEvento)
	*/
	tareasRouter.Path("/all").Methods(http.MethodGet).HandlerFunc(api.TareaAPI.InsertTarea)
}
