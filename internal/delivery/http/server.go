package http

import (
	"net/http"

	"article/pkg/grace"

	"github.com/rs/cors"
)

// ArticleHandler ...
type ArticleHandler interface {
	CreateArticle(w http.ResponseWriter, r *http.Request)
	GetArticleById(w http.ResponseWriter, r *http.Request)
	GetArticlePagination(w http.ResponseWriter, r *http.Request)
	UpdateArticle(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	Article ArticleHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	handler := cors.AllowAll().Handler(s.Handler())
	return grace.Serve(port, handler)
}
