package vo

type TagsRes struct {
	Nums  int    `json:"nums"`
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Thumb string `json:"thumb"`
}
