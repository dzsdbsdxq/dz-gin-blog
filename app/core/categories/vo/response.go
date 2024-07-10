package vo

type MenuRes struct {
	ID       uint64    `json:"id"`
	ParentId uint      `json:"parent_id"`
	SortBy   int       `json:"sort_by"`
	Name     string    `json:"name"`
	Slug     string    `json:"slug"`
	Thumb    string    `json:"thumb"`
	Password string    `json:"password"`
	Desc     string    `json:"desc"`
	Url      string    `json:"url"`
	Children []MenuRes `json:"children"`
}
type MenuTreeMap map[string][]MenuRes
