package vo

type PostReq struct {
	Title           string `json:"title"`
	Body            string `json:"body"`
	Extend          string `json:"extend"`
	Thumb           string `json:"thumb"`
	CommentAccepted int    `json:"comment_accepted"`
	Recommend       int    `json:"recommend"`
	Status          uint   `json:"status"`
	Top             uint   `json:"top"`
	SortBy          uint   `json:"sortby"`
	Flag            uint   `json:"flag"`     //文章类型(1=图文，2=视频)
	Views           uint   `json:"views"`    //查看数
	Likes           uint   `json:"likes"`    //点赞数
	Password        string `json:"password"` //密码
	Slug            string `json:"slug"`     //别名(可设置成路径访问)
	Desc            string `json:"desc"`     //描述
	SeoKey          string `json:"seo_key"`  //seo关键词
	SeoDesc         string `json:"seo_desc"` //seo描述
}
type PageInfo struct {
	PageNo   int64  `json:"pageNo" form:"pageNo"`     // 当前页
	PageSize int64  `json:"pageSize" form:"pageSize"` // 每页数量
	Category int64  `json:"category" form:"category"` //栏目ID
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}
