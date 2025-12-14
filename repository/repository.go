package repository

import "github.com/dandimuzaki/project-app-inventaris-cli-dandi-muhamad-zaki/database"

type Repository struct {
	RepoCategory RepositoryCategory
	RepoItem     RepositoryItem
	RepoReport RepositoryReport
}

func NewRepository(db database.PgxIface) Repository {
	return Repository{
		RepoCategory: NewRepoCategory(db),
		RepoItem: NewRepoItem(db),
		RepoReport: NewRepoReport(db),
	}
}