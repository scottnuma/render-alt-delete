package render

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/scottnuma/render-alt-delete/pkg/rad"
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
		log.Println("failed to send HTTP request: ", err)
		return err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		log.Println("received non 2XX HTTP status: ", res.Status)
		return errors.New("non 2XX HTTP response status")
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("failed to read all of HTTP response body: ", err)
		return err
	}

	err = json.Unmarshal(body, dst)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ListServices(ownerID string) ([]rad.Service, error) {
	url := fmt.Sprintf("https://%s/v1/services?type=&limit=20", c.apiEndpoint)
	if ownerID != "" {
		url += fmt.Sprintf("&ownerId=%s", ownerID)
	}

	req, _ := http.NewRequest("GET", url, nil)

	var services []ServiceResponseObject
	err := c.requestAndParse(req, &services)
	if err != nil {
		return nil, err
	}

	svcs := make([]rad.Service, 0, len(services))
	for _, svc := range services {
		svcs = append(svcs, svc.Service)
	}

	return svcs, nil
}

type ServiceResponseObject struct {
	Service rad.Service
	Cursor  string
}

func (c *Client) DeleteService(serviceID string) error {
	url := fmt.Sprintf("https://%s/v1/services/%s", c.apiEndpoint, serviceID)

	req, _ := http.NewRequest("DELETE", url, nil)

	var errResponse DeleteResponse
	err := c.requestAndParse(req, &errResponse)
	if err != nil {
		return err
	}

	return fmt.Errorf(errResponse.Message)
}

func (c *Client) ListAuthorizedOwners() ([]rad.Owner, error) {
	url := fmt.Sprintf("https://%s/v1/owners?limit=20", c.apiEndpoint)

	req, _ := http.NewRequest("GET", url, nil)

	var ownerResps []OwnerResponseObject
	err := c.requestAndParse(req, &ownerResps)
	if err != nil {
		return nil, err
	}

	owners := make([]rad.Owner, 0, len(ownerResps))
	for _, owner := range ownerResps {
		owners = append(owners, owner.Owner)
	}

	return owners, nil
}

type OwnerResponseObject struct {
	Owner  rad.Owner
	Cursor string
}

type DeleteResponse struct {
	Message string
}
