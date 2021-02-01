package utils

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

func SendData(ctx context.Context, res http.ResponseWriter, data interface{}) {
	if ctx.Value(RESPONSE_TYPE) == "xml" {
		xmlData, err := xml.MarshalIndent(data, "", "  ")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/xml")
		res.Write(xmlData)
	} else {
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		res.Write(jsonData)
	}
}
