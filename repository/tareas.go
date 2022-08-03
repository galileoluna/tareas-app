package repository

import (
	"context"
	"github.com/galileoluna/tareas/domain"
)

func InsertTarea(ctx context.Context, tarea *domain.Tarea) error {
	return repository.InsertTarea(ctx, tarea)
}

func ListTareas(ctx context.Context) ([]*domain.Tarea, error) {
	return repository.ListTareas(ctx)
}
func Close() {
	repository.Close()
}
