package vo

type ThemeReq struct {
	Value string `json:"value"`
	Key   string `json:"key"`
	Type  string `json:"type"`
	Desc  string `json:"desc"`
	Hook  string `json:"hook"`
}
