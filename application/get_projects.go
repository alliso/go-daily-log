package application

import (
	"go-daily-log/domain/project"
	domain "go-daily-log/domain/repository"
	infr "go-daily-log/infrastructure/db/sqlite/repository"
)

type GetProjectsCommand struct {
	projectRepository domain.ProjectRepository
}

func GetProjects() *GetProjectsCommand {
	return &GetProjectsCommand{
		projectRepository: infr.NewProjectRepositoryImpl(),
	}
}

func (g *GetProjectsCommand) Get() []project.Project {
	return g.projectRepository.FindAll()
}
