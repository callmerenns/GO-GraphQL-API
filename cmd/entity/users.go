package entity

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName string         `gorm:"not null;column:firstname" json:"firstname"`
	LastName  string         `gorm:"not null;column:lastname" json:"lastname"`
	Email     string         `gorm:"not null;unique" json:"email"`
	Password  string         `gorm:"not null" json:"password"`
	Role      string         `json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Products  []*Product     `gorm:"many2many:enrollments;" json:"products"`
}

func GetUserIdByEmail(db *gorm.DB, username string) (uint, error) {
	var user User
	if err := db.Where("email = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, fmt.Errorf("user not found")
		}
		return 0, err
	}
	return user.ID, nil
}
