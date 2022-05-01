package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"resourcesAndFactories/pkg/domain/services"
)

type Handler struct {
	Router *mux.Router
	*services.Service
}

func New(service *services.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()
	h.Router.Use(cors)
	h.Router.HandleFunc("/dashboard", h.getDashboard).Methods("GET", "OPTIONS")
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
		return
	})
}

func (h *Handler) getDashboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dashboardData, err := h.Service.GetDashboardData()
	if err != nil {
		throwInternalError(w, err)
		return
	}
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"resources": dashboardData.Resources,
		"factories": dashboardData.Factories,
	}); err != nil {
		throwInternalError(w, err)
		return
	}
}
