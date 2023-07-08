package schemas

import (
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02T15:04:05.999999"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

type User struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	DisplayName string     `json:"display_name"`
	AvatarUrl   string     `json:"avatar_url"`
	Profile     string     `json:"profile"`
	TwitterID   string     `json:"twitter_id"`
	GithubID    string     `json:"github_id"`
	CreatedAt   CustomTime `json:"created_at"`
	UpdatedAt   CustomTime `json:"updated_at"`
}

type Asset struct {
	AssetType string     `json:"asset_type"`
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	WorkID    string     `json:"work_id"`
	Extension string     `json:"extension"`
	Url       string     `json:"url"`
	CreatedAt CustomTime `json:"created_at"`
	UpdatedAt CustomTime `json:"updated_at"`
}

type Url struct {
	Url       string     `json:"url"`
	UrlType   string     `json:"url_type"`
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	CreatedAt CustomTime `json:"created_at"`
	UpdatedAt CustomTime `json:"updated_at"`
}

type Tag struct {
	Name  string `json:"name"`
	Color string `json:"color"`
	ID    string `json:"id"`
}

type Thumbnail struct {
	AssetType string     `json:"asset_type"`
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	WorkID    string     `json:"work_id"`
	Extension string     `json:"extension"`
	Url       string     `json:"url"`
	CreatedAt CustomTime `json:"created_at"`
	UpdatedAt CustomTime `json:"updated_at"`
}

type Work struct {
	ID              string     `json:"id"`
	Title           string     `json:"title"`
	Description     string     `json:"description"`
	DescriptionHtml string     `json:"description_html"`
	User            User       `json:"user"`
	Assets          []Asset    `json:"assets"`
	Urls            []Url      `json:"urls"`
	Visibility      string     `json:"visibility"`
	Tags            []Tag      `json:"tags"`
	Thumbnail       Thumbnail  `json:"thumbnail"`
	Likes           []Like     `json:"likes"`
	Comments        []Comment  `json:"comments"`
	CreatedAt       CustomTime `json:"created_at"`
	UpdatedAt       CustomTime `json:"updated_at"`
}

type Works struct {
	Works           []Work `json:"works"`
	WorksTotalCount int    `json:"works_total_count"`
}
