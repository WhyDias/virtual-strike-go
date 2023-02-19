package modules

type PointRequest struct {
	Username string `form:"username" binding:"required"`
}
