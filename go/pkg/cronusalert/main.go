package cronusalert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type cronusAlertClient struct {
	client HTTPClient
	token  string
}

type AlertType string

const (
	FIRING   AlertType = "firing"
	RESOLVED AlertType = "resolved"
	NEUTRAL  AlertType = "neutral"

	host = "https://cronusmonitoring.com"
)

type Alert struct {
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	Status      AlertType `json:"status"`
}

type alertReq struct {
	Status      string      `json:"status"`
	Annotations annotations `json:"annotations"`
}

type alerts struct {
	Alerts []alertReq `json:"alerts"`
}

type annotations struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

type CronusAlertClient interface {
	Fire(alert Alert) error
}

func buildReq(token string, alert Alert) (*http.Request, error) {
	alrt := alertReq{
		Status: string(alert.Status),
		Annotations: annotations{
			Summary:     alert.Summary,
			Description: alert.Description,
		},
	}

	alerts := alerts{
		Alerts: []alertReq{alrt},
	}

	b, err := json.Marshal(alerts)
	if err != nil {
		return nil, fmt.Errorf("failed to build alert %v", err)
	}

	uri := fmt.Sprintf("%s/api/alert", host)
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("failed to build request %v", err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func (c *cronusAlertClient) Fire(alert Alert) error {

	req, err := buildReq(c.token, alert)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to fire alert %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return fmt.Errorf(string(b))
	}

	return nil
}

func NewCronusAlertClient(token string) CronusAlertClient {
	return &cronusAlertClient{
		client: &http.Client{},
		token:  token,
	}
}
