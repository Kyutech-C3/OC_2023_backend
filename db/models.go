package db

import "github.com/google/uuid"

// データベースの表をイメージして読む

type Likes struct {
	UserID uuid.UUID `gorm:"primaryKey" type:"uuid" json:"user_id"`
	WorkID uuid.UUID `gorm:"primaryKey" type:"uuid" json:"work_id"`
}

type Comments struct {
	CommentID       string `gorm:"primaryKey" json:"comment_id"`
	WorkID          uuid.UUID `gorm:"primaryKey" json:"work_id"`
	CommentUserName string `json:"user_name"`
	Comment         string `json:"comment"`
}
