package modules

type PointRequest struct {
	Username string `json:"username" binding:"required"`
}
