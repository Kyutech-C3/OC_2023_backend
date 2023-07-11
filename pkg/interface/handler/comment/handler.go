package comment

import (
	"net/http"
	"oc-2023/pkg/domain/entity"
	"oc-2023/pkg/usecase"

	"github.com/gin-gonic/gin"
)

type CommentHandler interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type commentHandler struct {
	interactor usecase.Interactor
}

func New(interactor usecase.Interactor) CommentHandler {
	return &commentHandler{
		interactor,
	}
}

func (lh *commentHandler) Create(ctx *gin.Context) {
	input := new(entity.PostComments)

	if err := ctx.ShouldBindJSON(input); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Parse failed.",
		})
		return
	}
	if err := lh.interactor.CreateComment(ctx, input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Insert failed",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success POST",
		})
	}
}

func (lh *commentHandler) Delete(ctx *gin.Context) {
	input := new(entity.DeleteComments)

	if err := ctx.ShouldBindJSON(input); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Parse failed.",
		})
		return
	}

	if err := lh.interactor.DeleteComment(ctx, input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Eject failed",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success DELETE",
		})
	}
}
