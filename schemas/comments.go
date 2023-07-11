package schemas

import (
	"oc-2023/db"
	"time"

	"github.com/google/uuid"
)

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

type Comment struct {
	CommentID uuid.UUID `json:"comment_id"`
	WorkID    uuid.UUID `json:"work_id"`
	UserID    uuid.UUID `json:"user_id"`
	UserName  string    `json:"user_name"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

func ConvertCommentModelToSchema(comments []db.Comments) []Comment {
	schemaComments := make([]Comment, len(comments))

	for i, mc := range comments {
		schemaComments[i] = Comment{
			CommentID: mc.CommentID,
			UserID:    mc.UserID,
			WorkID:    mc.WorkID,
			UserName:  mc.UserName,
			Comment:   mc.Comment,
			CreatedAt: mc.CreatedAt,
		}
	}

	return schemaComments
}
