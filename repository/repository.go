package repository

import (
	"context"
	"github.com/galileoluna/tareas/domain"
)

type Repository interface {
	Close()
	InsertTarea(ctx context.Context, tarea *domain.Tarea) error
	ListTareas(ctx context.Context) ([]*domain.Tarea, error)
}

var repository Repository

func SetRepository(repo Repository) {
	repository = repo
}
