package controller

import (
	"net/http"

	"github.com/gocroot/config"
	"github.com/gocroot/helper"
	"github.com/gocroot/helper/atdb"
	"github.com/gocroot/model"
	"go.mongodb.org/mongo-driver/bson"
	
)


func Getparkirgratis(respw http.ResponseWriter, req *http.Request) {
	park:=atdb.GetAllDoc[[]model .parkirgratis](config.Mongoconn,"tempat",bson.M{})
	park.Response = "Not Found"
	helper.WriteJSON(respw, http.StatusOk, park)
}
