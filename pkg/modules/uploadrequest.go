package modules

import "time"

type UploadRequest struct {
	Identification string `form:"identification"`
	Data           string `form:"data" binding:"required"`
}

type Data struct {
	IdentifierTariff string
	StartWorkDate    time.Time
	EndWorkDate      string
	Tariffs          []struct {
		GetTimeCreate   string
		GetTimeExpired  string
		GetTimeUsed     float64
		GetIdentifier   string
		GetTitle        string
		GetCost         float64
		GetTime         float64
		GetTimeLeft     float64
		GetTariffStatus int
	}
}
