package dto

import "time"

type Task struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TaskType    string `json:"task_type"` /* онлайн либо не онлайн.
	онлайн, например, сделать что-то на ноуте или на телефоне.
	не онлайн - бытовые дела, либо на бумаге
	*/
	DifficultyLevel string     `json:"difficulty_level"` // легко, средне, сложно
	Deadline        *time.Time `json:"deadline"`
	Completed       bool       `json:"completed"`
}

// // Выводим задачу
// func PrintTask(task Task){
// 	fmt.Printf("Задача ID %d: %s - %s (Тип: %s, Сложность: %s, Дедлайн: %s)\n", task.id, task.name, task.description, task.taskType, task.difficultyLevel, task.deadline.Format(time.RFC1123))
// }

// func main(){
// 	storage := NewMemoryStorage()

// 	storage.CreateTask("Сделать домашку", "Сделать домашнее задание по математике", "в тетради", "средне", time.Now().Add(48*time.Hour))
// 	storage.CreateTask("Прочитать книгу", "Прочитать главу из книги, которую я читаю", "онлайн", "легко", time.Now().Add(3*time.Hour))

// 	err := storage.UpdateTask(2, "Сделать домашку", "Сделать домашнее задание по математике и физике", "в тетради", "сложно", time.Now().Add(96*time.Hour))
// 	if err != nil {
// 		fmt.Println("Ошибка при обновлении задачи:", err)
// 	}

// 	task, err := storage.GetTask(3)
// 	if err != nil {
// 		fmt.Println("Ошибка при получении задачи:", err)
// 	} else {
// 		PrintTask(task)
// 	}

// 	err = storage.DeleteTask(2)
// 	if err != nil {
// 		fmt.Println("Ошибка при удалении задачи:", err)
// 	}
// 	for _, task := range storage.tasks {
// 		PrintTask(task)
// 	}
// }
