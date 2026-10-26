package usecase

import "todo-api/repository"

// Интерфейс TaskUsecase — то, что увидит handler.
// В нём 3 метода — по одному на каждую операцию, которую должен уметь usecase-слой
type TaskUsecase interface {
	CreateTask(title, description string) (repository.Task, error) // создание задачи
	GetTask(id int) (repository.Task, error)                       // получение задачи по id
	ListTask() ([]repository.Task, error)                          // получение всех задач
	DeleteTask(id int) error
	DoneTask(id int) (repository.Task, error)// статус дон
}

// taskUsecase — конкретная реализация интерфейса TaskUsecase
type taskUsecase struct {
	repo repository.TaskRepository
}

func NewTaskUsecase(repo1 repository.TaskRepository) *taskUsecase {
	return &taskUsecase{repo: repo1}

}

func (u *taskUsecase) DeleteTask(id int) error {
	return u.repo.Delete(id)
}
func (u *taskUsecase) DoneTask(id int) (repository.Task, error) {
	return u.repo.Done(id)
}

func (u *taskUsecase) ListTask() ([]repository.Task, error) {
	return u.repo.List() //вызываем метод List() у репозитория, который вернёт слайс задач и ошибку
}

func (u *taskUsecase) GetTask(id int) (repository.Task, error) {
	return u.repo.GetByID(id) //вызываем метод GetByID() у репозитория, который вернёт задачу и ошибку
}
func (u *taskUsecase) CreateTask(title, description string) (repository.Task, error) {
	//IDсоздается в репозитории
	//новая задача
	task := repository.Task{
		Title:       title,
		Description: description,
		Status:      "new", //по умолчанию статус новой задачи "new"
	}
	return u.repo.Create(task) //Метод Create уже возвращает сразу пару (Task, error), и return просто передаёт её дальше. Тип error дописывать нельзя, это снова тип вместо значения
}
