package processes

import (
	"net/http"
	"partiq/internal/shared"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	processes, err := h.svc.List(ctx)
	if err != nil {
		shared.RespondError(w, http.StatusInternalServerError, "failed to fetch processes")
		return
	}

	shared.RespondJSON(w, http.StatusOK, processes)
}
