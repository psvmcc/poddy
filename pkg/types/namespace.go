package types

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Namespaces [][]string

func (n *Namespaces) Get(serverEndpoint, token string) error {
	url := fmt.Sprintf("%s/api/v1/namespaces", serverEndpoint)
	client := http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return fmt.Errorf("Failed to generate request: %s", err)
	}

	req.Header.Set("User-Agent", fmt.Sprintf("poddy/%s (%s)", Version, Commit))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to send request: %s", err)
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("Wrong exit code: %d", res.StatusCode)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return fmt.Errorf("Failed to get request: %s", err)
	}

	jsonErr := json.Unmarshal(body, &n)
	if jsonErr != nil {
		return fmt.Errorf("Failed to unmarshal request: %s", err)
	}
	return nil
}

type Namespace struct {
	Name       string            `yaml:"Name" json:"name"`
	AccessType string            `yaml:"Access Type" json:"accessType"`
	Network    string            `yaml:"Network" json:"network"`
	ENV        map[string]string `yaml:"ENV" json:"env"`
}

func (n *Namespace) Get(serverEndpoint, token, name string) error {
	url := fmt.Sprintf("%s/api/v1/namespaces/%s", serverEndpoint, name)
	client := http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return fmt.Errorf("Failed to generate request: %s", err)
	}

	req.Header.Set("User-Agent", fmt.Sprintf("poddy/%s (%s)", Version, Commit))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Failed to send request: %s", err)
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("Wrong exit code: %d", res.StatusCode)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return fmt.Errorf("Failed to get request: %s", err)
	}

	jsonErr := json.Unmarshal(body, &n)
	if jsonErr != nil {
		return fmt.Errorf("Failed to unmarshal request: %s", err)
	}
	return nil
}
