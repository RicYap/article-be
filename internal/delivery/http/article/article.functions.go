package article

import (
	httpHelper "article/internal/delivery/http"
	"article/internal/entity/article"
	"article/pkg/response"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateArticle(w http.ResponseWriter, r *http.Request) {

	var (
		err     error
		article article.Posts
	)

	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	err = json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		return
	}

	err = h.articleSvc.CreateArticle(ctx, article)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		return
	}

	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}

func (h *Handler) GetArticleById(w http.ResponseWriter, r *http.Request) {

	var (
		err error
	)

	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		return
	}

	result, err := h.articleSvc.GetArticleById(ctx, id)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		return
	}

	resp.Data = result
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}

func (h *Handler) GetArticlePagination(w http.ResponseWriter, r *http.Request) {

	var (
		err error
	)

	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	vars := mux.Vars(r)

	limit, err := strconv.Atoi(vars["limit"])
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		return
	}

	offset, err := strconv.Atoi(vars["offset"])
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		return
	}

	result, err := h.articleSvc.GetArticlePagination(ctx, limit, offset)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		return
	}

	resp.Data = result
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}