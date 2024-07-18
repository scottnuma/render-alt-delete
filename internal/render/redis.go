package render

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/scottnuma/render-alt-delete/internal/rad"
)

type RedisResponseObject struct {
	Redis  rad.Redis
	Cursor string
}

func (c *Client) ListRedis(ownerID string) ([]rad.Redis, error) {
	url := fmt.Sprintf("https://%s/v1/redis?limit=20", c.apiEndpoint)
	if ownerID != "" {
		url += fmt.Sprintf("&ownerId=%s", ownerID)
	}
	log.Info("got list postgres url", "url", url)

	req, _ := http.NewRequest("GET", url, nil)

	var dbresps []RedisResponseObject
	err := c.requestAndParse(req, &dbresps)
	if err != nil {
		return nil, err
	}

	dbs := make([]rad.Redis, 0, len(dbresps))
	for _, dbresp := range dbresps {
		dbs = append(dbs, dbresp.Redis)
	}

	return dbs, nil
}

func (c *Client) DeleteRedis(postgresID string) error {
	return c.deleteResource("redis", postgresID)
}
