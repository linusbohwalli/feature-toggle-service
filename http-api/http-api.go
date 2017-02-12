package http_api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	api "github.com/linusbohwalli/feature-toggle-service/api"
	"github.com/pkg/errors"
)

func (h *Handler) handleGetFeaturesForProperties(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var req api.GetFeaturesByPropertiesRequest

	props := make(map[string][]string)
	props = r.URL.Query()

	req.Properties = make(map[string]string)
	for k, vs := range props {
		for _, v := range vs {
			req.Properties[k] = v
		}
	}

	if err := h.Client.Open(); err != nil {
		err := errors.New("Failed to initialize connection to server")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}
	defer h.Client.Close()

	resp, err := h.Client.GetFeaturesForProperties(&req)
	if err != nil {
		err := errors.New("Failed to get features for properties")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}

	createJSONandSendResponse(w, resp, h.Logger)

}

func (h *Handler) handleCreateToggleRule(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var req api.CreateToggleRuleRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err := errors.New("Invalid JSON")
		Error(w, http.StatusBadRequest, err, h.Logger)
		return
	}

	if err := h.Client.Open(); err != nil {
		err := errors.New("Failed to open client connection")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}

	defer h.Client.Close()

	resp, err := h.Client.CreateToggleRule(&req)
	if err != nil {
		err := errors.New("Failed to create toggle rule")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}

	createJSONandSendResponse(w, resp, h.Logger)
}

func (h *Handler) handleReadToggleRule(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var req api.ReadToggleRuleRequest

	req.Id = ps.ByName("id")

	if err := h.Client.Open(); err != nil {
		err := errors.New("Failed to initialize connection to server")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}
	defer h.Client.Close()

	resp, err := h.Client.ReadToggleRule(&req)
	if err != nil {
		err := errors.New("Failed to read toggle rule")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}

	createJSONandSendResponse(w, resp, h.Logger)
}

func (h *Handler) handleDeleteToggleRule(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var req api.DeleteToggleRuleRequest

	req.Id = ps.ByName("id")

	if err := h.Client.Open(); err != nil {
		err := errors.New("Failed to initialize connection to server")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}
	defer h.Client.Close()

	resp, err := h.Client.DeleteToggleRule(&req)
	if err != nil {
		err := errors.New("Failed to delete toggle rule")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}

	createJSONandSendResponse(w, resp, h.Logger)
}

func (h *Handler) handleSearchToggleRule(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var req api.SearchToggleRuleRequest

	req.Name = r.URL.Query().Get("name")

	if err := h.Client.Open(); err != nil {
		err := errors.New("Failed to initialize connection to server")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}
	defer h.Client.Close()

	resp, err := h.Client.SearchToggleRule(&req)
	if err != nil {
		err := errors.New("Failed to search toggle rule")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}

	createJSONandSendResponse(w, resp, h.Logger)
}

func (h *Handler) handleCreateFeature(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var req api.CreateFeatureRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err := errors.New("Invalid JSON")
		Error(w, http.StatusBadRequest, err, h.Logger)
		return
	}

	if err := h.Client.Open(); err != nil {
		err := errors.New("Failed to initialize connection to server")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}
	defer h.Client.Close()

	resp, err := h.Client.CreateFeature(&req)
	if err != nil {
		err := errors.New("Failed to create feature")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}

	createJSONandSendResponse(w, resp, h.Logger)
}

func (h *Handler) handleReadFeature(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var req api.ReadFeatureRequest

	req.Id = ps.ByName("id")

	if err := h.Client.Open(); err != nil {
		err := errors.New("Failed to initialize connection to server")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}
	defer h.Client.Close()

	resp, err := h.Client.ReadFeature(&req)
	if err != nil {
		err := errors.New("Failed to read feature")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}

	createJSONandSendResponse(w, resp, h.Logger)
}

func (h *Handler) handleDeleteFeature(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var req api.DeleteFeatureRequest

	req.Id = ps.ByName("id")

	if err := h.Client.Open(); err != nil {
		err := errors.New("Failed to initialize connection to server")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}
	defer h.Client.Close()

	resp, err := h.Client.DeleteFeature(&req)
	if err != nil {
		err := errors.New("Failed to delete feature")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}

	createJSONandSendResponse(w, resp, h.Logger)
}

func (h *Handler) handleSearchFeature(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var req api.SearchFeatureRequest

	req.Name = r.URL.Query().Get("name")

	if err := h.Client.Open(); err != nil {
		err := errors.New("Failed to initialize connection to server")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}
	defer h.Client.Close()

	resp, err := h.Client.SearchFeature(&req)
	if err != nil {
		err := errors.New("Failed to search feature")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}

	createJSONandSendResponse(w, resp.Features, h.Logger)
}

func (h *Handler) handleCreateProperty(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var req api.CreatePropertyRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err := errors.New("Invalid JSON")
		Error(w, http.StatusBadRequest, err, h.Logger)
		return
	}

	if err := h.Client.Open(); err != nil {
		err := errors.New("Failed to initialize connection to server")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}
	defer h.Client.Close()

	resp, err := h.Client.CreateProperty(&req)
	if err != nil {
		err := errors.New("Failed to create property")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}

	createJSONandSendResponse(w, resp.Name, h.Logger)
}

func (h *Handler) handleReadProperty(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var req api.ReadPropertyRequest

	req.Name = ps.ByName("name")

	if err := h.Client.Open(); err != nil {
		err := errors.New("Failed to initialize connection to server")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}
	defer h.Client.Close()

	resp, err := h.Client.ReadProperty(&req)
	if err != nil {
		err := errors.New("Failed to read property")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}

	createJSONandSendResponse(w, resp, h.Logger)
}

func (h *Handler) handleDeleteProperty(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var req api.DeletePropertyRequest

	req.Name = ps.ByName("name")

	if err := h.Client.Open(); err != nil {
		err := errors.New("Failed to initialize connection to server")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}
	defer h.Client.Close()

	resp, err := h.Client.DeleteProperty(&req)
	if err != nil {
		err := errors.New("Failed to delete property")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}

	createJSONandSendResponse(w, resp, h.Logger)
}

func (h *Handler) handleSearchProperty(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var req api.SearchPropertyRequest

	req.Name = r.URL.Query().Get("name")

	if err := h.Client.Open(); err != nil {
		err := errors.New("Failed to initialize connection to server")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}
	defer h.Client.Close()

	resp, err := h.Client.SearchProperty(&req)
	if err != nil {
		err := errors.New("Failed to search property")
		Error(w, http.StatusInternalServerError, err, h.Logger)
		return
	}

	createJSONandSendResponse(w, resp.Properties, h.Logger)
}