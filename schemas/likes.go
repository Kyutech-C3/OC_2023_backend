package schemas

import (
	"oc-2023/db"

	"github.com/google/uuid"
)

type CreateLikes struct {
	UserId uuid.UUID `type:"uuid" json:"user_id"`
	WorkId uuid.UUID `type:"uuid" json:"work_id"`
}

type DeleteLikes struct {
	UserId uuid.UUID `type:"uuid" json:"user_id"`
	WorkId uuid.UUID `type:"uuid" json:"work_id"`
}

type Like struct {
	UserID uuid.UUID `json:"user_id"`
	WorkID uuid.UUID `json:"work_id"`
}

func ConvertLikeModelToSchemas(likes []db.Likes) []Like {
	schemaLikes := make([]Like, len(likes))

	for i, ml := range likes {
		schemaLikes[i] = Like{
			UserID: ml.UserID,
			WorkID: ml.WorkID,
		}
	}

	return schemaLikes
}
