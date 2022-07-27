package domain

import "time"

type User struct {
	Id               uint64    `json:"id" gorm:"primary_key:auto_increment"`
	Username         string    `json:"username" gorm:"type:varchar(255);not null, unique"`
	Password         string    `json:"password" gorm:"type:varchar(255);not null"`
	Email            string    `json:"email" gorm:"type:varchar(255);not null, unique"`
	VerificationTime time.Time `json:"verification_time" gorm:"type:datetime"`
	Token            string    `json:"token,omitempty" gorm:"-"`
}
