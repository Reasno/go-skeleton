package repositories

import (
	"github.com/DoNewsCode/core/otgorm"
	"gorm.io/gorm"
	"time"
)

func ProviderMigrations() []*otgorm.Migration {
	return []*otgorm.Migration{
		{
			ID: "202010280100",
			Migrate: func(db *gorm.DB) error {
				type User struct {
					gorm.Model
					email           string
					username        string
					password        string
					rememberToken   string
					emailVerifiedAt time.Time
				}
				return db.AutoMigrate(
					&User{},
				)
			},
			Rollback: func(db *gorm.DB) error {
				type User struct{}
				return db.Migrator().DropTable(&User{})
			},
		},
	}
}
