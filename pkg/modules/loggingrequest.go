package modules

type LoggingRequest struct {
	Identification string `form:"identification" binding:"required"`
	Data           string `form:"data" binding:"required"`
}
