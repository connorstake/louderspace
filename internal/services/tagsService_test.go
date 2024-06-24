package services

import (
	"github.com/stretchr/testify/assert"
	"louderspace/internal/repositories"
	"testing"
)

func TestCreateTag(t *testing.T) {
	storage := repositories.NewMockTagStorage()
	service := NewTagService(storage)

	tag, err := service.CreateTag("New Tag")
	assert.NoError(t, err)
	assert.NotNil(t, tag)
	assert.Equal(t, "New Tag", tag.Name)
}

func TestUpdateTag(t *testing.T) {
	storage := repositories.NewMockTagStorage()
	service := NewTagService(storage)

	tag, err := service.CreateTag("Old Tag")
	assert.NoError(t, err)

	tag.Name = "Updated Tag"
	err = service.UpdateTag(tag.ID, tag.Name)
	assert.NoError(t, err)

}

func TestDeleteTag(t *testing.T) {
	storage := repositories.NewMockTagStorage()
	service := NewTagService(storage)

	tag, err := service.CreateTag("Tag to be deleted")
	assert.NoError(t, err)

	err = service.DeleteTag(tag.ID)
	assert.NoError(t, err)

}

func TestGetAllTags(t *testing.T) {
	storage := repositories.NewMockTagStorage()
	service := NewTagService(storage)

	_, err := service.CreateTag("Tag 1")
	assert.NoError(t, err)
	_, err = service.CreateTag("Tag 2")
	assert.NoError(t, err)

	tags, err := service.GetAllTags()
	assert.NoError(t, err)
	assert.Len(t, tags, 2)
}
