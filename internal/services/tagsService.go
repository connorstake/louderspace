package services

import (
	"louderspace/internal/repositories"
)

type TagManagement interface {
	GetAllTags() ([]string, error)
}

type TagService struct {
	tagStorage repositories.TagStorage
}

func NewTagService(tagStorage repositories.TagStorage) TagManagement {
	return &TagService{tagStorage}
}

func (s *TagService) GetAllTags() ([]string, error) {
	return s.tagStorage.GetAllTags()
}
