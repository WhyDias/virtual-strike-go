package models

import "time"

type DailyInfo struct {
	Identification string    `json:"Identification"`
	StartWorkDate  time.Time `json:"StartWorkDate"`
	EndWorkDate    string    `json:"EndWorkDate"`
	Tariffs        []struct {
		GetTimeCreate   string  `json:"GetTimeCreate"`
		GetTimeExpired  string  `json:"GetTimeExpired"`
		GetTimeUsed     float64 `json:"GetTimeUsed"`
		GetIdentifier   string  `json:"GetIdentifier"`
		GetTitle        string  `json:"GetTitle"`
		GetCost         float64 `json:"GetCost"`
		GetTime         float64 `json:"GetTime"`
		GetTimeLeft     float64 `json:"GetTimeLeft"`
		GetTariffStatus int     `json:"GetTariffStatus"`
	} `json:"Tariffs"`
}
