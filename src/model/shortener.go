package model

// tag anotation para adicionar funcionalidades
type Shortener struct {
	ID          uint   `json:"id" gorm:"primaryKey; autoIncrement"`
	Slug        string `json:"slug" gorm:"unique;not null; varchar(50)"`
	OriginalUrl string `json:"url" gorm:"not null; varchar(2083)" binding:"required"`
}
