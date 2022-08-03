package database

import (
	"context"
	"fmt"
	"github.com/galileoluna/tareas/domain"
)

func (repo *PostgresRepository) InsertTarea(ctx context.Context, tarea *domain.Tarea) error {
	fmt.Println("Insert Tarea")
	_, err := repo.db.ExecContext(ctx, "INSERT INTO tareas (id, titulo) VALUES ($1, $2)", tarea.ID, tarea.Titulo)
	return err
}

func (repo *PostgresRepository) ListTareas(ctx context.Context) ([]*domain.Tarea, error) {
	fmt.Println("List Tareas")
	rows, err := repo.db.QueryContext(ctx, "SELECT id, titulo FROM tareas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tareas := []*domain.Tarea{}
	for rows.Next() {
		tarea := &domain.Tarea{}
		if err := rows.Scan(&tarea.ID, &tarea.Titulo); err != nil {
			return nil, err
		}
		tareas = append(tareas, tarea)
	}
	return tareas, nil
}
