package templates

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tarmo/internal/adapters/inbound/rest/common"
	"tarmo/internal/core/templates"
	"tarmo/internal/core/templates/ports/inbound"

	"github.com/go-chi/chi/v5"
)

var (
	ErrInternal         = "internal error"
	ErrInvalidId        = "invalid id"
	ErrInvalidReq       = "invalid request"
	ErrInvalidTemplate  = "invalid template"
	ErrTemplateNotFound = "template not found"
)

type Handler struct {
	service inbound.TemplatePort
}

func NewHandler(service inbound.TemplatePort) *Handler {
	return &Handler{service: service}
}

// GetAll templates
// @Summary      List all templates
// @Description  Get a list of all templates in the collection
// @Tags         templates
// @Produce      json
// @Success      200  {array}   TemplateListJSONResponseDTO
// @Failure      500  {string}  string "internal error"
// @Router       /templates [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	templates, err := h.service.GetAll()
	if err != nil {
		http.Error(w, ErrInternal, http.StatusInternalServerError)
		return
	}

	common.WriteJSON(w, http.StatusOK, ToResponseList(templates))
}

// GetByID template
// @Summary      Get a template by ID
// @Description  Get detailed information about a single template
// @Tags         templates
// @Produce      json
// @Param        id   path      int  true  "Template ID"
// @Success      200  {object}  TemplateJSONResponseDTO
// @Failure      400  {string}  string "invalid id"
// @Failure      404  {string}  string "template not found"
// @Failure      500  {string}  string "internal error"
// @Router       /templates/{id} [get]
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, ErrInvalidId, http.StatusBadRequest)
		return
	}

	template, err := h.service.GetByID(id)
	if err != nil {
		switch err {
		case templates.ErrTemplateNotFound:
			http.Error(w, ErrTemplateNotFound, http.StatusNotFound)
		default:
			http.Error(w, ErrInternal, http.StatusInternalServerError)
		}
		return
	}

	common.WriteJSON(w, http.StatusOK, ToResponse(&template))
}

// Create template
// @Summary      Create a new template
// @Description  Add a new template to the collection
// @Tags         templates
// @Accept       json
// @Produce      json
// @Param        template  body  CreateTemplateRequestDTO  true  "Template object"
// @Success      201     "Created"
// @Failure      400     {string}  string "invalid request"
// @Router       /templates [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateTemplateRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, ErrInvalidReq, http.StatusBadRequest)
		return
	}

	_, err := h.service.Create(ToCreateCommand(req))
	if err != nil {
		http.Error(w, ErrInvalidTemplate, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Update template
// @Summary      Update a template
// @Description  Update an existing template's information
// @Tags         templates
// @Accept       json
// @Produce      json
// @Param        id      path  int                     true  "Template ID"
// @Param        template  body  UpdateTemplateRequestDTO  true  "Updated template object"
// @Success      200     "OK"
// @Failure      400     {string}  string "invalid request"
// @Failure      404     {string}  string "template not found"
// @Failure      500     {string}  string "internal error"
// @Router       /templates/{id} [put]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, ErrInvalidReq, http.StatusBadRequest)
		return
	}

	var req UpdateTemplateRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, ErrInvalidReq, http.StatusBadRequest)
		return
	}

	cmd := ToUpdateCommand(req)
	cmd.ID = id

	err = h.service.Update(cmd)
	if err != nil {
		switch err {
		case templates.ErrTemplateNotFound:
			http.Error(w, ErrTemplateNotFound, http.StatusNotFound)
		default:
			http.Error(w, ErrInternal, http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete template
// @Summary      Delete a template
// @Description  Remove a template from the collection
// @Tags         templates
// @Param        id   path  int  true  "Template ID"
// @Success      204  "No Content"
// @Failure      400  {string}  string "invalid id"
// @Failure      404  {string}  string "template not found"
// @Failure      500  {string}  string "internal error"
// @Router       /templates/{id} [delete]
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
		case templates.ErrTemplateNotFound:
			http.Error(w, ErrTemplateNotFound, http.StatusNotFound)
		default:
			http.Error(w, ErrInternal, http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
