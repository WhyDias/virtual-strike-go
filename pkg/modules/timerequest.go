package modules

type TimeRequest struct {
	Identification string `json:"identification" binding:"required"`
}
