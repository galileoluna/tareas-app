package api

import "net/http"

type TareasAPI interface {
	InsertTarea(w http.ResponseWriter, r *http.Request)
	//GetTarea(w http.ResponseWriter, r *http.Request)
	//UpdateTarea(w http.ResponseWriter, r *http.Request)
	GetTareas(w http.ResponseWriter, r *http.Request)
	//DeleteTarea(w http.ResponseWriter, r *http.Request)
}

type TareasAPIImpl struct {
}

var (
	TareaAPI TareasAPI = &TareasAPIImpl{}
)
