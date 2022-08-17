package db

import (
	"github.com/esonhugh/update-alternative-java/lib/misc"
	"github.com/glebarez/sqlite"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Database struct {
	Main *gorm.DB
}

var DB *Database

func init() {
	assemble()
}

func assemble() {
	main := CreateMain()
	// other db connections
	DB = &Database{
		Main: main,
	}
}

func CreateMain() *gorm.DB {
	db, err := gorm.Open(
		sqlite.Open(misc.DBlocate()),
		&gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}
	return db
}

// func CreateOtherDB[T]() T
