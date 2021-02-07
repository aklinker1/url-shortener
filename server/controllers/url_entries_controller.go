package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/aklinker1/url-shortener/server/models"
	"github.com/aklinker1/url-shortener/server/repos"
	"github.com/aklinker1/url-shortener/server/utils"
)

func readInputUrlEntry(res http.ResponseWriter, req *http.Request) (*models.InputUrlEntry, error) {
	input := &models.InputUrlEntry{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	err = json.Unmarshal(body, input)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	if input.URL == "" {
		http.Error(res, "URL cannot be empty", http.StatusInternalServerError)
		return nil, err
	}

	return input, nil
}

func ListURLEntries() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		pagination := req.Context().Value(utils.PAGINATION).(*models.Pagination)
		entries, err := repos.URLEntryRepo.List(pagination)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.SendData(req.Context(), res, models.ToDTOs(entries))
	}
}

func SearchURLEnties() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		panic("SearchURLEnties is not implemented")
	}
}

func GetURLEntry() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		entry := req.Context().Value(utils.URL_ENTRY).(*models.URLEntry)
		utils.SendData(req.Context(), res, entry.ToDTO())
	}
}

func CreateURLEntry() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		input, err := readInputUrlEntry(res, req)
		if err != nil {
			return
		}

		entry, err := repos.URLEntryRepo.Create(input.URL)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		utils.SendData(req.Context(), res, entry.ToDTO())
	}
}

func UpdateURLEntry() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		input, err := readInputUrlEntry(res, req)
		originalEntry := req.Context().Value(utils.URL_ENTRY).(*models.URLEntry)
		if err != nil {
			return
		}

		updatedEntry, err := repos.URLEntryRepo.Update(originalEntry, input.URL)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		utils.SendData(req.Context(), res, updatedEntry.ToDTO())
	}
}

func DeleteURLEntry() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		entry := ctx.Value(utils.URL_ENTRY).(*models.URLEntry)
		err := repos.URLEntryRepo.Delete(entry)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.SendData(ctx, res, entry.ToDTO())
	}
}
