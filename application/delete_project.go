package application

import (
	domain "go-daily-log/domain/repository"
	infr "go-daily-log/infrastructure/db/sqlite/repository"
)

type DeleteProjectCommand struct {
	projectRepository domain.ProjectRepository
}

func DeleteProject() *DeleteProjectCommand {
	return &DeleteProjectCommand{
		projectRepository: infr.NewProjectRepositoryImpl(),
	}
}

func (d *DeleteProjectCommand) Delete(id int) {
	d.projectRepository.Delete(id)
}
