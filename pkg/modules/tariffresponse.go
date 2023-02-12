package modules

type TariffResponse struct {
	Owner            string `json:"owner"`
	IdentifierTariff string `json:"identifier_tariff"`
	ErrorMessage     string `json:"errorMessage"`
}
