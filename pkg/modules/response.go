package modules

type Response struct {
	Status  bool   `json:"status" binding:"required"`
	Message string `json:"message" binding:"required"`
}
