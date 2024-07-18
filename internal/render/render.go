package render

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/charmbracelet/log"
)

type Client struct {
	apiToken    string
	apiEndpoint string
}

func NewClient(apiEndpoint, apiToken string) *Client {
	return &Client{
		apiToken:    apiToken,
		apiEndpoint: apiEndpoint,
	}
}

func (c *Client) requestAndParse(req *http.Request, dst any) error {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("%d response", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if len(body) == 0 {
		return nil
	}

	err = json.Unmarshal(body, dst)
	if err != nil {
		return fmt.Errorf("failed to parse json: %s", err)
	}

	return nil
}

func (c *Client) ListServices(ownerID string) ([]Service, error) {
	url := fmt.Sprintf("https://%s/v1/services?type=&limit=20", c.apiEndpoint)
	if ownerID != "" {
		url += fmt.Sprintf("&ownerId=%s", ownerID)
	}

	req, _ := http.NewRequest("GET", url, nil)

	var services []serviceResponse
	err := c.requestAndParse(req, &services)
	if err != nil {
		return nil, err
	}

	svcs := make([]Service, 0, len(services))
	for _, svc := range services {
		svcs = append(svcs, svc.Service)
	}

	return svcs, nil
}

type serviceResponse struct {
	Service Service
	Cursor  string
}

func (c *Client) deleteResource(resourceType, resourceID string) error {
	url := fmt.Sprintf("https://%s/v1/%s/%s", c.apiEndpoint, resourceType, resourceID)
	log.Info("deleting resource", "url", url)

	req, _ := http.NewRequest("DELETE", url, nil)

	var errResponse deleteResponse
	err := c.requestAndParse(req, &errResponse)
	if err != nil {
		return err
	}

	// In the event of success, we don't care about the response
	return nil
}

func (c *Client) DeleteService(serviceID string) error {
	return c.deleteResource("services", serviceID)
}

func (c *Client) ListAuthorizedOwners() ([]Owner, error) {
	url := fmt.Sprintf("https://%s/v1/owners?limit=20", c.apiEndpoint)

	req, _ := http.NewRequest("GET", url, nil)

	var ownerResps []ownerResponseObject
	err := c.requestAndParse(req, &ownerResps)
	if err != nil {
		return nil, err
	}

	owners := make([]Owner, 0, len(ownerResps))
	for _, owner := range ownerResps {
		owners = append(owners, owner.Owner)
	}

	return owners, nil
}

type ownerResponseObject struct {
	Owner  Owner
	Cursor string
}

type deleteResponse struct {
	Message string
}
