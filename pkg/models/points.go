package models

type Points struct {
	ID         int    `json:"id"`
	PointName  string `json:"point_name"`
	Identifier string `json:"identifier"`
	IsAccess   int    `json:"isAccess"`
	BundleID   int    `json:"bundleId"`
	Owner      string `json:"owner"`
}
