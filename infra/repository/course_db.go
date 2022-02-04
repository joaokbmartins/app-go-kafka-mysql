package repository

import (
	"database/sql"

	"github.com/joaokbmartins/app-go-kafka-mysql/entity"
)

type CourseMySqlRepository struct {
	Db *sql.DB
}

func (c CourseMySqlRepository) Insert(course entity.Course) error {

	stmt, err := c.Db.Prepare(`INSERT INTO courses (id, name, description, status) VALUES (?, ?, ?, ?)`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		course.ID,
		course.Name,
		course.Description,
		course.Status,
	)

	if err != nil {
		return err
	}

	return nil

}
