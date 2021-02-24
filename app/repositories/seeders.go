package repositories

import (
	"fmt"
	"github.com/DoNewsCode/core/otgorm"
	"github.com/nfangxu/core-skeleton/app/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"math/rand"
)

func ProvideSeeders() []*otgorm.Seed {
	rf := func(n int) string {
		letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
		b := make([]rune, n)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		return string(b)
	}

	return []*otgorm.Seed{
		{
			Name: "seeding mysql",
			Run: func(db *gorm.DB) error {
				for i := 0; i < 100; i++ {
					pwd, _ := bcrypt.GenerateFromPassword([]byte("L0vePhPFore4er"), bcrypt.DefaultCost)
					db.Create(&entities.User{
						Username: rf(10),
						Email:    fmt.Sprintf("%s@%s.com", rf(10), rf(6)),
						Password: string(pwd),
					})
				}
				return nil
			},
		},
	}
}
