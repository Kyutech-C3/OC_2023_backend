package like

import (
	"context"
	"oc-2023/pkg/domain/entity"
	"oc-2023/pkg/domain/model"
	"oc-2023/pkg/domain/repository/like"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	conn *gorm.DB
}

func New(conn *gorm.DB) like.Repository {
	return &repository{
		conn: conn,
	}
}

func (r *repository) Insert(_ context.Context, like *entity.PostLikes) error {
	lk := model.Likes{
		WorkID: like.WorkId,
		UserID: like.UserId,
	}

	if err := r.conn.Create(lk).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(_ context.Context, like *entity.DeleteLikes) error {
	var lk model.Likes
	if err := r.conn.Where("work_id = ? AND user_id = ?", like.WorkId, like.UserId).First(&lk).Error; err != nil {
		return err
	}

	if err := r.conn.Delete(&lk).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) SelectByID(_ context.Context, ID uuid.UUID) ([]entity.Like, error) {
	var likes model.LikeArray
	if err := r.conn.Where("work_id = ?", ID).Find(&likes).Error; err != nil {
		return nil, err
	}

	return likes.ConvertLikeModelToEntity(), nil
}
