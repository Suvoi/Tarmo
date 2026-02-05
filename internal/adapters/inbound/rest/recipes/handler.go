package recipes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tarmo/internal/adapters/inbound/rest/common"
	"tarmo/internal/core/recipes"
	"tarmo/internal/core/recipes/ports/inbound"

	"github.com/go-chi/chi/v5"
)

var (
	ErrInternal       = "internal error"
	ErrInvalidId      = "invalid id"
	ErrInvalidReq     = "invalid request"
	ErrInvalidRecipe  = "invalid recipe"
	ErrRecipeNotFound = "recipe not found"
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
		http.Error(w, ErrInternal, http.StatusInternalServerError)
		return
	}

	common.WriteJSON(w, http.StatusOK, ToResponseList(recipes))
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, ErrInvalidId, http.StatusBadRequest)
		return
	}

	recipe, err := h.service.GetByID(id)
	if err != nil {
		switch err {
		case recipes.ErrRecipeNotFound:
			http.Error(w, ErrRecipeNotFound, http.StatusNotFound)
		default:
			http.Error(w, ErrInternal, http.StatusInternalServerError)
		}
		return
	}

	common.WriteJSON(w, http.StatusOK, ToResponse(&recipe))
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateRecipeRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, ErrInvalidReq, http.StatusBadRequest)
		return
	}

	_, err := h.service.Create(ToCreateCommand(req))
	if err != nil {
		http.Error(w, ErrInvalidRecipe, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, ErrInvalidReq, http.StatusBadRequest)
		return
	}

	var req UpdateRecipeRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, ErrInvalidReq, http.StatusBadRequest)
		return
	}

	cmd := ToUpdateCommand(req)
	cmd.ID = id

	err = h.service.Update(cmd)
	if err != nil {
		switch err {
		case recipes.ErrRecipeNotFound:
			http.Error(w, ErrRecipeNotFound, http.StatusNotFound)
		default:
			http.Error(w, ErrInternal, http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, ErrInvalidId, http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		switch err {
		case recipes.ErrRecipeNotFound:
			http.Error(w, ErrRecipeNotFound, http.StatusNotFound)
		default:
			http.Error(w, ErrInternal, http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
