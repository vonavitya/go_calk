package main

import (
	"calc_service/calc"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type CalcRequest struct {
	Expression string `json:"expression"`
}

type CalcResponse struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	var req CalcRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Expression is not valid"}`, http.StatusUnprocessableEntity)
		return
	}

	result, err := calc.Calc(req.Expression)
	if err != nil {
		if err.Error() == "division by zero" || strings.Contains(err.Error(), "invalid") {
			http.Error(w, `{"error": "Expression is not valid"}`, http.StatusUnprocessableEntity)
		} else {
			http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		}
		return
	}

	resp := CalcResponse{Result: fmt.Sprintf("%f", result)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api/v1/calculate", calculateHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
