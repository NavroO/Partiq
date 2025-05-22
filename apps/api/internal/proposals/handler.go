package proposals

import (
	"net/http"
	"partiq/internal/shared"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	proposals, err := h.svc.List(ctx)
	if err != nil {
		shared.RespondError(w, http.StatusInternalServerError, "failed to fetch proposals")
		return
	}

	shared.RespondJSON(w, http.StatusOK, proposals)
}

func (h *Handler) GetByProcessID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	processIDStr := chi.URLParam(r, "processID")
	if processIDStr == "" {
		shared.RespondError(w, http.StatusBadRequest, "missing process ID")
		return
	}

	processID, err := strconv.Atoi(processIDStr)
	if err != nil {
		shared.RespondError(w, http.StatusBadRequest, "invalid process ID")
		return
	}

	proposals, err := h.svc.ListByProcessID(ctx, processID)
	if err != nil {
		shared.RespondError(w, http.StatusInternalServerError, "failed to fetch proposals")
		return
	}

	shared.RespondJSON(w, http.StatusOK, proposals)
}
