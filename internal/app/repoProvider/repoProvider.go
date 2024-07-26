package repoProvider

import (
	"database/sql"
	"github.com/pkg/errors"
	"kinolove/internal/repository"
	"kinolove/pkg/logger"
)

var isInitializedRepo = false

type RepoProvider struct {
	db  *sql.DB
	log logger.Common

	user  repository.UserRepository
	movie repository.MovieRepository
}

func InitRepos(db *sql.DB, log logger.Common) *RepoProvider {

	isInitializedRepo = true

	return &RepoProvider{
		db:  db,
		log: log,
	}
}

func (r *RepoProvider) Storage() *sql.DB {
	if !isInitializedRepo {
		r.log.Fatal(errors.New("Init error"), "Provider is not initialized")
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
