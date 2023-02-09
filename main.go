package main

import (
	"fmt"
	"os"

	"github.com/scottnuma/render-alt-delete/pkg/rad"
	"github.com/scottnuma/render-alt-delete/pkg/render"
	"github.com/scottnuma/render-alt-delete/pkg/tui"
)

func main() {
	token := os.Getenv("RAD_RENDER_API_TOKEN")
	if token == "" {
		fmt.Println("RAD_RENDER_API_TOKEN is not set")
		os.Exit(1)
	}

	endpoint := os.Getenv("RAD_RENDER_API_ENDPOINT")
	if endpoint == "" {
		endpoint = "api.render.com"
	}
	var client rad.RenderService = render.NewClient(endpoint, token)

	p := tui.NewTUI(client)
	if err := p.Start(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
