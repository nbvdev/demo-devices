package database

import (
	"database/sql"
	"embed"
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"os"
)

func NewMysqlConfig() mysql.Config {
	return mysql.Config{
		User:      os.Getenv("DB_USER_NAME"),
		Passwd:    os.Getenv("DB_USER_PASSWORD"),
		Net:       "tcp",
		Addr:      os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName:    os.Getenv("DB_NAME"),
		ParseTime: true,
	}
}

func NewDbConnection(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ApplyMigrations(fs *embed.FS, path string, cfg mysql.Config) error {
	d, err := iofs.New(fs, path)
	if err != nil {
		return err
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, "mysql://"+cfg.FormatDSN())
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
