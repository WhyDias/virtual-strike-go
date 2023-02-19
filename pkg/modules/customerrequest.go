package modules

type CustomerRequest struct {
	Username string `form:"username" binding:"required"`
}
