package vo

type AttachmentsRes struct {
	Url       string `json:"url"`
	Path      string `json:"path"`
	Driver    string `json:"driver"`
	CreatedAt string `json:"created_at"`
	Name      string `json:"name"`
	Key       string `json:"key"`
	FileType  string `json:"file_type"`
}
