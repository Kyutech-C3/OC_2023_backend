// workIdはフロントのパスパラメータから、userIdはフロントのuuidから
package handler

import (
	"net/http"
	"oc-2023/cruds"
	"oc-2023/schemas"

	"github.com/gin-gonic/gin"
)

// gin.Contextは、リクエストの (GETされた, PUTした) データやパラメータ、またはエラー情報などいろんなものを含んだもの
// gin.Hはmap[string]interface{}と同義。
func HandlePostLike(ctx *gin.Context) {
	input := new(schemas.CreateLikes)

	if err := ctx.ShouldBindJSON(input); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Parse failed.",
		})
		return
	}

	if err := cruds.CreateLike(input.WorkId, input.UserId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Insert failed",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success POST",
		})
	}
}

func HandleDeleteLike(ctx *gin.Context) {
	input := new(schemas.DeleteLikes)

	if err := ctx.ShouldBindJSON(input); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Parse failed.",
		})
		return
	}

	if err := cruds.DeleteLike(input.WorkId, input.UserId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Eject failed",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success DELETE",
		})
	}
}
