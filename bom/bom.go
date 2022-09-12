package bom

import "database/sql"

type App struct {
	DB *sql.DB
}

func New(c Config) (*App, error) {
	db, err := sql.Open("postgres", c.DBURL)
	if err != nil {
		return nil, err
	}
	return &App{DB: db}, nil
}

type Config struct {
	Listen string
	DBURL  string
}
