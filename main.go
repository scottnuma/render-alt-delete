package main

import (
	"fmt"
	"os"

	"github.com/scottnuma/render-alt-delete/pkg/rad"
	"github.com/scottnuma/render-alt-delete/pkg/render"
	"github.com/scottnuma/render-alt-delete/pkg/tui"
)

func main() {
	token, endpoint := render.GetConfig()
	var client rad.RenderService = render.NewClient(endpoint, token)

	p := tui.NewTUI(client)
	if err := p.Start(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
