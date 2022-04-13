package render

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/scottnuma/render-alt-delete/pkg/rad"
)

type Client struct {
	apiToken string
}

func NewClient(apiToken string) *Client {
	return &Client{
		apiToken: apiToken,
	}
}

func (c *Client) ListServices(ownerID string) ([]rad.Service, error) {

	url := "https://api.render.com/v1/services?type=&limit=20"
	if ownerID != "" {
		url += fmt.Sprintf("&ownerId=%s", ownerID)
	}

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var services []ServiceResponseObject
	err := json.Unmarshal(body, &services)
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

	url := fmt.Sprintf("https://api.render.com/v1/services/%s", serviceID)

	req, _ := http.NewRequest("DELETE", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))

	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode == http.StatusNoContent {
		return nil
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var errResponse DeleteResponse
	err := json.Unmarshal(body, &errResponse)
	if err != nil {
		return nil
	}

	return fmt.Errorf(errResponse.Message)
}

func (c *Client) ListAuthorizedOwners() ([]rad.Owner, error) {
	url := "https://api.render.com/v1/owners?limit=20"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var ownerResps []OwnerResponseObject
	err := json.Unmarshal(body, &ownerResps)
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
