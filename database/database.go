package database

import "gorm.io/gorm"

type Database interface {
	Connector() *gorm.DB
	Close() error
}
