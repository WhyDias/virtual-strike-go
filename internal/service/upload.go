package service

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lithammer/shortuuid"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"os"
	"time"
	"virtual-strike-backend-go/pkg/models"
	"virtual-strike-backend-go/pkg/modules"
)

type UploadService struct{}

func NewUploadService() *UploadService {
	return &UploadService{}
}

func (u *UploadService) UploadLogic(jsonInput modules.UploadRequest) (code int, any modules.Response) {
	requestBodyBytes := new(bytes.Buffer)
	json.NewEncoder(requestBodyBytes).Encode(jsonInput)

	var request modules.UploadRequest
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

	var points models.Points

	req := db.QueryRow("SELECT * FROM points WHERE points.identifier = ?", request.Identification).Scan(&points.ID, &points.PointName, &points.Identifier, &points.IsAccess, &points.BundleID, &points.Owner)
	switch {
	case req == sql.ErrNoRows:
		var response modules.Response
		response.Status = false
		response.Message = req.Error()
		logrus.Error(req.Error())
		return 500, response
	case req != nil:
		var response modules.Response
		response.Status = false
		response.Message = req.Error()
		logrus.Error(req.Error())
		return 500, response
	default:
		path := "./" + request.Identification
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(path, os.ModePerm)
			if err != nil {
				var response modules.Response
				response.Status = false
				response.Message = err.Error()
				logrus.Error(err.Error())
				return 500, response
			}
		}

		var isAccess int
		requestForAccess := db.QueryRow("SELECT IF((SELECT count(*) as cnt FROM statistics WHERE identifier = ? and date = ?) > 3, 1, 0)", request.Identification, time.Now().Format("2006-01-02")).Scan(&isAccess)
		switch {
		case requestForAccess == sql.ErrNoRows:
			var response modules.Response
			response.Status = false
			response.Message = requestForAccess.Error()
			logrus.Error(requestForAccess.Error())
			return 500, response
		case requestForAccess != nil:
			var response modules.Response
			response.Status = false
			response.Message = requestForAccess.Error()
			logrus.Error(requestForAccess.Error())
			return 500, response
		default:
			if isAccess == 0 {

				pathToFile := "/statistic_" + time.Now().Format("2006-01-02_3-4-5")
				data, err := json.Marshal(request.Data)
				if err != nil {
					panic(err)
				}
				saveId := shortuuid.New()
				saveQuery := "INSERT INTO statistics (`id`, `date`, `data`, `identifier`) VALUES (?, ?, ?, ?)"
				saveResult, err := db.ExecContext(context.Background(), saveQuery, saveId, time.Now().Format("2006-01-02"), data, request.Identification)
				if err != nil {
					logrus.Fatalf("impossible insert data: %s", err)
				}
				idForSave, err := saveResult.LastInsertId()
				if err != nil {
					logrus.Fatalf("impossible to retrieve last inserted id: %s", err)
				}
				logrus.Printf("inserted id: %d, %s", idForSave, saveId)

				write := ioutil.WriteFile(path+pathToFile, data, 0700)
				if write != nil {
					var response modules.Response
					response.Status = false
					response.Message = write.Error()
					logrus.Error(write.Error())
					return 500, response
				}

				idPoint := shortuuid.New()
				var Data modules.Data
				s := request.Data
				a := []byte(s)
				json.Unmarshal(a, &Data)
				queryForPoint := "INSERT INTO point (`id`, `owner`, `identifier_tariff`, `StartWorkDate`, `EndWorkDate`) VALUES (?, ?, ?, ?, ?)"
				insertResult, err := db.ExecContext(context.Background(), queryForPoint, idPoint, request.Identification, Data.IdentifierTariff, Data.StartWorkDate, Data.EndWorkDate)
				if err != nil {
					var response modules.Response
					response.Status = false
					response.Message = err.Error()
					logrus.Error(err.Error())
					return 500, response
				}
				idForPoints, err := insertResult.LastInsertId()
				if err != nil {
					var response modules.Response
					response.Status = false
					response.Message = err.Error()
					logrus.Error(err.Error())
					return 500, response
				}
				logrus.Printf("inserted id: %d, %s", idForPoints, idPoint)

				tariffs := Data.Tariffs

				for _, tariff := range tariffs {
					idPointTariff := shortuuid.New()
					queryForPointTariff := "INSERT INTO point_tariffs (`point_id`,`id`, `GetTimeCreate`, `GetTimeExpired`, `GetTimeUsed`, `GetIdentifier`, `GetTitle`, `GetCost`, `GetTime`, `GetTimeLeft`, `GetTariffStatus`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
					insertResult, err := db.ExecContext(context.Background(), queryForPointTariff, idPoint, idPointTariff, tariff.GetTimeCreate, tariff.GetTimeExpired, tariff.GetTimeUsed, tariff.GetIdentifier, tariff.GetTitle, tariff.GetCost, tariff.GetTime, tariff.GetTimeLeft, tariff.GetTariffStatus)
					if err != nil {
						var response modules.Response
						response.Status = false
						response.Message = err.Error()
						logrus.Error(err.Error())
						return 500, response
					}
					idForPointsTariff, err := insertResult.LastInsertId()
					if err != nil {
						var response modules.Response
						response.Status = false
						response.Message = err.Error()
						logrus.Error(err.Error())
						return 500, response
					}
					logrus.Printf("inserted id: %d, %s", idForPointsTariff, idPointTariff)

				}

				var response modules.Response
				response.Status = true
				response.Message = "Create successfully"
				return 200, response
			} else {
				var response modules.Response
				response.Status = false
				response.Message = "Database have more than 3 record on this date"
				return 500, response
			}
		}
	}
}
