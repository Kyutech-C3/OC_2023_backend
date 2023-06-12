package db

// データベースの表をイメージして読む

type Likes struct {
	UserID string `gorm:"primaryKey" json:"user_id"`
	WorkID string `gorm:"primaryKey" json:"work_id"`
}

type Comments struct {
	CommentID       string `gorm:"primaryKey" json:"comment_id"`
	WorkID          string `gorm:"primaryKey" json:"work_id"`
	CommentUserName string `json:"user_name"`
	Comment         string `json:"comment"`
}
