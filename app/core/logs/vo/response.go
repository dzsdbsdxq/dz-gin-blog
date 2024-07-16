package vo

import "time"

type LogsRes struct {
	ID        uint64    `json:"ID"`         //主键ID
	CreatedAt time.Time `json:"created_at"` // 创建时间
	Type      string    `json:"type"`
	Key       string    `json:"key"`
	Content   string    `json:"content"`
	Ip        string    `json:"ip"`
	Creator   string    `json:"creator"`
}
