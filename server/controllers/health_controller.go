package controllers

import (
	"net/http"

	"github.com/aklinker1/url-shortener/server/utils"
)

func HealthEndpoint() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		utils.SendData(req.Context(), res, map[string]interface{}{
			"server": "up",
		})
	}
}
