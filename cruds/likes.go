package cruds

import (
	"errors"
	"oc-2023/db"

	"github.com/google/uuid"
)

func CreateLike(workId uuid.UUID, userId uuid.UUID) error {
	lk := db.Likes{
		WorkID: workId,
		UserID: userId,
	}

	if err := db.Psql.Create(lk).Error; err != nil {
		return err
	}

	return nil
}

func DeleteLike(workId uuid.UUID, userId uuid.UUID) (err error) {
	if err = db.Psql.Where("work_id = ? AND user_id = ?", workId, userId).First(&db.Likes{}).Error; err != nil {
		err = errors.New("Like not found")
		return
	}

	lk := db.Likes{
		WorkID: workId,
		UserID: userId,
	}
	if err = db.Psql.Delete(lk).Error; err != nil {
		return
	}
	return
}
