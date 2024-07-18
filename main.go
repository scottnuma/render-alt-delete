package main

import (
	"fmt"
	"os"

	"github.com/scottnuma/render-alt-delete/internal/render"
	"github.com/scottnuma/render-alt-delete/internal/tui"
)

func main() {
	token, endpoint := render.GetConfig()
	client := render.NewClient(endpoint, token)

	p := tui.NewTUI(client)
	if err := p.Start(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
