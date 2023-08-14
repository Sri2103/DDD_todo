package db

import (
	todo_repository "github.com/sri2103/domain_DD_todo/internal/app/todo/repository"
	user_repository "github.com/sri2103/domain_DD_todo/internal/app/user/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// ConnectToDB establishes a connection to the PostgreSQL database.
func ConnectToDB() (*gorm.DB, error) {
    dsn := "host=localhost user=postgresUser password=postgresPW dbname=postgresDB port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}

// GetDB returns the global GORM database instance.
func GetDB() *gorm.DB {
    return db
}

// MigrateModels runs the auto-migration for the models.
func MigrateModels(db *gorm.DB) error {
    return db.AutoMigrate(
		&todo_repository.Pg_Todo{},
		&user_repository.User_Pg_Todo{},
	)
}
