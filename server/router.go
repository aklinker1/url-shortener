package server

import (
	"context"
	"fmt"
	"io/fs"
	"net/http"
	"strconv"
	"strings"

	"github.com/aklinker1/url-shortener/server/controllers"
	"github.com/aklinker1/url-shortener/server/models"
	"github.com/aklinker1/url-shortener/server/repos"
	"github.com/aklinker1/url-shortener/server/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func createRouter(ui *fs.FS, metaJSON string) *chi.Mux {
	r := chi.NewRouter()
	if !IS_PROD {
		r.Use(cors)
	}
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	if !IS_PROD {
		r.Use(middleware.Logger)
	}
	r.Use(middleware.Recoverer)

	r.Get("/", controllers.RootEndpoint())

	r.Get("/@/api/health", controllers.HealthEndpoint(metaJSON))

	r.Route("/@/api/urlEntries", func(r chi.Router) {
		r.With(paginate).Get("/", controllers.ListURLEntries())
		r.Post("/", controllers.CreateURLEntry())

		r.Route(fmt.Sprintf("/{%s}", utils.URL_ENTRY_ID_PARAM), func(r chi.Router) {
			r.Use(urlEntryCtx)
			r.Get("/", controllers.GetURLEntry())
			r.Put("/", controllers.UpdateURLEntry())
			r.Delete("/", controllers.DeleteURLEntry())
		})
	})

	fileServer(r, "/@/ui", ui)

	r.Route("/{shortUrl:[a-zA-Z0-9]+}", func(r chi.Router) {
		r.With(shortURLCtx).Handle("/", controllers.Redirect())
	})

	return r
}

func fileServer(r chi.Router, public string, assets *fs.FS) {
	index, err := fs.ReadFile(*assets, "index.html")
	if err != nil {
		panic(err)
	}
	embeddedFs := http.StripPrefix(public, http.FileServer(http.FS(*assets)))

	if strings.ContainsAny(public, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	if public != "/" && public[len(public)-1] != '/' {
		r.Get(public, http.RedirectHandler(public+"/", 301).ServeHTTP)
		public += "/"
	}

	r.Get(public+"*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.ContainsRune(r.URL.Path, '.') {
			w.WriteHeader(200)
			w.Header().Add("Content-Type", "text/html")
			w.Write(index)
		} else {
			embeddedFs.ServeHTTP(w, r)
		}
	}))
}

func urlEntryCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		hashedID := chi.URLParam(req, utils.URL_ENTRY_ID_PARAM)
		id, err := strconv.ParseInt(hashedID, 32, 0)
		if err != nil {
			http.Error(res, "Failed to read hashed ID="+hashedID, 404)
			return
		}
		urlEntry, err := repos.URLEntryRepo.Read(id)
		if err != nil {
			http.Error(res, "URL Entry not found with id="+fmt.Sprint(id), 404)
			return
		}
		ctx := context.WithValue(req.Context(), utils.URL_ENTRY, urlEntry)
		next.ServeHTTP(res, req.WithContext(ctx))
	})
}

func shortURLCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		hashedID := chi.URLParam(req, utils.SHORT_URL_PATH_PARAM)
		id, err := strconv.ParseInt(hashedID, 32, 0)
		if err != nil {
			http.Error(res, "Failed to read hashed ID="+hashedID, 404)
			return
		}
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
		pagination := &models.Pagination{
			Page: 0,
			Size: 20,
		}
		pageStr := req.URL.Query().Get(utils.PAGE_QUERY_PARAM)
		if pageStr != "" {
			pageInt, err := strconv.Atoi(pageStr)
			if err != nil {
				http.Error(res, "'page' query param was not a number", http.StatusBadRequest)
				return
			}
			pagination.Page = pageInt
		}

		sizeStr := req.URL.Query().Get(utils.SIZE_QUERY_PARAM)
		if sizeStr != "" {
			sizeInt, err := strconv.Atoi(sizeStr)
			if err != nil {
				http.Error(res, "'size' query param was not a number", http.StatusBadRequest)
				return
			}
			pagination.Size = sizeInt
		}
		paginationCtx := context.WithValue(req.Context(), utils.PAGINATION, pagination)

		next.ServeHTTP(res, req.WithContext(paginationCtx))
	})
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(1, req.URL, 2, req.Host, 3, req.URL.Host, 4, req.URL.Hostname())
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Headers", strings.Join([]string{
			"Accept",
			"Authorization",
			"Content-Type",
		}, ","))
		res.Header().Set("Access-Control-Allow-Methods", strings.Join([]string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		}, ","))

		if req.Method == http.MethodOptions {
			res.WriteHeader(http.StatusOK)
		} else {
			next.ServeHTTP(res, req)
		}
	})
}
