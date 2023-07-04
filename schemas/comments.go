package schemas

import "github.com/google/uuid"

type PostComments struct {
	UserId   uuid.UUID `type:"uuid" json:"user_id" binding:"required"`
	WorkId   uuid.UUID `type:"uuid" json:"work_id" binding:"required"`
	UserName string    `json:"user_name" binding:"required"`
	Comment  string    `json:"comment" binding:"required"`
}

type DeleteComments struct {
	CommentId uuid.UUID `type:"uuid" json:"comment_id" binding:"required"`
	UserId    uuid.UUID `type:"uuid" json:"user_id" binding:"required"`
}
