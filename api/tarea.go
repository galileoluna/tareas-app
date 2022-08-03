package api

import (
	"encoding/json"
	"fmt"
	"github.com/galileoluna/tareas/domain"
	"github.com/galileoluna/tareas/repository"
	"github.com/segmentio/ksuid"
	"net/http"
)

func (tareaAPI *TareasAPIImpl) InsertTarea(w http.ResponseWriter, r *http.Request) {
	var nuevaTarea domain.Tarea
	err := json.NewDecoder(r.Body).Decode(&nuevaTarea)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	id, err := ksuid.NewRandom()
	tarea := domain.Tarea{ID: id.String(), Titulo: nuevaTarea.Titulo}
	if err := repository.InsertTarea(r.Context(), &tarea); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func (tareaAPI *TareasAPIImpl) GetTareas(w http.ResponseWriter, r *http.Request) {
	tareas, _ := repository.ListTareas(r.Context())
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tareas); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
