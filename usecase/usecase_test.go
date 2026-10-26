package usecase

import (
	"testing"
	"todo-api/repository"
)

func TestCreateTask(t *testing.T) {
	repo := repository.NewInMemoryRepo() //хранилище
	uc := NewTaskUsecase(repo)           //uc-то что мы проверяем(подключен или нет)

	_, err := uc.CreateTask("тест", "описание") //действие которое проверяем
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	tasks, _ := uc.ListTask() //достаем список задач
	if len(tasks) != 1 {      //ожидаем одну задачу
		t.Errorf("expected 1 task, got %d", len(tasks))
	}
}

func TestDoneTask(t *testing.T) {
	repo := repository.NewInMemoryRepo() // чистое хранилище, чтобы тест не зависел от других
	uc := NewTaskUsecase(repo)           // usecase поверх этого хранилища

	// подготовка: создаём задачу, которую потом отметим выполненной
	task, err := uc.CreateTask("тест", "описание")
	if err != nil {
		// если подготовка не удалась, дальше тест не имеет смысла (Fatalf его останавливает)
		t.Fatalf("unexpected error: %v", err)
	}

	// взял ID созданной задачи и передал в DoneTask
	done, err := uc.DoneTask(task.ID)
	if err != nil {
		// ошибки быть не должно: задача существует
		t.Fatalf("unexpected error donetask: %v", err)
	}

	// главная проверка: статус должен стать "done"
	if done.Status != "done" {
		t.Fatalf("expected status done, got %q", done.Status)
	}
}

func TestDoneTaskNotFound(t *testing.T) { // тест на проверки ошибки TaskNotFound
	repo := repository.NewInMemoryRepo() // чистое хранилище, чтобы тест не зависел от других
	uc := NewTaskUsecase(repo)           // usecase поверх этого хранилища

	_, err := uc.DoneTask(999)
	if err == nil { // если все ок тест падает ПРОВЕРЯЕМ ОШИБКУ
		t.Fatal("expected error, got nil")
	}
}

func TestGetTaskNotFound(t *testing.T) { // тест для проверки ошибки получения задания
	repo := repository.NewInMemoryRepo() // чистое хранилище, чтобы тест не зависел от других
	uc := NewTaskUsecase(repo)           // usecase поверх этого хранилища

	_, err := uc.GetTask(999)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestDeleteTaskNotFound(t *testing.T) { // тест на проверку условия задачи нет, ждём ошибку
	repo := repository.NewInMemoryRepo() // чистое хранилище, чтобы тест не зависел от других
	uc := NewTaskUsecase(repo)           // usecase поверх этого хранилища

	err := uc.DeleteTask(999)
	if err == nil { // если все ок тест падает ПРОВЕРЯЕМ ОШИБКУ
		t.Fatal("expected error, got nil")
	}
}
