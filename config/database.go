package config

import (
	"TodoList-GO/app/entities"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error) {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DbHost, DbPort, DbUser, DbPassword, DbName)

	DB, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		return DB, err
	}

	err = DB.AutoMigrate(
		entities.Todo{},
		entities.SubTodo{},
	)
	if err != nil {
		return DB, err
	}

	return DB, nil
}
