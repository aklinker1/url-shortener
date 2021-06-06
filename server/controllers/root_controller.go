package controllers

import "net/http"

func RootEndpoint() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		http.Redirect(res, req, "/@/ui", http.StatusSeeOther)
	}
}
