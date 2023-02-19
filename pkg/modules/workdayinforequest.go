package modules

type WorkDayInfoRequest struct {
	Identification      string `form:"identification" binding:"required"`
	IdentificationTarif string `form:"identificationTarif" binding:"required"`
}
