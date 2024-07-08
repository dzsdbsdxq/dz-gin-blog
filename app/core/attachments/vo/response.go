package vo

import "time"

type AttachmentsUploadSuccessResponse struct {
	Url       string    `json:"url"`
	Path      string    `json:"path"`
	Driver    string    `json:"driver"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Key       string    `json:"key"`
}
