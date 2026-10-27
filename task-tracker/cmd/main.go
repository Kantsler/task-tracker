package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"task-tracker/internal/app"
	"task-tracker/internal/storage"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := sql.Open(
		"pgx",
		"postgres://task_tracker:task_tracker_password@localhost:5432/task_tracker?sslmode=disable",
	)
	if err != nil {
		panic(fmt.Sprintf("Ошибка открытия соединения с БД: %v", err))
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		panic(fmt.Sprintf("Ошибка подключения к БД: %v", err))
	}

	fmt.Println("Подключение к PostgreSQL успешно!")

	taskStorage := storage.NewPostgresStorage(db)

	e := app.BuildApp(taskStorage)
	defer e.Close()

	if err := e.Start(":8080"); err != nil {
		panic(fmt.Sprintf("Ошибка при запуске приложения: %v", err))
	}
}
