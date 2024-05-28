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
	park, _ := atdb.GetAllDoc[[]model.Tempat](config.Mongoconn, "tempat", bson.M{})
	helper.WriteJSON(respw, http.StatusOK, park)
}
