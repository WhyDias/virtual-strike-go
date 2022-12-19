package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"log"
	"strconv"
	"time"
	"virtual-strike-backend-go/pkg/models"
	"virtual-strike-backend-go/pkg/modules"
)

type TimeService struct{}

func NewTimeService() *TimeService {
	return &TimeService{}
}

func (t *TimeService) TimeLogic(jsonInput modules.TimeRequest) (code int, any modules.Response) {
	requestBodyBytes := new(bytes.Buffer)
	json.NewEncoder(requestBodyBytes).Encode(jsonInput)

	var request modules.TimeRequest
	json.Unmarshal(requestBodyBytes.Bytes(), &request)

	unixTime := time.Now().Unix() + 21600

	db, err := sql.Open("mysql", "admin:admin@tcp(localhost:3306)/virtual-strike")
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
		response.Message = strconv.Itoa(int(unixTime))
		return 200, response
	}
}
