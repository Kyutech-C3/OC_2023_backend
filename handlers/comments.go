package handler

import (
	"fmt"
	"net/http"
	"oc-2023/cruds"
	"oc-2023/schemas"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandlePostComment(ctx *gin.Context) {
	input := new(schemas.PostComments)

	if err := ctx.ShouldBindJSON(input); err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Parse failed.",
		})
		return
	}
	commentId, _ := uuid.NewUUID()
	if err := cruds.CreateComment(commentId, input.WorkId, input.UserId, input.UserName, input.Comment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Insert failed",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success POST",
		})
	}
	return
}

func HandleDeleteComment(ctx *gin.Context) {
	input := new(schemas.DeleteComments)

	if err := ctx.ShouldBindJSON(input); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Parse failed",
		})
		return
	}

	if err := cruds.DeleteComment(input.CommentId, input.UserId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Eject failed",
		})
		fmt.Print(err)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success DELETE",
		})
	}
	return
}
