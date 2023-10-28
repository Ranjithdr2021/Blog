package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Credentials struct {
	Username string
	Password string
	Server   string
	Dbname   string
}

var Database = Credentials{
	Username: "root",
	Password: "Ranjith",
	Server:   "tcp(localhost:3306)",
	Dbname:   "blogposts",
}

func (m Credentials) ConnectToDB() *gorm.DB {
	// Connect to the database
	dataSourceName := m.Username + ":" + m.Password + "@" + m.Server + "/" + m.Dbname
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	return db
}

func NewDb() *Credentials {
	return &Database
}
