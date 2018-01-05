package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"

	"github.com/k0kubun/pp"
)

var payloadReg = regexp.MustCompile("payload=(.*)")

func payloadUnpacker(payload string) (string, error) {
	group := payloadReg.FindSubmatch([]byte(payload))
	if len(group) == 2 {
		return string(group[1]), nil
	}

	return "", fmt.Errorf("RegFail : %v", string(payload))
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading request body: err=%s", err)
		return
	}

	payloadStr, err := url.QueryUnescape(string(payload))

	payloadStr, err = payloadUnpacker(payloadStr)
	if err != nil {
		log.Printf("error reading request body: err=%s", err)
		return
	}
	pp.Print(payloadStr)

	var activity Activity
	if err := json.Unmarshal([]byte(payloadStr), &activity); err != nil {
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
