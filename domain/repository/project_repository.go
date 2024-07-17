package repository

import "go-daily-log/domain/project"

type ProjectRepository interface {
	FindAll() []project.Project
	Create(project.Project)
	Delete(id int)
}
