package modules

type WorkDayInfoRequest struct {
	Identification      string `json:"identification" binding:"required"`
	IdentificationTarif string `json:"identificationTarif" binding:"required"`
}
