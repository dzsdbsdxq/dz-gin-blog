package vo

type PostTagReq struct {
	PostId uint `json:"post_id"`
	TagId  uint `json:"tag_id"`
}
