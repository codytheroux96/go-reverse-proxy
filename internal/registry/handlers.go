package registry

import (
	"encoding/json"
	"net/http"
	"time"
)

func (reg *Registry) HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var srv Server
	if err := json.NewDecoder(r.Body).Decode(&srv); err != nil {
		http.Error(w, "invalid payload in request", http.StatusBadRequest)
		return
	}

	if srv.Name == "" || srv.BaseURL == "" || len(srv.Prefixes) == 0 {
		http.Error(w, "missing a required field in payload", http.StatusBadRequest)
		return
	}

	srv.RegisteredAt = time.Now()

	if err := reg.Register(srv); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message:": "server registration successful"})
}

func (reg *Registry) HandleDeregister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" {
		http.Error(w, "invalid payload in request", http.StatusBadRequest)
		return
	}

	if err := reg.Deregister(req.Name); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "server deregistered successfully"})
}

func (reg *Registry) HandleRegistryList(w http.ResponseWriter, r *http.Request) {

}
