package services

import (
	"gin-fleamarket/dto"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(createItemInput dto.CreateItemInput) (*models.Item, error)
	Update(itemId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error)
	Delete(itemId uint) error
}

type ItemService struct {
	repository repositories.IItemRepository
}

func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{repository: repository}
}

func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}

func (s *ItemService) FindById(itemId uint) (*models.Item, error) {
	return s.repository.FindById(itemId)
}

func (s *ItemService) Create(createItemInput dto.CreateItemInput) (*models.Item, error) {
	newItem := models.Item{
		Price:       createItemInput.Price,
		Name:        createItemInput.Name,
		Description: createItemInput.Description,
		SoldOut:     false,
	}
	return s.repository.Create(newItem)
}

func (s *ItemService) Update(itemId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error) {
	targetItems, err := s.repository.FindById(itemId)
	if err != nil {
		return nil, err
	}
	if updateItemInput.Name != nil {
		targetItems.Name = *updateItemInput.Name
	}
	if updateItemInput.Price != nil {
		targetItems.Price = *updateItemInput.Price
	}
	if updateItemInput.Description != nil {
		targetItems.Description = *updateItemInput.Description
	}
	if updateItemInput.SoldOut != nil {
		targetItems.SoldOut = *updateItemInput.SoldOut
	}
	return s.repository.Update(*targetItems)
}

func (s *ItemService) Delete(itemId uint) error {
	return s.repository.Delete(itemId)
}
