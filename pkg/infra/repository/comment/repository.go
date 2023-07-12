package comment

import (
	"context"
	"oc-2023/pkg/domain/entity"
	"oc-2023/pkg/domain/model"
	"oc-2023/pkg/domain/repository/comment"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	conn *gorm.DB
}

func New(conn *gorm.DB) comment.Repository {
	return &repository{
		conn: conn,
	}
}

func (r *repository) Insert(_ context.Context, comment *entity.PostComments) error {
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	cm := model.Comments{
		CommentID: uuid.New(),
		WorkID:    comment.WorkId,
		UserID:    comment.UserId,
		UserName:  comment.UserName,
		Comment:   comment.Comment,
		CreatedAt: time.Now().In(location),
	}

	if err := r.conn.Create(cm).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(_ context.Context, comment *entity.DeleteComments) error {
	var cm model.Comments
	if err := r.conn.Where("comment_id = ? AND user_id = ?", comment.CommentId, comment.UserId).First(&cm).Error; err != nil {
		return err
	}

	if err := r.conn.Delete(&cm).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) SelectByID(_ context.Context, ID uuid.UUID) ([]entity.Comment, error) {
	var comments model.CommentArray
	if err := r.conn.Where("work_id = ?", ID).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments.ConvertCommentModelToEntity(), nil
}
