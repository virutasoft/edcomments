package models

import "github.com/jinzhu/gorm"

//Comment hace y es la estructura de los commentarios
type Comment struct {
	gorm.Model
	UserID   uint      `json:"userId"`
	ParentID uint      `json:"parentId"`
	Votes    int32     `json:"votes"`
	Content  string    `json:"content"`
	HasVote  uint      `json:"hasVote" gorm:"-"`
	User     []User    `json:"user,omitempty"`
	Children []Comment `json:"children,omitempty"`
}
