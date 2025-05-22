package processes

import (
	"database/sql"
	"errors"
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

	processes, err := h.svc.List(ctx)
	if err != nil {
		shared.RespondError(w, http.StatusInternalServerError, "failed to fetch processes")
		return
	}

	shared.RespondJSON(w, http.StatusOK, processes)
}

func (h *Handler) GetProcessByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStr := chi.URLParam(r, "processID")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		shared.RespondError(w, http.StatusBadRequest, "invalid process ID")
		return
	}

	process, err := h.svc.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			shared.RespondError(w, http.StatusNotFound, "process not found")
		} else {
			shared.RespondError(w, http.StatusInternalServerError, "failed to fetch process")
		}
		return
	}

	shared.RespondJSON(w, http.StatusOK, process)
}
