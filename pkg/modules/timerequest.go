package modules

type TimeRequest struct {
	Identification string `form:"identification" binding:"required"`
}
