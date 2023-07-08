package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"oc-2023/config"
	"oc-2023/cruds"
	"oc-2023/schemas"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleGetWorks(ctx *gin.Context) {
	tagNames := ctx.Query("tag_names")
	res, err := http.Get(fmt.Sprintf("%s/api/v1/works?tag_names=%s", config.APIHost, tagNames))
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var works schemas.Works
	if err := json.Unmarshal(body, &works); err != nil {
		fmt.Println(err)
	}

	for i, w := range works.Works {
		uid, err := uuid.Parse(w.ID)
		if err != nil {
			fmt.Printf("Invalid UUID string: %v\n", err)
			return
		}
		likes, err := cruds.GerLikesByID(uid)
		if err != nil {
			fmt.Println(err)
		}
		works.Works[i].Likes = schemas.ConvertLikeModelToSchemas(likes)
	}

	ctx.JSON(http.StatusOK, &works)
}

func HandleGetWork(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := http.Get(fmt.Sprintf("%s/api/v1/works/%s", config.APIHost, id))
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var work schemas.Work
	if err := json.Unmarshal(body, &work); err != nil {
		fmt.Println(err)
	}

	uid, err := uuid.Parse(work.ID)
	if err != nil {
		fmt.Printf("Invalid UUID string: %v\n", err)
		return
	}

	likes, err := cruds.GerLikesByID(uid)
	if err != nil {
		fmt.Println(err)
	}

	comments, err := cruds.GetCommetsByID(uid)
	if err != nil {
		fmt.Println(err)
	}

	work.Likes = schemas.ConvertLikeModelToSchemas(likes)
	work.Comments = schemas.ConvertCommentModelToSchema(comments)

	ctx.JSON(http.StatusOK, &work)
}
