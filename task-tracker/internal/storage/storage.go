
package storage

import (
	"context"
	"fmt"

	"task-tracker/internal/dto"
)

type Storage interface {
	CreateTask(ctx context.Context, task dto.Task) (dto.Task, error)
	GetTask(ctx context.Context, id int64) (dto.Task, error)
	UpdateTask(ctx context.Context, task dto.Task) error
	DeleteTask(ctx context.Context, id int64) error
}

type MemoryStorage struct {
	tasks  map[int64]dto.Task
	nextId int64
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tasks:  make(map[int64]dto.Task),
		nextId: 1,
	}
}

// Создание задачи
func (ms *MemoryStorage) CreateTask(ctx context.Context, task dto.Task) (dto.Task, error) {
	if err := ctx.Err(); err != nil {
		return dto.Task{}, err
	}

	task.Id = ms.nextId
	ms.tasks[task.Id] = task
	ms.nextId++

	return task, nil
}

// Получение задачи
func (ms *MemoryStorage) GetTask(ctx context.Context, id int64) (dto.Task, error) {
	if err := ctx.Err(); err != nil {
		return dto.Task{}, err
	}

	task, exists := ms.tasks[id]
	if !exists {
		return dto.Task{}, fmt.Errorf("задача с ID %d не найдена", id)
	}

	return task, nil
}

// Обновление задачи
func (ms *MemoryStorage) UpdateTask(ctx context.Context, task dto.Task) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	if _, exists := ms.tasks[task.Id]; !exists {
		return fmt.Errorf("задача с ID %d не найдена", task.Id)
	}

	ms.tasks[task.Id] = task

	return nil
}

// Удаление задачи
func (ms *MemoryStorage) DeleteTask(ctx context.Context, id int64) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	if _, exists := ms.tasks[id]; !exists {
		return fmt.Errorf("задача с ID %d не найдена", id)
	}

	delete(ms.tasks, id)

	return nil
}
