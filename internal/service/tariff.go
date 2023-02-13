package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"virtual-strike-backend-go/pkg/modules"
)

type TariffService struct{}

func NewTariffService() *TariffService {
	return &TariffService{}
}

func (u *TariffService) TariffLogic(jsonInput modules.TariffRequest) (code int, any []modules.TariffResponse) {
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

	var tariffs []modules.TariffResponse

	rows, err := db.Query("SELECT id, data FROM statistics WHERE date = ?", request.Date)
	if err != nil {
		log.Print(err.Error())
		return 500, tariffs
	}

	var tariff modules.TariffResponse

	for rows.Next() {
		var id, data string
		err = rows.Scan(&id, &data)
		if err != nil {
			log.Println(err.Error())
			return 500, tariffs
		}

		tariff.ID = id
		stri := strings.ReplaceAll(data, "/", "")
		fmt.Println(stri)
		fmt.Println(data)
		tariff.Data = data
		tariffs = append(tariffs, tariff)
	}
	return 200, tariffs
}
