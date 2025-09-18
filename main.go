package main

import (
	"fmt"
	"os"

	"github.com/scottnuma/render-alt-delete/internal/render"
	"github.com/scottnuma/render-alt-delete/internal/tui"
)

func main() {
	token, endpoint, err := render.GetConfig()
	if err != nil {
		fmt.Printf("failed to get config: %s\n", err)
		os.Exit(1)
	}
	client := render.NewClient(endpoint, token)

	p := tui.NewTUI(client)
	if err := p.Start(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
