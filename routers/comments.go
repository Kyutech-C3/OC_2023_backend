package routers

import (
	"oc-2023/cruds"

	"net/http"

	"github.com/gin-gonic/gin"
)

func InitCommentRouter(cr *gin.RouterGroup) {
	cr.POST("/:post_id", createMyComment)
	cr.DELETE("/:comment_id", deleteComment)
}

func createMyComment(ctx *gin.Context) {
	var (
		// payload types.NewComment
		// userId any
		isExist bool
	)
	if /* userId, */ isExist = ctx.Get( /*"user_id" */ ); !isExist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "token is invalid",
		})
		return

	}

	workId := ctx.Param("work_id")
	//c.Bind(&payload)
	/* p, */
	err := cruds.CreateComment( /* payload.Content, userId.(string), */ workId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "cannot create your comment",
		})
		return
	}

	ctx.JSON(http.StatusOK /* p */)
}

func deleteComment(ctx *gin.Context) {
	var (
		isExist bool
	)

	if _, isExist = ctx.Get( /* "user_id" */ ); !isExist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "token is invalid",
		})
		return
	}

	commentId := ctx.Param("comment_id")
	err := cruds.DeleteComment(commentId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
