package model

import (
	"oc-2023/pkg/domain/entity"

	"github.com/google/uuid"
)

type Likes struct {
	UserID uuid.UUID `gorm:"primaryKey" type:"uuid" json:"user_id"`
	WorkID uuid.UUID `gorm:"primaryKey" type:"uuid" json:"work_id"`
}

type LikeArray []Likes

func (la LikeArray) ConvertLikeModelToEntity() []entity.Like {
	entityLikes := make([]entity.Like, len(la))

	for i, ml := range la {
		entityLikes[i] = entity.Like(ml.UserID.String())
	}

	return entityLikes
}
