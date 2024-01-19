package store

import (
	"database/sql"
	"fmt"
	"log"
	"main/config"
	"main/model"
)

const driveName = "mysql"

type Store interface {
	GetUsers() (model.User, error)
}

type Mysql struct {
	DB *sql.DB
}

func NewMysqlStore(c config.Mysql) *Mysql {
	db, err := sql.Open(driveName,
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			c.UserName,
			c.Password,
			c.Host,
			c.Port,
			c.Database,
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &Mysql{DB: db}
}

func (m *Mysql) GetUsers() (model.User, error) {
	var user model.User

	query := `select * from TestDB`

	err := m.DB.
		QueryRow(query).
		Scan(
			&user.ID,
			&user.Username,
			&user.FullName,
			&user.Email,
		)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
