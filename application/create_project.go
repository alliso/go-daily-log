package application

import (
	"go-daily-log/domain/project"
	domain "go-daily-log/domain/repository"
	infr "go-daily-log/infrastructure/db/sqlite/repository"
)

type CreateProjectsCommand struct {
	projectRepository domain.ProjectRepository
}

func CreateProject() *CreateProjectsCommand {
	return &CreateProjectsCommand{
		projectRepository: infr.NewProjectRepositoryImpl(),
	}
}

func (c *CreateProjectsCommand) Create(project project.Project) {
	c.projectRepository.Create(project)
}
