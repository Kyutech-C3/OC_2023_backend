package like

import (
	"net/http"
	"oc-2023/pkg/domain/entity"
	"oc-2023/pkg/usecase"

	"github.com/gin-gonic/gin"
)

type LikeHandler interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type likeHandler struct {
	interactor usecase.Interactor
}

func New(interactor usecase.Interactor) LikeHandler {
	return &likeHandler{
		interactor,
	}
}

func (lh *likeHandler) Create(ctx *gin.Context) {
	input := new(entity.PostLikes)

	if err := ctx.ShouldBindJSON(input); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Parse failed.",
		})
		return
	}
	if err := lh.interactor.CreateLike(ctx, input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Insert failed",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success POST",
		})
	}
}

func (lh *likeHandler) Delete(ctx *gin.Context) {
	input := new(entity.DeleteLikes)

	if err := ctx.ShouldBindJSON(input); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Parse failed.",
		})
		return
	}

	if err := lh.interactor.DeleteLike(ctx, input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Eject failed",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success DELETE",
		})
	}
}
