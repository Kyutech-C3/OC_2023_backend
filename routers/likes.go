package routers

import (
	"net/http"
	"oc-2023/cruds"

	"github.com/gin-gonic/gin"
)

 func InitLikeRouter(lr *gin.RouterGroup) {
 	lr.POST("/:work_id", itIsLike)
 	lr.DELETE("/:work_id", itIsNotLike)
 }

// *gin.Context: リクエストの (GETされた, PUTした) データやパラメータ、またはエラー情報などいろんなものを含んだもの
func itIsLike(ctx *gin.Context) {
	var (
		/* userId any */
		isExist bool
	)

	// URIの 「/:work_id」 に入る部分を workIdに代入する
	 workId := ctx.Param("work_id")

	// user_idを得られなかったらmessageをJSONで出力する例外処理
	if /* userId, */ isExist = ctx.Get( /* "user_id" */); !isExist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":"token is invalid",
		})
		return
	}

	// エラーがあれば、messageでエラー内容をJSONで出力する例外処理
	work, err := cruds.GiveLike(workId /*, userId string */)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":err.Error(),
		})
		return
	}

	// 成功した場合にStatusOKを返す
	ctx.JSON(http.StatusOK, work )
}

func itIsNotLike(ctx *gin.Context) {
	var (
		/* userId any*/
		isExist bool
	)

	// URIの 「/:work_id」 に入る部分を workIdに代入する
	/* workId := ctx.Param("work_id") */
	
	// user_idを得られなかったらmessageをJSONで出力する例外処理
	if /* userId, */ isExist = ctx.Get(/* "user_id" */); !isExist {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":"token is invalid",
		})
		return
	} 

	// エラーがあれば、messageでエラー内容をJSONで出力する例外処理
	 work, err := cruds.DeleteLike(workId ,/* userId string*/)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message":err.Error(),
		})
		return
	}

	// 成功した場合にStatusOKを返す
	ctx.JSON(http.StatusOK, work )
}
