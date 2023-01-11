package service

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"log"
	"strings"
	"virtual-strike-backend-go/pkg/modules"
)

type PointService struct{}

func NewPointService() *PointService {
	return &PointService{}
}

func (p *PointService) PointLogic(jsonInput modules.PointRequest) (code int, any []modules.PointResponse) {
	requestBodyBytes := new(bytes.Buffer)
	json.NewEncoder(requestBodyBytes).Encode(jsonInput)

	var request modules.PointRequest
	json.Unmarshal(requestBodyBytes.Bytes(), &request)

	var res []modules.PointResponse

	db, err := sql.Open("mysql", "admin:admin@tcp(localhost:3306)/virtual-strike")
	if err != nil {
		log.Print(err.Error())
		return 500, res
	}

	rowForPoints, err := db.Query("SELECT points from customers WHERE username = ?", request.Username)
	if err != nil {
		log.Print(err.Error())
		return 500, res
	}

	var points string

	for rowForPoints.Next() {
		error := rowForPoints.Scan(&points)
		if error != nil {
			log.Println(error.Error())
			return 500, res
		}
	}

	pointsDecode, err := base64.StdEncoding.DecodeString(points)
	if err != nil {
		panic(err)
	}

	ownerPoints := strings.Split(string(pointsDecode), "|")

	defer rowForPoints.Close()

	for _, ownerPoint := range ownerPoints {
		rows, err := db.Query("SELECT identifier_tariff, id, StartWorkDate, EndWorkDate FROM point WHERE owner = ?", ownerPoint)
		if err != nil {
			log.Print(err.Error())
			return 500, res
		}

		var info modules.PointResponse

		for rows.Next() {
			var IdentifierTariff, Id, StartWorkDate, EndWorkDate string
			err = rows.Scan(&IdentifierTariff, &Id, &StartWorkDate, &EndWorkDate)
			if err != nil {
				log.Println(err.Error())
				return 500, res
			}

			info.IdentifierTariff = IdentifierTariff
			info.Id = Id
			info.StartWorkDate = StartWorkDate
			info.EndWorkDate = EndWorkDate

			var tariffs []modules.Tariff
			info.Tariffs = tariffs
			rows, err := db.Query("SELECT GetTimeCreate, GetTimeExpired, GetTimeUsed / 60, GetIdentifier, GetTitle, GetCost, GetTime, GetTimeLeft, GetTariffStatus FROM point_tariffs WHERE point_id = ?", info.Id)
			if err != nil {
				log.Print(err.Error())
				return 500, res
			}

			var tariff modules.Tariff

			for rows.Next() {
				var GetCost, GetTime, GetTimeLeft, GetTimeUsed float64
				var GetTariffStatus int
				var GetTimeCreate, GetTimeExpired, GetIdentifier, GetTitle string
				err = rows.Scan(&GetTimeCreate, &GetTimeExpired, &GetTimeUsed, &GetIdentifier, &GetTitle, &GetCost, &GetTime, &GetTimeLeft, &GetTariffStatus)
				if err != nil {
					log.Println(err.Error())
					return 500, res
				}

				tariff.GetTimeCreate = GetTimeCreate
				tariff.GetTimeExpired = GetTimeExpired
				tariff.GetTimeUsed = GetTimeUsed
				tariff.GetIdentifier = GetIdentifier
				tariff.GetTitle = GetTitle
				tariff.GetCost = GetCost
				tariff.GetTime = GetTime
				tariff.GetTimeLeft = GetTimeLeft
				tariff.GetTariffStatus = GetTariffStatus
				tariffs = append(tariffs, tariff)
			}

			info.Tariffs = tariffs
			res = append(res, info)
		}
	}

	defer db.Close()
	return 200, res
}
