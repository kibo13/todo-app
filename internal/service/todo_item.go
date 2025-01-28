package service

import (
	"github.com/kibo13/todo-app/internal/entity"
	"github.com/kibo13/todo-app/internal/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item entity.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]entity.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (entity.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Update(userId, itemId int, input entity.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(userId, itemId, input)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}
