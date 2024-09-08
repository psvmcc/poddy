package types

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ConfigMap [][]string

func (n *ConfigMap) List(serverEndpoint, token, namespace string) error {
	url := fmt.Sprintf("%s/api/v1/namespaces/%s/configmaps", serverEndpoint, namespace)
	client := http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return fmt.Errorf("failed to generate request: %s", err)
	}

	req.Header.Set("User-Agent", fmt.Sprintf("poddy/%s (%s)", Version, Commit))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %s", err)
	}

	if res.StatusCode == 404 {
		return nil
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("wrong exit code: %d", res.StatusCode)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return fmt.Errorf("failed to get request: %s", err)
	}

	jsonErr := json.Unmarshal(body, &n)
	if jsonErr != nil {
		return fmt.Errorf("failed to unmarshal request: %s", err)
	}
	return nil
}

func (n *ConfigMap) Get(serverEndpoint, token, namespace, name string) ([]byte, error) {
	url := fmt.Sprintf("%s/api/v1/namespaces/%s/configmaps/%s", serverEndpoint, namespace, name)
	client := http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed to generate request: %s", err)
	}

	req.Header.Set("User-Agent", fmt.Sprintf("poddy/%s (%s)", Version, Commit))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %s", err)
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("wrong exit code: %d", res.StatusCode)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return nil, fmt.Errorf("failed to get request: %s", err)
	}

	return body, nil
}

func (n *ConfigMap) Delete(serverEndpoint, token, namespace, name string) error {
	url := fmt.Sprintf("%s/api/v1/namespaces/%s/configmaps/%s", serverEndpoint, namespace, name)
	client := http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return fmt.Errorf("failed to generate request: %s", err)
	}

	req.Header.Set("User-Agent", fmt.Sprintf("poddy/%s (%s)", Version, Commit))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Method = "DELETE"

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %s", err)
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("wrong exit code: %d", res.StatusCode)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	return nil
}
