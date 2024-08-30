package repoProvider

import (
	"database/sql"
	"github.com/pkg/errors"
	"kinolove/internal/repository"
	"kinolove/pkg/logger"
)

var isInitializedRepo = false

type RepoProvider struct {
	db *sql.DB

	user  repository.UserRepository
	movie repository.MovieRepository
	perm  repository.PermissionRepository
	role  repository.RoleRepository
}

func InitRepos(db *sql.DB) *RepoProvider {

	isInitializedRepo = true

	return &RepoProvider{
		db: db,
	}
}

func (r *RepoProvider) Storage() *sql.DB {
	if !isInitializedRepo {
		logger.Log().Fatal(errors.New("Init error"), "Provider is not initialized")
	}

	return r.db
}

func (r *RepoProvider) UserRepo() repository.UserRepository {
	if r.user != nil {
		return r.user
	}

	repo := repository.NewUserRepository(r.Storage())
	r.user = repo
	return r.user
}

func (r *RepoProvider) MovieRepo() repository.MovieRepository {
	if r.movie != nil {
		return r.movie
	}

	repo := repository.NewMoviesRepository(r.Storage())
	r.movie = repo
	return r.movie
}

func (r *RepoProvider) RoleRepo() repository.RoleRepository {
	if r.role != nil {
		return r.role
	}

	repo := repository.NewRoleRepository(r.Storage())
	r.role = repo
	return r.role
}

func (r *RepoProvider) PermissionRepo() repository.PermissionRepository {
	if r.perm != nil {
		return r.perm
	}

	repo := repository.NewPermissionRepository(r.Storage())
	r.perm = repo
	return r.perm
}
