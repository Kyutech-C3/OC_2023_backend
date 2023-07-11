//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE
//go:generate goimports -w --local "github.com/sivchari/chat-answer" mock_$GOPACKAGE/mock_$GOFILE
package usecase

import (
	"oc-2023/pkg/domain/entity"
	commentrepository "oc-2023/pkg/domain/repository/comment"
	likerepository "oc-2023/pkg/domain/repository/like"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Interactor interface {
	CreateComment(ctx *gin.Context, comment *entity.PostComments) error
	DeleteComment(ctx *gin.Context, comment *entity.DeleteComments) error
	GetCommentsByID(ctx *gin.Context, ID uuid.UUID) ([]entity.Comment, error)
	CreateLike(ctx *gin.Context, like *entity.PostLikes) error
	DeleteLike(ctx *gin.Context, like *entity.DeleteLikes) error
	GetLikesByID(ctx *gin.Context, ID uuid.UUID) ([]entity.Like, error)
}

type interactor struct {
	commentRepository commentrepository.Repository
	likeRepository    likerepository.Repository
}

func NewInteractor(
	commentRepository commentrepository.Repository,
	likeRepository likerepository.Repository,
) Interactor {
	return &interactor{
		commentRepository,
		likeRepository,
	}
}

func (i *interactor) CreateComment(ctx *gin.Context, comment *entity.PostComments) error {
	if err := i.commentRepository.Insert(ctx, comment); err != nil {
		return err
	}

	return nil
}

func (i *interactor) DeleteComment(ctx *gin.Context, comment *entity.DeleteComments) error {
	if err := i.commentRepository.Delete(ctx, comment); err != nil {
		return err
	}

	return nil
}

func (i *interactor) GetCommentsByID(ctx *gin.Context, ID uuid.UUID) ([]entity.Comment, error) {
	comments, err := i.commentRepository.SelectByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (i *interactor) CreateLike(ctx *gin.Context, like *entity.PostLikes) error {
	if err := i.likeRepository.Insert(ctx, like); err != nil {
		return err
	}

	return nil
}

func (i *interactor) DeleteLike(ctx *gin.Context, like *entity.DeleteLikes) error {
	if err := i.likeRepository.Delete(ctx, like); err != nil {
		return err
	}

	return nil
}

func (i *interactor) GetLikesByID(ctx *gin.Context, ID uuid.UUID) ([]entity.Like, error) {
	likes, err := i.likeRepository.SelectByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return likes, nil
}
