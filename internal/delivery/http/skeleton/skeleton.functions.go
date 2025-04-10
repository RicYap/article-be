package skeleton

import (
	"log"
	"net/http"
	httpHelper "skeleton/internal/delivery/http"
	"skeleton/pkg/response"
)

// GetSkeleton godoc
// @Summary Get entries of all skeletons
// @Description Get entries of all skeletons
// @Tags Skeleton
// @Accept  json
// @Produce  json
// @Security BearerAuth
// @Success 200
// @Router /v1/profiles [get]
func (h *Handler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	result, err := h.skeletonSvc.GetAllUser(ctx)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		return
	}

	resp.Data = result
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}

func (h *Handler) GeneratePDF(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	ctx := r.Context()

	result, err := h.skeletonSvc.GeneratePDF(ctx)
	if err != nil {
		defer resp.RenderJSON(w, r)
		
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "inline; filename=Daftar_User.pdf")
	w.Write(result)

	resp.Data = result
	log.Printf("[INFO] %s %s\n", r.Method, r.URL)
}