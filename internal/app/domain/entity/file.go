package entity

// FileResponse represents the standard response format for an uploaded file
type FileResponse struct {
	URL      string `json:"url"`
	Filename string `json:"filename"`
}
