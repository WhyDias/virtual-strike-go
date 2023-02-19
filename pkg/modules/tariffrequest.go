package modules

type TariffRequest struct {
	Date string `form:"date" binding:"required"`
}
