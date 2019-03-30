package models

type User struct {
	ID        uint `gorm:"primary_key"`
	Name string
	Languages         []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	ID        uint `gorm:"primary_key"`
	Name string
	Users              []User `gorm:"many2many:user_languages;"`
}
