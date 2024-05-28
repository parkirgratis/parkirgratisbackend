package controller

import (
	"net/http"

	"github.com/gocroot/helper"
	"github.com/gocroot/model"
	
)

func GetHome(respw http.ResponseWriter, req *http.Request) {
	var park model.Response
	park.Response = helper.GetIPaddress()
	helper.WriteJSON(respw, http.StatusOK, park)
}


func NotFound(respw http.ResponseWriter, req *http.Request) {
	var park model.Response
	park.Response = "Not Found"
	helper.WriteJSON(respw, http.StatusNotFound, park)
}
