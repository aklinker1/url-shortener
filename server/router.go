package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/aklinker1/url-shortener/server/controllers"
	"github.com/aklinker1/url-shortener/server/repos"
	"github.com/aklinker1/url-shortener/server/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func createRouter() *chi.Mux {
	r := chi.NewRouter()
	if !IS_PROD {
		r.Use(middleware.Logger)
	}
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/api/health", controllers.HealthEndpoint())

	r.Route("/api/urlEntries", func(r chi.Router) {
		r.With(paginate).Get("/", controllers.ListURLEntries())
		// r.With(paginate).Get("/search", controllers.SearchURLEnties())
		r.Post("/", controllers.CreateURLEntry())

		r.Route(fmt.Sprintf("/{%s}", utils.URL_ENTRY_ID_PARAM), func(r chi.Router) {
			r.Use(urlEntryCtx)
			r.Get("/", controllers.GetURLEntry())
			r.Put("/", controllers.UpdateURLEntry())
			r.Delete("/", controllers.DeleteURLEntry())
		})
	})

	fileServer(r, "/ui", "/app/ui")

	r.Route("/{shortUrl:[a-zA-Z0-9]+}", func(r chi.Router) {
		r.With(shortURLCtx).Handle("/", controllers.Redirect())
	})

	return r
}

func fileServer(r chi.Router, public string, static string) {

	if strings.ContainsAny(public, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	root, _ := filepath.Abs(static)
	if _, err := os.Stat(root); os.IsNotExist(err) {
		panic("Static Documents Directory Not Found")
	}

	fs := http.StripPrefix(public, http.FileServer(http.Dir(root)))

	if public != "/" && public[len(public)-1] != '/' {
		r.Get(public, http.RedirectHandler(public+"/", 301).ServeHTTP)
		public += "/"
	}

	r.Get(public+"*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file := strings.Replace(r.RequestURI, public, "/", 1)
		if _, err := os.Stat(root + file); os.IsNotExist(err) {
			http.ServeFile(w, r, path.Join(root, "index.html"))
			return
		}
		fs.ServeHTTP(w, r)
	}))
}

func urlEntryCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		urlEntryID := chi.URLParam(req, utils.URL_ENTRY_ID_PARAM)
		urlEntry, err := repos.URLEntryRepo.Read(urlEntryID)
		if err != nil {
			http.Error(res, "URL Entry not found with id="+urlEntryID, 404)
			return
		}
		ctx := context.WithValue(req.Context(), utils.URL_ENTRY, urlEntry)
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}

func shortURLCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		id := chi.URLParam(req, utils.SHORT_URL_PATH_PARAM)
		idCtx := context.WithValue(req.Context(), utils.SHORT_URL_PATH_PARAM, id)
		urlEntry, _ := repos.URLEntryRepo.Read(id)

		var ctx context.Context
		if urlEntry != nil {
			ctx = context.WithValue(idCtx, utils.URL_ENTRY, urlEntry)
		} else {
			ctx = idCtx
		}
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}

func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		var page = 0
		pageStr := req.URL.Query().Get(utils.PAGE_QUERY_PARAM)
		if pageStr != "" {
			var err error
			page, err = strconv.Atoi(pageStr)
			if err != nil {
				http.Error(res, "'page' query param was not a number", http.StatusBadRequest)
			}
		}
		pageCtx := context.WithValue(req.Context(), utils.PAGE_QUERY_PARAM, page)

		var size = 20
		sizeStr := req.URL.Query().Get(utils.SIZE_QUERY_PARAM)
		if sizeStr != "" {
			var err error
			size, err = strconv.Atoi(sizeStr)
			if err != nil {
				http.Error(res, "'size' query param was not a number", http.StatusBadRequest)
			}
		}
		pageAndSizeCtx := context.WithValue(pageCtx, utils.SIZE_QUERY_PARAM, size)

		next.ServeHTTP(res, req.WithContext(pageAndSizeCtx))
	})
}
