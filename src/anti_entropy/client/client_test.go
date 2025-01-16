package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateState(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/state" {
			t.Errorf("expected /state path, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewNodeClient(server.URL, 0)
	err := client.UpdateState("test_key", "test_value")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGetState(t *testing.T) {
	expectedState := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/state" {
			t.Errorf("expected /state path, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(expectedState)
	}))
	defer server.Close()

	client := NewNodeClient(server.URL, 0)
	state, err := client.GetState()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(state) != len(expectedState) {
		t.Errorf("expected state with %d entries, got %d", len(expectedState), len(state))
	}
}

func TestTriggerGossip(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/gossip" {
			t.Errorf("expected /gossip path, got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewNodeClient(server.URL, 0)
	err := client.TriggerGossip()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
} 