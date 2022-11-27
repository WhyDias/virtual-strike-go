package modules

type UploadRequest struct {
	Identification string `json:"identification" binding:"required"`
	Data           string `json:"data" binding:"required"`
}
