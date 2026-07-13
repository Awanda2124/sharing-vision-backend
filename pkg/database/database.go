package database

import (
	"github.com/awanda/backend-repo/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(cfg *configs.Config) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(cfg.DSN()), &gorm.Config{})
}
