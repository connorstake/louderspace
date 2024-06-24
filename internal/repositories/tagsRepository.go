package repositories

import (
	"database/sql"
	"louderspace/internal/models"
)

type TagStorage interface {
	GetAllTags() ([]*models.Tag, error)
	Create(tag *models.Tag) error
	Update(tag *models.Tag) error
	Delete(id int) error
}

type TagDatabase struct {
	db *sql.DB
}

func NewTagDatabase(db *sql.DB) TagStorage {
	return &TagDatabase{db}
}

func (r *TagDatabase) GetAllTags() ([]*models.Tag, error) {
	var tags []*models.Tag
	rows, err := r.db.Query("SELECT id, name FROM tags")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tag models.Tag
		if err := rows.Scan(&tag.ID, &tag.Name); err != nil {
			return nil, err
		}
		tags = append(tags, &tag)
	}

	return tags, nil
}

func (r *TagDatabase) Create(tag *models.Tag) error {
	query := "INSERT INTO tags (name) VALUES ($1) RETURNING id"
	return r.db.QueryRow(query, tag.Name).Scan(&tag.ID)
}

func (r *TagDatabase) Update(tag *models.Tag) error {
	query := "UPDATE tags SET name = $1 WHERE id = $2"
	_, err := r.db.Exec(query, tag.Name, tag.ID)
	return err
}

func (r *TagDatabase) Delete(id int) error {
	query := "DELETE FROM tags WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}
