package vo

type UserRes struct {
	ID        uint64 `json:"id"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	NickName  string `json:"nick_name"`
	HeaderImg string `json:"header_img"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Enable    int    `json:"enable"`
	WebLink   string `json:"web_link"`
	Desc      string `json:"desc"`
}
