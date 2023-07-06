package db

import "github.com/google/uuid"

// データベースの表をイメージして読む

type Likes struct {
	UserID uuid.UUID `gorm:"primaryKey" type:"uuid" json:"user_id"`
	WorkID uuid.UUID `gorm:"primaryKey" type:"uuid" json:"work_id"`
}

type Comments struct {
	CommentID uuid.UUID `gorm:"primaryKey" type:"uuid" json:"comment_id"`
	WorkID    uuid.UUID `type:"uuid" json:"work_id"`
	UserID    uuid.UUID `type:"uuid" json:"user_id"`
	UserName  string    `json:"user_name"`
	Comment   string    `json:"comment"`
}
