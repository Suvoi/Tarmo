package recipes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tarmo/internal/adapters/inbound/rest/common"
	"tarmo/internal/core/recipes/ports/inbound"
	"tarmo/internal/lib/logger"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service inbound.RecipePort
}

func NewHandler(service inbound.RecipePort) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.WriteJSON(w, http.StatusOK, ToResponseList(recipes))
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	recipe, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.WriteJSON(w, http.StatusOK, ToResponse(recipe))
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateRecipeRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Error("Failed to decode JSON: %v", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	logger.Info("Received request: %+v", req)

	if req.Name == "" {
		logger.Error("Name is empty")
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	err := h.service.Create(ToCommand(req))
	if err != nil {
		logger.Error("Service.Create failed: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
