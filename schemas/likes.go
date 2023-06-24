package schemas

type CreateLikes struct {
	UserId string `json:"user_id"`
	PostId string `json:"post_id"`
}