package db

import (
	"database/sql"
	"fmt"
	"github.com/friendsofgo/errors"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"random.chars.jp/git/misskey/config"
)

var DB *sql.DB

func Open() error {
	sslMode := "disable"
	if config.External.Postgres.SSL {
		sslMode = "enable"
	}
	if db, err := sql.Open("postgres",
		fmt.Sprintf(
			"host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
			config.External.Postgres.Host,
			config.External.Postgres.Port,
			config.External.Postgres.DB,
			config.External.Postgres.Username,
			config.External.Postgres.Password,
			sslMode,
		)); err != nil {
		return err
	} else {
		DB = db
	}

	boil.SetDB(DB)
	return nil
}

func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return errors.New("database was not initialized")
}
