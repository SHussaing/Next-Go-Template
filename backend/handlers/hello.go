package handlers

import (
    "encoding/json"
    "net/http"
)

type Response struct {
    Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    response := Response{
        Message: "Hello, this is the Go server!",
    }

    w.WriteHeader(http.StatusOK)

    json.NewEncoder(w).Encode(response)
}
