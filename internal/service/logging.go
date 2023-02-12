package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"os"
	"time"
	"virtual-strike-backend-go/pkg/models"
	"virtual-strike-backend-go/pkg/modules"
)

type LoggingService struct{}

func NewLoggingService() *LoggingService {
	return &LoggingService{}
}

func (u *LoggingService) LoggingLogic(jsonInput modules.LoggingRequest) (code int, any modules.Response) {
	requestBodyBytes := new(bytes.Buffer)
	json.NewEncoder(requestBodyBytes).Encode(jsonInput)

	var request modules.LoggingRequest
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

	req := db.QueryRow("SELECT * FROM points WHERE points.identifier = ?", request.Identification).Scan(&points.ID, &points.PointName, &points.Identifier, &points.IsAccess, &points.BundleID)
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
				pathLogs := "./logs"
				if _, _ = os.Stat(pathLogs); errors.Is(err, os.ErrNotExist) {
					errLog := os.Mkdir(pathLogs, os.ModePerm)
					if errLog != nil {
						var response modules.Response
						response.Status = false
						response.Message = err.Error()
						logrus.Error(err.Error())
						return 500, response
					}
				}
			}
		}

		pathToFile := "/logging_" + time.Now().Format("2006-01-02_3-4-5")
		data, err := json.Marshal(request.Data)
		if err != nil {
			panic(err)
		}

		write := ioutil.WriteFile(path+pathToFile, data, 0700)
		if write != nil {
			var response modules.Response
			response.Status = false
			response.Message = write.Error()
			logrus.Error(write.Error())
			return 500, response
		}

		var response modules.Response
		response.Status = true
		response.Message = "Create successfully"
		return 200, response
	}
}
