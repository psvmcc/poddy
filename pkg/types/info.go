package types

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Info struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
}

func (n *Info) Get(serverEndpoint, token string) error {
	url := fmt.Sprintf("%s/api/v1/info", serverEndpoint)
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

	if res.StatusCode == 404 {
		return nil
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
