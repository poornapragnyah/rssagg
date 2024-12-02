package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter,code int, msg string){
	if code > 499{
		log.Printf("Responding with 5XX error code:%v",msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w,code,errResponse{
		Error: msg,
	})

}


func respondWithJSON(w http.ResponseWriter, code int,payload interface{}){
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Failed to marshal JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
