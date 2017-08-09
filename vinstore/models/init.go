package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// Empty import necessary to use the MySQL dialect.
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Create creates the gorm database and creates the tables in the database if
// necessary.
func Create(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dsn+"?charset=utf8&parseTime=True")
	if err != nil {
		return nil, err
	}

	//db = db.Debug()

	db.AutoMigrate(&DBVersion{})
	dbv := DBVersion{-1}
	db.First(&dbv)
	if dbv.Version == -1 {
		db.Create(&DBVersion{Version: 0})
		dbv.Version = 0
	}

	if dbv.Version < len(migrations) {
		for _, f := range migrations[dbv.Version:] {
			fmt.Println("Running migration", dbv.Version)
			f(db)
			dbv.Version++
		}
	}

	// FOR SOME REASON IT DOESN'T WORK WITH db.Update AND I AM TIRED OF THIS SHIT
	db.Exec("UPDATE db_versions SET version = ?", dbv.Version)

	return db, nil
}

// DBVersion is a table holding a single value: the version of the database.
type DBVersion struct {
	Version int
}

var migrations = []func(db *gorm.DB){
	func(db *gorm.DB) {
		db.CreateTable(&Match{}, &MatchRedirect{}, &Game{})
		db.Model(&Game{}).AddForeignKey("match_id", "matches(id)", "CASCADE", "CASCADE")
	},
}
