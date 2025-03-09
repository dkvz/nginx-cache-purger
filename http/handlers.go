package http

import (
	"encoding/json"
	"net/http"
)

type statusResponse struct {
	Success bool `json:"success"`
}

func successResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statusResponse{Success: true})
}

/*func jsonResponse(w http.ResponseWriter, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}*/

func (s *Server) handlePurgeRequest(w http.ResponseWriter, r *http.Request) {
	var b PurgeRequest
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// We should sanitize the request here but I'll do it
	// further down.
	s.requester.PurgeUrl(b.Url)
	successResponse(w)
}
