package http

import (
	"errors"
	"log"
	"net/http"

	"article/pkg/response"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

func (s *Server) Handler() *mux.Router {
	r := mux.NewRouter()
	
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	r.HandleFunc("", defaultHandler).Methods("GET")
	r.HandleFunc("/", defaultHandler).Methods("GET")

	// Routes
	article := r.PathPrefix("/article").Subrouter()
	article.HandleFunc("", s.Article.CreateArticle).Methods("POST")
	article.HandleFunc("/{id:[0-9]+}", s.Article.GetArticleById).Methods("GET")
	article.HandleFunc("/{limit:[0-9]+}/{offset:[0-9]+}", s.Article.GetArticlePagination).Methods("GET")
	article.HandleFunc("/{id:[0-9]+}", s.Article.UpdateArticle).Methods("PUT")
	article.HandleFunc("/{id:[0-9]+}", s.Article.DeleteArticle).Methods("DELETE")

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	return r
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Example Service API"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	var (
		resp   *response.Response
		err    error
		errRes response.Error
	)
	resp = &response.Response{}
	defer resp.RenderJSON(w, r)

	err = errors.New("404 Not Found")

	if err != nil {
		// Error response handling
		errRes = response.Error{
			Code:   404,
			Msg:    "404 Not Found",
			Status: true,
		}

		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		resp.StatusCode = 404
		resp.Error = errRes
		return
	}
}
