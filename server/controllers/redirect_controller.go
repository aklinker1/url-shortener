package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aklinker1/url-shortener/server/models"
	"github.com/aklinker1/url-shortener/server/repos"
	"github.com/aklinker1/url-shortener/server/utils"
)

func Redirect() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		hashedID := req.Context().Value(utils.SHORT_URL_PATH_PARAM).(string)
		id, err := strconv.ParseInt(hashedID, 32, 0)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		entry := req.Context().Value(utils.URL_ENTRY).(*models.URLEntry)
		fmt.Println("id =", id, "entry =", entry) // TODO: remove
		if entry == nil {
			http.Redirect(res, req, fmt.Sprintf("/ui?badId=%s", id), http.StatusSeeOther)
		} else {
			repos.URLEntryRepo.UpdateVisits(entry)
			http.Redirect(res, req, entry.URL, http.StatusFound)
		}
	}
}
