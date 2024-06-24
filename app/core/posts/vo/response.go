package vo

type PageRes[T any] struct {
	Page
	// 总页数
	TotalPages int64 `json:"total"`
	// 总数量
	TotalCount int64 `json:"count"`
	List       []T   `json:"list"`
}
