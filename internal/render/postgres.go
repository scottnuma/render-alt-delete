package render

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/scottnuma/render-alt-delete/internal/rad"
)

type PostgresResponseObject struct {
	Postgres rad.Postgres
	Cursor   string
}

func (c *Client) ListPostgres(ownerID string) ([]rad.Postgres, error) {
	url := fmt.Sprintf("https://%s/v1/postgres?limit=20", c.apiEndpoint)
	if ownerID != "" {
		url += fmt.Sprintf("&ownerId=%s", ownerID)
	}
	log.Info("got list postgres url", "url", url)

	req, _ := http.NewRequest("GET", url, nil)

	var dbresps []PostgresResponseObject
	err := c.requestAndParse(req, &dbresps)
	if err != nil {
		return nil, err
	}

	dbs := make([]rad.Postgres, 0, len(dbresps))
	for _, dbresp := range dbresps {
		dbs = append(dbs, dbresp.Postgres)
	}

	return dbs, nil
}

func (c *Client) DeletePostgres(postgresID string) error {
	return c.deleteResource("postgres", postgresID)
}
