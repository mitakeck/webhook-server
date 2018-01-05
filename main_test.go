package main

import (
	"testing"
)

func TestPayloadUnpacker(t *testing.T) {
	payload := "payload=123"
	result, err := payloadUnpacker(payload)
	if err != nil {
		t.Fatalf("[FAIL] payloadUnpacker err: %v", err)
	}

	if result != "123" {
		t.Fatalf("[FAIL] payloadUnpacker ac: %v, ex: %v", result, "123")
	}
}

func TestPayloadUnpacker2(t *testing.T) {
	payload := "payload=aaaaaaaaaa="
	result, err := payloadUnpacker(payload)
	if err != nil {
		t.Fatalf("[FAIL] payloadUnpacker err: %v", err)
	}

	if result != "aaaaaaaaaa=" {
		t.Fatalf("[FAIL] payloadUnpacker ac: %v, ex: %v", result, "aaaaaaaaaa=")
	}
}

func TestPayloadUnpacker3(t *testing.T) {
	payload := "123123"
	_, err := payloadUnpacker(payload)
	if err == nil {
		t.Fatalf("[FAIL] payloadUnpacker payload: %v", payload)
	}
}
