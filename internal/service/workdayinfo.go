package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"log"
	"virtual-strike-backend-go/pkg/models"
	"virtual-strike-backend-go/pkg/modules"
)

type WorkDayInfoService struct{}

func NewWorkDayInfoService() *WorkDayInfoService {
	return &WorkDayInfoService{}
}

func (u *WorkDayInfoService) WorkDayInfoLogic(jsonInput modules.WorkDayInfoRequest) (code int, any modules.Response) {
	requestBodyBytes := new(bytes.Buffer)
	json.NewEncoder(requestBodyBytes).Encode(jsonInput)

	var request modules.WorkDayInfoRequest
	json.Unmarshal(requestBodyBytes.Bytes(), &request)

	db, err := sql.Open("mysql", "admin:admin@tcp(31.172.64.249:3306)/virtual-strike")
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

		var response modules.Response
		response.Status = true
		response.Message = "Success"
		return 200, response
	}
}
