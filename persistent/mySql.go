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
	GetUsers() (*model.User, error)
}

type Mysql struct {
	DB *sql.DB
}

func NewMysqlStore(config config.Mysql) *Mysql {
	db, err := sql.Open(driveName,
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			config.UserName,
			config.Password,
			config.Host,
			config.Port,
			config.Database,
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

func (m *Mysql) GetUsers() (*model.User, error) {
	var user *model.User

	query := `select 
    		id, user_name, full_name, email
		from TestDB
		where user_name = ? and pass = ?`

	err := m.DB.
		QueryRow(query).
		Scan(
			&user.ID,
			&user.Username,
			&user.FullName,
			&user.Email,
		)

	if err != nil {
		return nil, err
	}

	return user, nil
}
