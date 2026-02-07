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

// GetAll recipes
// @Summary      List all recipes
// @Description  Get a list of all recipes in the collection
// @Tags         recipes
// @Produce      json
// @Success      200  {array}   RecipeListJSONResponseDTO
// @Failure      500  {string}  string "internal error"
// @Router       /recipes [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	recipes, err := h.service.GetAll()
	if err != nil {
		http.Error(w, ErrInternal, http.StatusInternalServerError)
		return
	}

	common.WriteJSON(w, http.StatusOK, ToResponseList(recipes))
}

// GetByID recipe
// @Summary      Get a recipe by ID
// @Description  Get detailed information about a single recipe
// @Tags         recipes
// @Produce      json
// @Param        id   path      int  true  "Recipe ID"
// @Success      200  {object}  RecipeJSONResponseDTO
// @Failure      400  {string}  string "invalid id"
// @Failure      404  {string}  string "recipe not found"
// @Failure      500  {string}  string "internal error"
// @Router       /recipes/{id} [get]
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

// Create recipe
// @Summary      Create a new recipe
// @Description  Add a new recipe to the collection
// @Tags         recipes
// @Accept       json
// @Produce      json
// @Param        recipe  body  CreateRecipeRequestDTO  true  "Recipe object"
// @Success      201     "Created"
// @Failure      400     {string}  string "invalid request"
// @Router       /recipes [post]
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

// Update recipe
// @Summary      Update a recipe
// @Description  Update an existing recipe's information
// @Tags         recipes
// @Accept       json
// @Produce      json
// @Param        id      path  int                     true  "Recipe ID"
// @Param        recipe  body  UpdateRecipeRequestDTO  true  "Updated recipe object"
// @Success      200     "OK"
// @Failure      400     {string}  string "invalid request"
// @Failure      404     {string}  string "recipe not found"
// @Failure      500     {string}  string "internal error"
// @Router       /recipes/{id} [put]
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

// Delete recipe
// @Summary      Delete a recipe
// @Description  Remove a recipe from the collection
// @Tags         recipes
// @Param        id   path  int  true  "Recipe ID"
// @Success      204  "No Content"
// @Failure      400  {string}  string "invalid id"
// @Failure      404  {string}  string "recipe not found"
// @Failure      500  {string}  string "internal error"
// @Router       /recipes/{id} [delete]
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
