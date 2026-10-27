package app

import (
	"net/http"

	"task-tracker/internal/dto"
	"task-tracker/internal/storage"

	"github.com/labstack/echo/v4"
)

func BuildApp(taskStorage storage.Storage) *echo.Echo {
	e := echo.New()

	e.POST("/tasks", func(c echo.Context) error {
		var task dto.Task

		if err := c.Bind(&task); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "некорректный JSON",
			})
		}

		createdTask, err := taskStorage.CreateTask(
			c.Request().Context(),
			task,
		)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "не удалось создать задачу",
			})
		}

		return c.JSON(http.StatusCreated, createdTask)
	})

	return e
}
