package services

import (
	"louderspace/internal/models"
	"louderspace/internal/repositories"
)

type TagManagement interface {
	GetAllTags() ([]*models.Tag, error)
	CreateTag(name string) (*models.Tag, error)
	UpdateTag(id int, name string) error
	DeleteTag(id int) error
}

type TagService struct {
	tagStorage repositories.TagStorage
}

func NewTagService(tagStorage repositories.TagStorage) TagManagement {
	return &TagService{tagStorage}
}

func (s *TagService) GetAllTags() ([]*models.Tag, error) {
	return s.tagStorage.GetAllTags()
}

func (s *TagService) CreateTag(name string) (*models.Tag, error) {
	tag := &models.Tag{Name: name}
	if err := s.tagStorage.Create(tag); err != nil {
		return nil, err
	}
	return tag, nil
}

func (s *TagService) UpdateTag(id int, name string) error {
	tag := &models.Tag{ID: id, Name: name}
	return s.tagStorage.Update(tag)
}

func (s *TagService) DeleteTag(id int) error {
	return s.tagStorage.Delete(id)
}
