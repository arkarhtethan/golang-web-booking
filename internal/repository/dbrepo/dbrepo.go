package dbrepo

import (
	"database/sql"

	"github.com/arkarhtethan/golang-web-booking/internal/config"
	"github.com/arkarhtethan/golang-web-booking/internal/repository"
)

type postgresDbRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}
type testDbRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewTestingRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDbRepo{
		App: a,
	}
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDbRepo{
		App: a,
		DB:  conn,
	}
}
