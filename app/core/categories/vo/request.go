package vo

type CategoryReq struct {
	ParentId uint   `json:"parent_id" form:"parent_id"`
	SortBy   int    `json:"sort_by" form:"sort_by"`
	Name     string `json:"name" form:"name"`
	Slug     string `json:"slug" form:"slug"`
	Thumb    string `json:"thumb" form:"thumb"`
	Password string `json:"password" form:"password"`
	Desc     string `json:"desc" form:"desc"`
}
