package domain

type Writer struct {
	Id   uint64  `json:"id" gorm:"primary_key:auto_increment"`
	Name string  `json:"name" gorm:"type:varchar(255);not null"`
	Book *[]Book `json:"book,omitempty"`
}