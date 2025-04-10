package article

import (
	httpHelper "article/internal/delivery/http"
	"article/internal/entity/article"
	"article/pkg/response"
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	
	var (
		err error
		article article.Post
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
