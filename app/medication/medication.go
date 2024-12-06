package medication

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	httpErrors "github.com/aborilov/hippo/api/sdk/http/errors"
	"github.com/aborilov/hippo/api/sdk/http/response"
	"github.com/aborilov/hippo/business/medication/model"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type App struct {
	service model.Service
	log     logr.Logger
}

func NewApp(log logr.Logger, svc model.Service) *App {
	return &App{
		service: svc,
		log:     log,
	}
}

func (app *App) RegisterHandlers(router *mux.Router) error {
	subrouter := router.PathPrefix("/medication").Subrouter()

	subrouter.Path("/").Methods("GET").HandlerFunc(app.List)
	subrouter.Path("/{id}").Methods("GET").HandlerFunc(app.Get)
	subrouter.Path("/{id}").Methods("DELETE").HandlerFunc(app.Delete)
	subrouter.Path("/").Methods("POST").HandlerFunc(app.Create)
	subrouter.Path("/{id}").Methods("PUT").HandlerFunc(app.Update)
	return nil
}

func (app *App) Update(w http.ResponseWriter, r *http.Request) {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		httpErrors.BadRequest(w, "Unable to obtain a medication ID from URL path")
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		httpErrors.BadRequest(w, fmt.Sprintf("unable to parse id: %s", err))
		return
	}
	_, err = app.service.Get(r.Context(), id)
	if err != nil {
		if errors.As(err, &model.ErrNotFound{}) {
			httpErrors.NotFound(w, err.Error())
			return
		}
		httpErrors.Internal(w, "unable to get medication", err)
		return
	}
	decoder := json.NewDecoder(r.Body)
	m := Medication{}
	if err := decoder.Decode(&m); err != nil {
		httpErrors.BadRequest(w, "Invalid JSON request body")
		return
	}
	s, err := m.ToService()
	if err != nil {
		httpErrors.Internal(w, "can't convert to service model", err)
		return
	}
	// force id from path
	s.ID = id
	n, err := app.service.Update(r.Context(), s)
	if err != nil {
		httpErrors.Internal(w, "unable to update medication", err)
		return
	}
	rv := serviceToMedication(n)
	response.WriteJSON(w, rv)
}

func (app *App) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	m := Medication{}
	if err := decoder.Decode(&m); err != nil {
		httpErrors.BadRequest(w, "Invalid JSON request body")
		return
	}
	s, err := m.ToService()
	if err != nil {
		httpErrors.Internal(w, "can't convert to service model", err)
		return
	}
	n, err := app.service.Create(r.Context(), s)
	if err != nil {
		httpErrors.Internal(w, "unable to create medication", err)
		return
	}
	rv := serviceToMedication(n)
	response.WriteJSON(w, rv)
}

func (app *App) List(w http.ResponseWriter, r *http.Request) {
	mm, err := app.service.List(r.Context())
	if err != nil {
		httpErrors.Internal(w, "unable to list medications", err)
		return
	}
	meds := []*Medication{}
	for _, m := range mm {
		meds = append(meds, serviceToMedication(m))
	}
	response.WriteJSON(w, meds)
}

func (app *App) Delete(w http.ResponseWriter, r *http.Request) {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		httpErrors.BadRequest(w, "Unable to obtain a medication ID from URL path")
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		httpErrors.BadRequest(w, fmt.Sprintf("unable to parse id: %s", err))
		return
	}
	if err := app.service.Delete(r.Context(), id); err != nil {
		httpErrors.Internal(w, "unable to get medication", err)
		return
	}

	response.WriteJSONWithStatus(w, http.StatusNoContent, nil)
}

func (app *App) Get(w http.ResponseWriter, r *http.Request) {
	idStr, ok := mux.Vars(r)["id"]
	if !ok {
		httpErrors.BadRequest(w, "Unable to obtain a medication ID from URL path")
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		httpErrors.BadRequest(w, fmt.Sprintf("unable to parse id: %s", err))
		return
	}
	m, err := app.service.Get(r.Context(), id)
	if err != nil {
		if errors.As(err, &model.ErrNotFound{}) {
			httpErrors.NotFound(w, err.Error())
			return
		}
		httpErrors.Internal(w, "unable to get medication", err)
		return
	}

	response.WriteJSON(w, serviceToMedication(m))
}
