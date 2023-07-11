package model

import (
	"oc-2023/pkg/domain/entity"
	"time"

	"github.com/google/uuid"
)

type Comments struct {
	CommentID uuid.UUID `gorm:"primaryKey" type:"uuid" json:"comment_id"`
	WorkID    uuid.UUID `type:"uuid" json:"work_id"`
	UserID    uuid.UUID `type:"uuid" json:"user_id"`
	UserName  string    `json:"user_name"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time
}

type CommentArray []Comments

func (ca CommentArray) ConvertCommentModelToEntity() []entity.Comment {
	entityComments := make([]entity.Comment, len(ca))

	for i, mc := range ca {
		entityComments[i] = entity.Comment{
			CommentID: mc.CommentID,
			UserID:    mc.UserID,
			WorkID:    mc.WorkID,
			UserName:  mc.UserName,
			Comment:   mc.Comment,
			CreatedAt: mc.CreatedAt,
		}
	}

	return entityComments
}
