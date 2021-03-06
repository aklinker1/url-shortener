package controllers

import (
	"fmt"
	"net/http"

	"github.com/aklinker1/url-shortener/server/models"
	"github.com/aklinker1/url-shortener/server/repos"
	"github.com/aklinker1/url-shortener/server/utils"
)

func Redirect() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		id := req.Context().Value(utils.SHORT_URL_PATH_PARAM).(int64)
		entry := req.Context().Value(utils.URL_ENTRY).(*models.URLEntry)
		fmt.Println("id =", id, "entry =", entry) // TODO: remove
		if entry == nil {
			http.Redirect(res, req, fmt.Sprintf("/@/ui?badId=%v", id), http.StatusSeeOther)
		} else {
			_, err := repos.URLEntryRepo.UpdateVisits(entry)
			if err != nil {
				fmt.Printf("Failed to update visits for ID=%v (%v)\n", id, err)
			}
			http.Redirect(res, req, entry.URL, http.StatusFound)
		}
	}
}
