package article

import (
	httpHelper "article/internal/delivery/http"
	"article/internal/entity/article"
	"article/pkg/response"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CheckRequirement(title string, content string, category string, status string) string {

	var checkResult []string

	validStatuses := map[string]bool{
		"publish": true,
		"draft":   true,
		"trash":   true,
	}

	loweCaseStatus := strings.ToLower(status)

	if len(title) < 20 {
		checkResult = append(checkResult, "Title minimal 20 karakter")
	}
	if len(content) < 200 {
		checkResult = append(checkResult, "Content minimal 200 karakter")
	}
	if len(category) < 3 {
		checkResult = append(checkResult, "Category minimal 3 karakter")
	}
	if !validStatuses[loweCaseStatus] {
		checkResult = append(checkResult, "Status antara 'publish', 'draft', 'thrash'")
	}

	result := strings.Join(checkResult, ", ")

	return result
}

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

	checkResult := CheckRequirement(article.Title, article.Content, article.Category, article.Status)
	if checkResult != "" {
		resp = httpHelper.ParseErrorCode(checkResult)
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, checkResult)
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

func (h *Handler) UpdateArticle(w http.ResponseWriter, r *http.Request) {

	var (
		err     error
		article article.PostsSlim
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

	err = json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		return
	}

	article.ID = id

	checkResult := CheckRequirement(article.Title, article.Content, article.Category, article.Status)
	if checkResult != "" {
		resp = httpHelper.ParseErrorCode(checkResult)
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, checkResult)
		return
	}

	err = h.articleSvc.UpdateArticle(ctx, article)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		return
	}

	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}
