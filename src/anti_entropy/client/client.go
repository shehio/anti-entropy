package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type NodeClient struct {
	baseURL string
	client  *http.Client
}

func NewNodeClient(host string, port int) *NodeClient {
	baseURL := host
	if port > 0 {
		baseURL = fmt.Sprintf("http://%s:%d", host, port)
	}
	return &NodeClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (nodeClient *NodeClient) UpdateState(key, value string) error {
	payload := map[string]string{
		"key":   key,
		"value": value,
	}
	
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	resp, err := nodeClient.client.Post(nodeClient.baseURL+"/state", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func (nodeClient *NodeClient) GetState() (map[string]string, error) {
	resp, err := nodeClient.client.Get(nodeClient.baseURL + "/state")
	if err != nil {
		return nil, fmt.Errorf("failed to get state: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var state map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&state); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return state, nil
}

func (nodeClient *NodeClient) TriggerGossip() error {
	resp, err := nodeClient.client.Post(nodeClient.baseURL+"/gossip", "", nil)
	if err != nil {
		return fmt.Errorf("failed to trigger gossip: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
} 