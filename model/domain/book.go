package domain

type Book struct {
	Id           uint64 `json:"id" gorm:"primary_key:auto_increment"`
	Title        string `json:"title" gorm:"type:varchar(255);not null"`
	Publisher    string `json:"publisher" gorm:"type:varchar(255);not null"`
	BookType     string `json:"book_type" gorm:"type:varchar(255);not null"`
	YearReleased string `json:"year_released" gorm:"type:varchar(4);not null"`
	Synopsis     string `json:"synopsis" gorm:"type:text;not null"`
	Genre        string `json:"genre" gorm:"type:varchar(255);not null"`
	Stock        uint64 `json:"stock" gorm:"type:int;not null"`
	WriterId     uint64 `json:"-" gorm:"not null"`
	Writer       Writer `json:"writer" gorm:"foreignkey:WriterId;constraint:onUpdate:CASCADE.onDelete:CASCADE"`
}