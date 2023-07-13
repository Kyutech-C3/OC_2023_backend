package work

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"oc-2023/pkg/config"
	"oc-2023/pkg/domain/entity"
	"oc-2023/pkg/usecase"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WorkHandler interface {
	Index(ctx *gin.Context)
	GetByID(ctx *gin.Context)
}

type workHandler struct {
	interactor usecase.Interactor
}

func New(interactor usecase.Interactor) WorkHandler {
	return &workHandler{
		interactor,
	}
}

func (wh *workHandler) Index(ctx *gin.Context) {
	tagNames := ctx.Request.URL.Query().Get("tag_names")
	searchWord := ctx.Request.URL.Query().Get("search_word")

	baseURL := fmt.Sprintf("%s/api/v1/works", config.APIHost)
	params := make([]string, 0)

	if tagNames != "" {
		params = append(params, "tag_names="+url.QueryEscape(tagNames))
	}

	if searchWord != "" {
		params = append(params, "search_word="+url.QueryEscape(searchWord))
	}

	if len(params) > 0 {
		baseURL = baseURL + "?" + strings.Join(params, "&")
	}

	fmt.Println(baseURL)
	res, err := http.Get(baseURL)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var works entity.Works
	if err := json.Unmarshal(body, &works); err != nil {
		fmt.Println(err)
	}

	for i, w := range works.Works {
		uid, err := uuid.Parse(w.ID)
		if err != nil {
			fmt.Printf("Invalid UUID string: %v\n", err)
			return
		}
		likes, err := wh.interactor.GetLikesByID(ctx, uid)
		if err != nil {
			fmt.Println(err)
		}
		works.Works[i].Likes = likes
	}

	ctx.JSON(http.StatusOK, &works)
}

func (wh *workHandler) GetByID(ctx *gin.Context) {
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

	var work entity.Work
	if err := json.Unmarshal(body, &work); err != nil {
		fmt.Println(err)
	}

	uid, err := uuid.Parse(work.ID)
	if err != nil {
		fmt.Printf("Invalid UUID string: %v\n", err)
		return
	}

	likes, err := wh.interactor.GetLikesByID(ctx, uid)
	if err != nil {
		fmt.Println(err)
	}

	comments, err := wh.interactor.GetCommentsByID(ctx, uid)
	if err != nil {
		fmt.Println(err)
	}

	work.Likes = likes
	work.Comments = comments

	ctx.JSON(http.StatusOK, &work)
}
