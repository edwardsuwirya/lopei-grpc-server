package manager

import "enigmacamp.com/lopei_grpc_srv/repository"

type RepositoryManager interface {
	LopeiRepo() repository.LopeiRepository
}

type repositoryManager struct {
	lopeiRepo repository.LopeiRepository
}

func (r *repositoryManager) LopeiRepo() repository.LopeiRepository {
	return r.lopeiRepo
}

func NewRepositoryManager() RepositoryManager {
	repo := new(repositoryManager)
	repo.lopeiRepo = repository.NewLopeiRepository()
	return repo
}
