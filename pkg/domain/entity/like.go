package entity

import "github.com/google/uuid"

type PostLikes struct {
	UserId uuid.UUID `type:"uuid" json:"user_id"`
	WorkId uuid.UUID `type:"uuid" json:"work_id"`
}

type DeleteLikes struct {
	UserId uuid.UUID `type:"uuid" json:"user_id"`
	WorkId uuid.UUID `type:"uuid" json:"work_id"`
}

type Like string
