package render

import (
	"fmt"
	"net/http"
)

type postgresResponse struct {
	Postgres Postgres
	Cursor   string
}

func (c *Client) ListPostgres(ownerID string) ([]Postgres, error) {
	url := fmt.Sprintf("https://%s/v1/postgres?limit=20", c.apiEndpoint)
	if ownerID != "" {
		url += fmt.Sprintf("&ownerId=%s", ownerID)
	}

	req, _ := http.NewRequest("GET", url, nil)

	var dbresps []postgresResponse
	err := c.requestAndParse(req, &dbresps)
	if err != nil {
		return nil, err
	}

	dbs := make([]Postgres, 0, len(dbresps))
	for _, dbresp := range dbresps {
		dbs = append(dbs, dbresp.Postgres)
	}

	return dbs, nil
}

func (c *Client) DeletePostgres(postgresID string) error {
	return c.deleteResource("postgres", postgresID)
}
