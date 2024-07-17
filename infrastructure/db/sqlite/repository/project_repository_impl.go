package repository

import (
	"database/sql"
	"go-daily-log/domain/project"
	"go-daily-log/infrastructure/db/sqlite/conf"
)

type ProjectRepositoryImpl struct {
	db *sql.DB
}

func NewProjectRepositoryImpl() *ProjectRepositoryImpl {
	return &ProjectRepositoryImpl{db: conf.Get()}
}

func (p *ProjectRepositoryImpl) Create(project project.Project) {

	sql := `INSERT INTO projects(name, description) 
			VALUES (?,?)`

	if _, err := p.db.Exec(sql, project.Name, project.Description); err != nil {
		panic(err)
	}
}

func (p *ProjectRepositoryImpl) FindAll() []project.Project {
	var projects []project.Project
	sql := `SELECT id, name, description
		FROM projects`

	if rows, err := p.db.Query(sql); err != nil {
		panic(err)
	} else {
		for rows.Next() {
			var project project.Project
			if err := rows.Scan(&project.Id, &project.Name, &project.Description); err != nil {
				panic(err)
			}
			projects = append(projects, project)
		}
	}

	return projects

}

func (p *ProjectRepositoryImpl) Delete(id int) {
	sql := `DELETE FROM projects WHERE id = ?`

	if _, err := p.db.Exec(sql, int64(id)); err != nil {
		panic(err)
	}
}
