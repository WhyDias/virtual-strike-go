package modules

type PointResponse struct {
	IdentifierTariff string   `json:"identifier-tariff"`
	Id               string   `json:"id"`
	StartWorkDate    string   `json:"StartWorkDate"`
	EndWorkDate      string   `json:"EndWorkDate"`
	Tariffs          []Tariff `json:"Tariffs"`
}

type Tariff struct {
	GetTimeCreate   string  `json:"GetTimeCreate"`
	GetTimeExpired  string  `json:"GetTimeExpired"`
	GetTimeUsed     float64 `json:"GetTimeUsed"`
	GetIdentifier   string  `json:"GetIdentifier"`
	GetTitle        string  `json:"GetTitle"`
	GetCost         float64 `json:"GetCost"`
	GetTime         float64 `json:"GetTime"`
	GetTimeLeft     float64 `json:"GetTimeLeft"`
	GetTariffStatus int     `json:"GetTariffStatus"`
}
