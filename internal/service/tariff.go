package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"virtual-strike-backend-go/pkg/modules"
)

type TariffService struct{}

func NewTariffService() *TariffService {
	return &TariffService{}
}

func (u *TariffService) TariffLogic(jsonInput modules.TariffRequest) (code int, any modules.TariffResponse) {
	requestBodyBytes := new(bytes.Buffer)
	json.NewEncoder(requestBodyBytes).Encode(jsonInput)

	var request modules.TariffRequest
	json.Unmarshal(requestBodyBytes.Bytes(), &request)
	Driver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	db, err := sql.Open(Driver, DBURL)
	if err != nil {
		log.Print(err.Error())
	}

	defer db.Close()

	var tariff modules.TariffResponse

	req := db.QueryRow("SELECT owner, identifier_tariff FROM point WHERE StartWorkDate = ?", request.Data).Scan(&tariff.Owner, &tariff.IdentifierTariff)
	switch {
	case req == sql.ErrNoRows:
		var response modules.TariffResponse
		response.ErrorMessage = req.Error()
		logrus.Error(req.Error())
		return 500, response
	case req != nil:
		var response modules.TariffResponse
		response.ErrorMessage = req.Error()
		logrus.Error(req.Error())
		return 500, response
	default:
		var response modules.TariffResponse
		response.Owner = tariff.Owner
		response.IdentifierTariff = tariff.IdentifierTariff
		return 200, response
	}
}
