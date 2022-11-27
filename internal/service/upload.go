package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
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

	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/DB_virtual")
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
				var response modules.Response
				response.Status = false
				response.Message = err.Error()
				logrus.Error(err.Error())
				return 500, response
			}
		}

		pathToFile := "/statistic_" + time.Now().Format("2006-01-02_3-4-5")
		write := ioutil.WriteFile(path+pathToFile, []byte(request.Data), 0700)
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
