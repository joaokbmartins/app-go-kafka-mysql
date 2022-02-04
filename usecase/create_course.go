package usecase

import (
	"github.com/google/uuid"
	"github.com/joaokbmartins/app-go-kafka-mysql/entity"
)

type CreateCourse struct {
	Repository entity.CourseRepository
}

func (c CreateCourse) Execute(input CreateCourseInputDTO) (CreateCourseOutputDTO, error) {

	course := entity.Course{}
	course.ID = uuid.New().String()
	course.Name = input.Name
	course.Description = input.Description
	course.Status = input.Status

	err := c.Repository.Insert(course)
	if err != nil {
		return CreateCourseOutputDTO{}, err
	}

	output := CreateCourseOutputDTO{}
	output.Description = course.Description
	output.Name = course.Name
	output.ID = course.ID
	output.Status = course.Status

	return output, nil

}
