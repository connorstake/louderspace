package repositories

import (
	"database/sql"
)

type TagStorage interface {
	GetAllTags() ([]string, error)
}

type TagDatabase struct {
	db *sql.DB
}

func NewTagDatabase(db *sql.DB) TagStorage {
	return &TagDatabase{db}
}

func (r *TagDatabase) GetAllTags() ([]string, error) {
	var tags []string
	rows, err := r.db.Query("SELECT name FROM tags")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
