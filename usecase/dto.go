package usecase

type CreateCourseInputDTO struct {
	Name        string `json: "name"`
	Description string `json: "description"`
	Status      string `json: "status"`
}

type CreateCourseOutputDTO struct {
	ID          string `json: "id"`
	Name        string `json: "name"`
	Description string `json: "description"`
	Status      string `json: "status"`
}
