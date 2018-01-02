package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	webhookData := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&webhookData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("payload: ")
	for k, v := range webhookData {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func main() {
	log.Println("server starts...")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
