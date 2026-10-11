package main

import(
	"fmt"
	"time"
)

type Task struct{
	id int
	name string
	description string
	taskType string /* онлайн либо не онлайн.
	онлайн, например, сделать что-то на ноуте или на телефоне.
	не онлайн - бытовые дела, либо на бумаге
	*/
	difficultyLevel string // легко, средне, сложно
	deadline time.Time
	completed bool
}

type Storage interface {
    CreateTask(name, description, taskType, difficultyLevel string, deadline time.Time) error
    GetTask(id int) (Task, error)
    UpdateTask(id int, name, description, taskType, difficultyLevel string, deadline time.Time) error
    DeleteTask(id int) error
}


type MemoryStorage struct{
	tasks []Task
	nextId int
}

func NewMemoryStorage() *MemoryStorage{
	return &MemoryStorage{tasks: make([]Task, 0), nextId: 1}
}

// Создание новой задачи
func (mStor *MemoryStorage) CreateTask(name, description, taskType, difficultyLevel string, deadline time.Time){
	task := Task{id: mStor.nextId, name: name, description: description, taskType: taskType, difficultyLevel: difficultyLevel, deadline: deadline, completed: false}
	mStor.tasks = append(mStor.tasks, task)
	mStor.nextId++
}

// Получение задачи по id
func (mStor *MemoryStorage) GetTask(id int) (Task, error){
	for _, task := range mStor.tasks {
		if task.id == id {
			return task, nil
		}
	}

	return Task{}, fmt.Errorf("задача с ID %d не найдена", id)
}

// Изменение данных задачи по id
func (mStor *MemoryStorage) UpdateTask(id int, name, description, taskType, difficultyLevel string, deadline time.Time) error{	
	for i, task := range mStor.tasks {
		if task.id == id {
			mStor.tasks[i] = Task{id: id, name: name, description: description, taskType: taskType, difficultyLevel: difficultyLevel, deadline: deadline, completed: task.completed}
			return nil
		}
	}
	return fmt.Errorf("задача с ID %d не найдена", id)
}


// Удаление задачи по id
func (mStor *MemoryStorage) DeleteTask(id int) error{
	for i, task := range mStor.tasks {
		if task.id == id {
			mStor.tasks = append(mStor.tasks[:i], mStor.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("задача с ID %d не найдена", id)
}

// Выводим задачу
func PrintTask(task Task){
	fmt.Printf("Задача ID %d: %s - %s (Тип: %s, Сложность: %s, Дедлайн: %s)\n", task.id, task.name, task.description, task.taskType, task.difficultyLevel, task.deadline.Format(time.RFC1123))
}


func main(){
	storage := NewMemoryStorage()

	storage.CreateTask("Сделать домашку", "Сделать домашнее задание по математике", "в тетради", "средне", time.Now().Add(48*time.Hour))
	storage.CreateTask("Прочитать книгу", "Прочитать главу из книги, которую я читаю", "онлайн", "легко", time.Now().Add(3*time.Hour))

	err := storage.UpdateTask(2, "Сделать домашку", "Сделать домашнее задание по математике и физике", "в тетради", "сложно", time.Now().Add(96*time.Hour))
	if err != nil {
		fmt.Println("Ошибка при обновлении задачи:", err)
	}

	task, err := storage.GetTask(3)
	if err != nil {
		fmt.Println("Ошибка при получении задачи:", err)
	} else {
		PrintTask(task)
	}

	err = storage.DeleteTask(2)
	if err != nil {
		fmt.Println("Ошибка при удалении задачи:", err)
	}
	for _, task := range storage.tasks {
		PrintTask(task)
	}
}