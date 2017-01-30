package http_api

//Utility package for http api package

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/linusbohwalli/feature-toggle-service/go-client"
	"github.com/pkg/errors"
)

//Handler is used to handle routes
type Handler struct {
	*httprouter.Router
	Logger *log.Logger
	*go_client.Client
}

//NewHandler returns a new handler
func NewHandler() *Handler {

	h := &Handler{
		Router: httprouter.New(),
		Logger: log.New(os.Stderr, "\t", 0),
		Client: go_client.NewClient(),
	}

	h.GET("/featuretree/features", h.handleGetFeaturesForProperties)

	h.POST("/togglerule", h.handleCreateToggleRule)
	h.GET("/togglerule/:id", h.handleReadToggleRule)
	h.DELETE("/togglerule/:id", h.handleDeleteToggleRule)
	h.GET("/togglerule", h.handleSearchToggleRule)

	h.POST("/feature", h.handleCreateFeature)
	h.GET("/feature/:id", h.handleReadFeature)
	h.DELETE("/feature/:id", h.handleDeleteFeature)
	h.GET("/feature", h.handleSearchFeature)

	h.POST("/property", h.handleCreateProperty)
	h.GET("/property/:name", h.handleReadProperty)
	h.DELETE("/property/:name", h.handleDeleteProperty)
	h.GET("/property", h.handleSearchProperty)

	return h
}

//ServeHTTP sends request to correct subhandler
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasPrefix(r.URL.Path, "/featuretree/features"):
		h.Router.ServeHTTP(w, r)
	case strings.HasPrefix(r.URL.Path, "/togglerule"):
		h.Router.ServeHTTP(w, r)
	case strings.HasPrefix(r.URL.Path, "/feature"):
		h.Router.ServeHTTP(w, r)
	case strings.HasPrefix(r.URL.Path, "/property"):
		h.Router.ServeHTTP(w, r)
	default:
		http.NotFound(w, r)

	}
}

type errorResponse struct {
	Error string
}

//Error comment
func Error(w http.ResponseWriter, code int, initErr error, logger *log.Logger) {

	logger.Printf("%+v", initErr)
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(&errorResponse{Error: initErr.Error()}); err != nil {
		err = errors.New("Something went wrong during json encoding of error response")
		logger.Printf("%+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func createJSONandSendResponse(w http.ResponseWriter, v interface{}, logger *log.Logger) {

	if err := json.NewEncoder(w).Encode(v); err != nil {
		err = errors.New("Something went wrong during json encoding")
		logger.Printf("%+v", err)
		Error(w, http.StatusInternalServerError, err, logger)
		return
	}
}
