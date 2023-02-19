package modules

import "time"

type UploadRequest struct {
	Identification string `form:"identification"`
	Data           struct {
		IdentifierTariff string    `form:"identifier-tariff"`
		StartWorkDate    time.Time `form:"StartWorkDate"`
		EndWorkDate      string    `form:"EndWorkDate"`
		Tariffs          []struct {
			GetTimeCreate   string  `form:"GetTimeCreate"`
			GetTimeExpired  string  `form:"GetTimeExpired"`
			GetTimeUsed     float64 `form:"GetTimeUsed"`
			GetIdentifier   string  `form:"GetIdentifier"`
			GetTitle        string  `form:"GetTitle"`
			GetCost         float64 `form:"GetCost"`
			GetTime         float64 `form:"GetTime"`
			GetTimeLeft     float64 `form:"GetTimeLeft"`
			GetTariffStatus int     `form:"GetTariffStatus"`
		} `json:"Tariffs"`
	} `json:"data"`
}
