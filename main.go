package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/k0kubun/pp"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading request body: err=%s", err)
		return
	}

	defer r.Body.Close()

	var activity Activity
	if err := json.Unmarshal(payload, &activity); err != nil {
		fmt.Println("Json Unmarshal Error")
		return
	}
	pp.Print(activity)
}

func main() {
	log.Println("server starts...")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
