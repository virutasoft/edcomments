package models

import "github.com/jinzhu/gorm"

//Vote es la estructura de los votos, permite controlar que un usuario solo
// vote una única vez por cada comentario
type Vote struct {
	gorm.Model
	CommentID uint `json:"commentId" gorm:"not null"`
	UserID    uint `json:"userId" gorm:"not null"`
	Value     bool `json:"value" gorm:"not null"`
}
