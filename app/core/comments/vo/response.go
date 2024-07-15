package vo

type CommentsRes struct {
	UserType  int    `json:"user_type" form:"user_type"`
	CommentId uint   `json:"comment_id" form:"comment_id"`
	Status    int    `json:"status" form:"status"`
	Type      int    `json:"type" form:"type"`
	FromId    uint   `json:"from_id" form:"from_id"`
	NickName  string `json:"nick_name" form:"nick_name"`
	Email     string `json:"email" form:"email"`
	Link      string `json:"link" form:"link"`
	Body      string `json:"body" form:"body"`
	FromTitle string `json:"from_title" form:"from_title"`
}
