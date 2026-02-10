package resources

import (
	"encoding/json"
	"net/http"
	"strconv"
	"tarmo/internal/adapters/inbound/rest/common"
	"tarmo/internal/core/resources"
	"tarmo/internal/core/resources/ports/inbound"

	"github.com/go-chi/chi/v5"
)

var (
	ErrInternal         = "internal error"
	ErrInvalidId        = "invalid id"
	ErrInvalidReq       = "invalid request"
	ErrInvalidResource  = "invalid resource"
	ErrResourceNotFound = "resource not found"
)

type Handler struct {
	service inbound.ResourcePort
}

func NewHandler(service inbound.ResourcePort) *Handler {
	return &Handler{service: service}
}

// GetAll resources
// @Summary      List all resources
// @Description  Get a list of all resources in the collection
// @Tags         resources
// @Produce      json
// @Success      200  {array}   ResourceJSONResponseDTO
// @Failure      500  {string}  string "internal error"
// @Router       /resources [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	resources, err := h.service.GetAll()
	if err != nil {
		http.Error(w, ErrInternal, http.StatusInternalServerError)
		return
	}

	common.WriteJSON(w, http.StatusOK, ToResponseList(resources))
}

// GetByID resource
// @Summary      Get a resource by ID
// @Description  Get detailed information about a single resource
// @Tags         resources
// @Produce      json
// @Param        id   path      int  true  "Resource ID"
// @Success      200  {object}  ResourceJSONResponseDTO
// @Failure      400  {string}  string "invalid id"
// @Failure      404  {string}  string "resource not found"
// @Failure      500  {string}  string "internal error"
// @Router       /resources/{id} [get]
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, ErrInvalidId, http.StatusBadRequest)
		return
	}

	resource, err := h.service.GetByID(id)
	if err != nil {
		switch err {
		case resources.ErrResourceNotFound:
			http.Error(w, ErrResourceNotFound, http.StatusNotFound)
		default:
			http.Error(w, ErrInternal, http.StatusInternalServerError)
		}
		return
	}

	common.WriteJSON(w, http.StatusOK, ToResponse(&resource))
}

// Create resource
// @Summary      Create a new resource
// @Description  Add a new resource to the collection
// @Tags         resources
// @Accept       json
// @Produce      json
// @Param        resource  body  CreateResourceRequestDTO  true  "Resource object"
// @Success      201     "Created"
// @Failure      400     {string}  string "invalid request"
// @Router       /resources [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateResourceRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, ErrInvalidReq, http.StatusBadRequest)
		return
	}

	_, err := h.service.Create(ToCreateCommand(req))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Update resource
// @Summary      Update a resource
// @Description  Update an existing resource's information
// @Tags         resources
// @Accept       json
// @Produce      json
// @Param        id      path  int                     true  "Resource ID"
// @Param        resource  body  UpdateResourceRequestDTO  true  "Updated resource object"
// @Success      200     "OK"
// @Failure      400     {string}  string "invalid request"
// @Failure      404     {string}  string "template not found"
// @Failure      500     {string}  string "internal error"
// @Router       /resources/{id} [put]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, ErrInvalidReq, http.StatusBadRequest)
		return
	}

	var req UpdateResourceRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, ErrInvalidReq, http.StatusBadRequest)
		return
	}

	cmd := ToUpdateCommand(req)
	cmd.ID = id

	err = h.service.Update(cmd)
	if err != nil {
		switch err {
		case resources.ErrResourceNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete resource
// @Summary      Delete a resource
// @Description  Remove a resource from the collection
// @Tags         resources
// @Param        id   path  int  true  "Resource ID"
// @Success      204  "No Content"
// @Failure      400  {string}  string "invalid id"
// @Failure      404  {string}  string "resource not found"
// @Failure      500  {string}  string "internal error"
// @Router       /resources/{id} [delete]
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
		case resources.ErrResourceNotFound:
			http.Error(w, ErrResourceNotFound, http.StatusNotFound)
		default:
			http.Error(w, ErrInternal, http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
