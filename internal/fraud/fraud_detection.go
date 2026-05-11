package fraud

import (
	"net/http"
	"encoding/json"
)

func FraudScoreHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req FraudScoreRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	score := calculateFraudScore(req)
	resp := FraudScoreResponse{
		Approved:  score < 0.5,
		FraudScore: score,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func calculateFraudScore(req FraudScoreRequest) float64 {
	score := 0.0
	if req.Transaction.Amount > 1000 {
		score += 0.3
	}
	return score
}