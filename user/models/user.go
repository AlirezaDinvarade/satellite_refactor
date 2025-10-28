package models

import (
	"time"

	"github.com/google/uuid"
)

type AccessLevelENUM string

const (
	AdminLevel  AccessLevelENUM = "admin"
	ViewerLevel AccessLevelENUM = "viewer"
	UserLevel   AccessLevelENUM = "user"
	ExpertLevel AccessLevelENUM = "expert"
)

type User struct {
	ID          uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	FullName    string          `gorm:"type:varchar(100)" json:"fullName"`
	NationalID  string          `gorm:"uniqueIndex;type:varchar(50)" json:"nationalID"`
	Email       string          `gorm:"type:varchar(50)" json:"email"`
	PhoneNumber string          `gorm:"type:varchar(11);uniqueIndex" json:"phoneNumber"`
	CreatedAt   time.Time       `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt   time.Time       `gorm:"type:timestamp" json:"updatedAt"`
	Password    string          `gorm:"type:varchar(255)" json:"-"`
	AccessLevel AccessLevelENUM `gorm:"type:varchar(10);default:'user'" json:"accessLevel"`
}
