package dbs

import (
	"time"

	"github.com/jinzhu/gorm"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func newGormModel(id uint, now time.Time) gorm.Model {
	return gorm.Model{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: &now,
	}
}

func connectDB(dns string) (*gorm.DB, sqlmock.Sqlmock) {
	_, mock, _ := sqlmock.NewWithDSN(dns)
	db, _ := gorm.Open("sqlmock", dns)

	return db, mock
}
