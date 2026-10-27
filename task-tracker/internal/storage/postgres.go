package storage

import (
	"context"
	"database/sql"
	"fmt"

	"task-tracker/internal/dto"
)

type PostgresStorage struct {
	db *sql.DB
}

// NewPostgresStorage создаёт PostgreSQL-хранилище.
func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{
		db: db,
	}
}

// CreateTask создаёт задачу и возвращает её с присвоенным ID.
func (ps *PostgresStorage) CreateTask(
	ctx context.Context,
	task dto.Task,
) (dto.Task, error) {
	query := `
		INSERT INTO tasks (
			title,
			description,
			task_type,
			difficulty_level,
			deadline,
			completed
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	err := ps.db.QueryRowContext(
		ctx,
		query,
		task.Title,
		task.Description,
		task.TaskType,
		task.DifficultyLevel,
		task.Deadline,
		task.Completed,
	).Scan(&task.Id)

	if err != nil {
		return dto.Task{}, fmt.Errorf("создание задачи: %w", err)
	}

	return task, nil
}

// GetTask получает задачу по ID.
func (ps *PostgresStorage) GetTask(
	ctx context.Context,
	id int64,
) (dto.Task, error) {
	query := `
		SELECT
			id,
			title,
			description,
			task_type,
			difficulty_level,
			deadline,
			completed
		FROM tasks
		WHERE id = $1
	`

	var task dto.Task
	var deadline sql.NullTime

	err := ps.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&task.Id,
		&task.Title,
		&task.Description,
		&task.TaskType,
		&task.DifficultyLevel,
		&deadline,
		&task.Completed,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return dto.Task{}, fmt.Errorf(
				"задача с ID %d не найдена",
				id,
			)
		}

		return dto.Task{}, fmt.Errorf("получение задачи: %w", err)
	}

	if deadline.Valid {
		task.Deadline = &deadline.Time
	}

	return task, nil
}

// UpdateTask обновляет существующую задачу.
func (ps *PostgresStorage) UpdateTask(
	ctx context.Context,
	task dto.Task,
) error {
	query := `
		UPDATE tasks
		SET
			title = $1,
			description = $2,
			task_type = $3,
			difficulty_level = $4,
			deadline = $5,
			completed = $6
		WHERE id = $7
	`

	result, err := ps.db.ExecContext(
		ctx,
		query,
		task.Title,
		task.Description,
		task.TaskType,
		task.DifficultyLevel,
		task.Deadline,
		task.Completed,
		task.Id,
	)

	if err != nil {
		return fmt.Errorf("обновление задачи: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("проверка обновления задачи: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("задача с ID %d не найдена", task.Id)
	}

	return nil
}

// DeleteTask удаляет задачу по ID.
func (ps *PostgresStorage) DeleteTask(
	ctx context.Context,
	id int64,
) error {
	query := `
		DELETE FROM tasks
		WHERE id = $1
	`

	result, err := ps.db.ExecContext(
		ctx,
		query,
		id,
	)

	if err != nil {
		return fmt.Errorf("удаление задачи: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("проверка удаления задачи: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("задача с ID %d не найдена", id)
	}

	return nil
}

// Проверка, что PostgresStorage реализует интерфейс Storage.
var _ Storage = (*PostgresStorage)(nil)
