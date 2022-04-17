package storage

import (
	"fmt"
	"visable/internal/config"
	"visable/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(conf *config.Config) (*Storage, error) {
	dns := fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Dbname,
	)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{Logger: logger.Default})

	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

func (x *Storage) Migrate() error {
	return x.db.AutoMigrate(
		&model.Project{},
		&model.Owner{},
		&model.Participant{},
	)
}
