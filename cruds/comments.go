package cruds

import (
	"errors"
	"oc-2023/db"
	"time"

	"github.com/google/uuid"
)

func GetCommetsByID(id uuid.UUID) (comments []db.Comments, err error) {
	if err = db.Psql.Where("work_id = ?", id).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func CreateComment(commentId uuid.UUID, workId uuid.UUID, userId uuid.UUID, userName string, comment string) error {
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	cmnt := db.Comments{
		CommentID: commentId,
		WorkID:    workId,
		UserID:    userId,
		UserName:  userName,
		Comment:   comment,
		CreatedAt: time.Now().In(location),
	}

	if err := db.Psql.Create(cmnt).Error; err != nil {
		return err
	}

	return nil
}

func DeleteComment(commentId uuid.UUID, userId uuid.UUID) (err error) {
	if err = db.Psql.Where("comment_id = ? AND user_id = ?", commentId, userId).First(&db.Comments{}).Error; err != nil {
		err = errors.New("Comment not found")
		return
	}

	cmnt := db.Comments{
		CommentID: commentId,
		UserID:    userId,
	}
	if err = db.Psql.Delete(cmnt).Error; err != nil {
		return
	}
	return
}
