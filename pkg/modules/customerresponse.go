package modules

type CustomerResponse struct {
	Status  bool `json:"status" binding:"required"`
	Message struct {
		ID           int    `json:"id"`
		PointName    string `json:"point_name"`
		Identifier   string `json:"identifier"`
		IsAccess     int    `json:"isAccess"`
		BundleID     int    `json:"bundleId"`
		ErrorMessage string `json:"errorMessage"`
	} `json:"message" binding:"required"`
}
