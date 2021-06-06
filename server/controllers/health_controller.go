package controllers

import (
	"net/http"
	"encoding/json"

	"github.com/aklinker1/url-shortener/server/utils"
)

func HealthEndpoint(metaJSON string) http.HandlerFunc {
	meta := map[string]interface{}{}
	err := json.Unmarshal([]byte(metaJSON), &meta)
	if err != nil {
		panic(err)
	}

	return func(res http.ResponseWriter, req *http.Request) {
		utils.SendData(req.Context(), res, map[string]interface{}{
			"server": "up",
			"version": meta["version"],
		})
	}
}
