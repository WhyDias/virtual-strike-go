package service

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"virtual-strike-backend-go/pkg/models"
	"virtual-strike-backend-go/pkg/modules"
)

type CustomerService struct{}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (u *CustomerService) CustomerLogic(jsonInput modules.CustomerRequest) (code int, any modules.CustomerResponse) {
	requestBodyBytes := new(bytes.Buffer)
	json.NewEncoder(requestBodyBytes).Encode(jsonInput)

	var request modules.CustomerRequest
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

	rowForPoints, err := db.Query("SELECT points from users WHERE username = ?", request.Username)
	if err != nil {
		var response modules.CustomerResponse
		response.Status = false
		response.Message.ErrorMessage = err.Error()
		logrus.Error(err)
		return 500, response
	}

	var pointOwner string

	for rowForPoints.Next() {
		error := rowForPoints.Scan(&pointOwner)
		if error != nil {
			var response modules.CustomerResponse
			response.Status = false
			response.Message.ErrorMessage = error.Error()
			logrus.Error(err)
			return 500, response
		}
	}

	pointsDecode, err := base64.StdEncoding.DecodeString(pointOwner)
	if err != nil {
		panic(err)
	}

	defer rowForPoints.Close()

	var points models.Points

	req := db.QueryRow("SELECT * FROM points WHERE points.identifier = ?", string(pointsDecode)).Scan(&points.ID, &points.PointName, &points.Identifier, &points.IsAccess, &points.BundleID, &points.Owner)
	switch {
	case req == sql.ErrNoRows:
		var response modules.CustomerResponse
		response.Status = false
		response.Message.ErrorMessage = req.Error()
		logrus.Error(req.Error())
		return 500, response
	case req != nil:
		var response modules.CustomerResponse
		response.Status = false
		response.Message.ErrorMessage = req.Error()
		logrus.Error(req.Error())
		return 500, response
	default:
		if request.Username != points.Owner {
			var response modules.CustomerResponse
			response.Status = false
			response.Message.ErrorMessage = req.Error()
			logrus.Error(req.Error())
			return 500, response
		} else {
			var response modules.CustomerResponse
			response.Status = true
			response.Message.ID = points.ID
			response.Message.PointName = points.PointName
			response.Message.Identifier = points.Identifier
			response.Message.IsAccess = points.IsAccess
			response.Message.BundleID = points.BundleID
			return 200, response
		}
	}
}
